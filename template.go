package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
)

//TemplateMessageURL 消息基本路径
const TemplateMessageURL = "https://api.weixin.qq.com/cgi-bin/template/"

// Template 模板
type Template struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
	AccessToken     string
}

//Industry 设置行业
func (t *Template) Industry(primary, second int) (CommonResp, error) {
	u := TemplateMessageURL + "api_set_industry"

	v := url.Values{}
	v.Add("access_token", t.AccessToken)
	v.Add("industry_id1", strconv.Itoa(primary))
	v.Add("industry_id2", strconv.Itoa(second))

	res := CommonResp{}
	resp, err := client.PostForm(u, v)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(content, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

//GetTemplateID 获取模板Id
func (t *Template) GetTemplateID(shortID string) (string, error) {
	u := TemplateMessageURL + "api_add_template"

	v := url.Values{}
	v.Add("access_token", t.AccessToken)
	v.Add("template_id_short", shortID)

	resp, err := client.PostForm(u, v)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res struct {
		CommonResp,
		TemplateIDTemp string
	}

	json.Unmarshal(content, &res)
	if err != nil {
		return "", err
	}
	return res.TemplateIDTemp, nil
}

//List 获取模板消息列表
func (t *Template) List() ([]Template, error) {
	u := TemplateMessageURL + "get_all_private_template"

	v := url.Values{}
	v.Add("access_token", t.AccessToken)
	furl := u + "?" + v.Encode()
	resp, err := client.Get(furl)
	var res struct {
		TemplateList []Template `json:"template_list"`
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}

	return res.TemplateList, nil
}

//Delete 删除模板消息
func (t *Template) Delete() (CommonResp, error) {
	u := TemplateMessageURL + "del_private_template"

	v := url.Values{}
	v.Add("access_token", t.AccessToken)
	v.Add("template_id", t.TemplateID)

	resp, err := client.PostForm(u, v)

	res := CommonResp{}
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(content, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

//TemplateMessage 模板消息
type TemplateMessage struct {
	OpenID             string `json:"touser"`
	TemplateID         string `json:"template_id"`
	URL                string `json:"url"`
	MiniProgramMessage `json:"miniprogram"`
	Content            []map[string]ContentMessage `json:"data"`
}

//ContentMessage 消息内容
type ContentMessage struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

//MiniProgramMessage 模板消息发送小程序
type MiniProgramMessage struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

//Send 发送模板消息
func (t *Template) Send(m TemplateMessage) (int, error) {
	u := "https://api.weixin.qq.com/cgi-bin/message/template/send"
	var res struct {
		CommonResp,
		MsgID int
	}

	data, err := json.Marshal(m)
	if err != nil {
		return res.MsgID, err
	}

	resp, err := client.Post(u, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return res.MsgID, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res.MsgID, err
	}

	err = json.Unmarshal(content, &res)
	if err != nil {
		return res.MsgID, err
	}
	return res.MsgID, nil
}
