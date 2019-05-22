package wechat

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//CustomerActionCreate 创建
const CustomerActionCreate = 1

//CustomerActionUpdate 更新
const CustomerActionUpdate = 2

//CustomerActionDelete 删除
const CustomerActionDelete = 3

//Customer 客服账号
type Customer struct {
	Account  string `json:"kf_account"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	ID       string `json:"kf_id"`
	Avatar   string `json:"kf_headimgurl"`
}

//CommonResp 客服操作响应
type CommonResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (c CommonResp) string() string {
	return c.ErrMsg
}

//CustomerList 客服列表
type CustomerList struct {
	List []Customer `json:"kf_list"`
}

//Action 客服的增加，更新删除操作
func (c *Customer) Action(token string, actionType int) CommonResp {
	var uri string

	switch actionType {
	case CustomerActionCreate:
		uri = "https://api.weixin.qq.com/customservice/kfaccount/add"
	case CustomerActionUpdate:
		uri = "https://api.weixin.qq.com/customservice/kfaccount/update"
	case CustomerActionDelete:
		uri = "https://api.weixin.qq.com/customservice/kfaccount/del"
	}
	v := url.Values{}
	v.Add("kf_account", c.Account)
	v.Add("nickname", c.NickName)
	v.Add("access_token", token)

	h := md5.New()
	h.Write([]byte(c.Password))
	p := fmt.Sprintf("%s", h.Sum(nil))
	v.Add("passwrod", p)
	body := strings.NewReader(v.Encode())

	res, err := http.Post(uri, "application/x-www-form-urlencoded", body)
	r := CommonResp{}
	if err != nil {
		r.ErrCode = 1
		r.ErrMsg = err.Error()
	} else {
		defer res.Body.Close()
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			r.ErrCode = 2
			r.ErrMsg = err.Error()
		} else {
			json.Unmarshal(content, &r)
		}
	}
	return r
}

//SetAvatar 创建用户头像
func (c *Customer) SetAvatar(token string) string {
	//TODO
	return token
}

//List 获取客服列表
func (c *Customer) List(token string) ([]Customer, CommonResp) {
	uri := "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"

	v := url.Values{}
	v.Add("access_token", token)
	url := uri + "?" + v.Encode()
	resp, err := http.Get(url)
	r := CommonResp{}
	var cList []Customer
	if err != nil {
		r.ErrCode = 1
		r.ErrMsg = err.Error()
		return cList, r
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.ErrCode = 2
		r.ErrMsg = err.Error()
		return cList, r
	}

	l := CustomerList{}

	err = json.Unmarshal(content, &l)
	if err != nil {
		r.ErrCode = 3
		r.ErrMsg = err.Error()
		return cList, r
	}

	return l.List, r
}
