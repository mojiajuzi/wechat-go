package wechat

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

var client http.Client

func init() {
	c := http.Client{}
	client = c
}

//Wechat 微信基本配置
type Wechat struct {
	AppID  string `json:"appid"`
	Secret string `json:"secret"`
	Token  string `json:"token"`
	AesKey string `json:"aes_key"`
}

//AccessToken 微信access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
	ErrorCode string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

//CheckSignature 判断消息是否来自微信
func (w *Wechat) CheckSignature(signature, nonce, echostr string, timestamp int64) (r bool) {
	//1: 将timestamp,nonce,token进行字典排序
	timeString := strconv.FormatInt(timestamp, 10)
	n := []string{w.Token, timeString, nonce}
	sort.Strings(n)

	//2:将三个参数字符串拼接成一个字符串
	var needEncodeStr string

	for _, s := range n {
		needEncodeStr += s
	}

	//3:进行sha1加密
	h := sha1.New()
	h.Write([]byte(needEncodeStr))

	bs := h.Sum(nil)

	//4:开发者获得加密后的字符串可与signature对比
	r = strings.EqualFold(signature, string(bs))
	return
}

//AccessToken 获取AccessToken
func (w *Wechat) AccessToken() (AccessToken, error) {
	uri := "https://api.weixin.qq.com/cgi-bin/token"
	v := url.Values{}
	v.Set("grant_type", "client_credential")
	v.Set("appid", w.AppID)
	v.Set("secret", w.Secret)
	url := uri + "?" + v.Encode()
	resp, err := http.Get(url)
	token := AccessToken{}
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return token, err
		}
		json.Unmarshal(data, &token)

		if !strings.EqualFold(token.Token, "") {
			return token, nil
		}
		return token, errors.New(token.ErrMsg)
	}
	return token, errors.New(resp.Status)
}
