package wechat

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

//MsgTypeText 文本
const MsgTypeText = "text"

//MsgTypeImage 图片
const MsgTypeImage = "image"

//MsgTypeVoice 语音
const MsgTypeVoice = "voice"

//MsgTypeVedio 视频
const MsgTypeVedio = "vedio"

//MsgTypeMusic 音乐
const MsgTypeMusic = "music"

//MsgTypeNews 图文
const MsgTypeNews = "news"

//Message 微信消息结构体
type Message struct {
	AppID string `xml:"ToUserName" json:"ToUserName"`

	OpenID string `xml:"FromUserName" json:"FromUserName"`

	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`

	MsgType string `xml:"MsgType" json:"MsgType"`

	MediaID string `xml:"MediaId" json:"MediaId"`

	MsgID string `xml:"MsgId" json:"MsgId"`

	Content string `xml:"Content" json:"Content"`

	PicURL string `xml:"PicUrl" json:"PicUrl"`

	Format string `xml:"Format" json:"Format"`

	Recognition string `xml:"Recognition" json:"Recognition"`

	ThumbMediaID string `xml:"ThumbMediaId" json:"ThumbMediaId"`

	LocationX string `xml:"Location_X" json:"Location_X"`

	LocationY string `xml:"Location_Y" json:"Location_Y"`

	Scale string `xml:"Scale" json:"Scale"`

	Lable string `xml:"Label" json:"Label"`

	Title string `xml:"Title" json:"Title"`

	Description string `xml:"Description" json:"Description"`

	URL string `xml:"Url" json:"Url"`

	Event string `xml:"Event" json:"Event"`

	EventKey string `xml:"EventKey" json:"EventKey"`

	Ticket string `xml:"Ticket" json:"Ticket"`

	Latitude string `xml:"Latitude" json:"Latitude"`

	Longitude string `xml:"Longitude" json:"Longitude"`

	Precision string `xml:"Precision" json:"Precision"`
}

//BaseMessage 基础结构
type BaseMessage struct {
	ToUserName string `xml:"ToUserName" json:"ToUserName"`

	FromUserName string `xml:"FromUserName" json:"FromUserName"`

	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`

	MsgType string `xml:"MsgType" json:"MsgType"`
}

//CommonAttr 设置消息通用属性
func (bm *BaseMessage) CommonAttr(openID, appID, mType string) {
	bm.ToUserName = openID
	bm.FromUserName = appID
	bm.MsgType = mType
	bm.CreateTime = time.Now().Unix()
}

//XML 格式化成xml格式
func (bm *BaseMessage) XML(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

//JSON 格式化成json格式
func (bm *BaseMessage) JSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

//TextMessage 文字回复
type TextMessage struct {
	XMLName xml.Name `xml:"xml"`
	Text
	BaseMessage
}

//Text 文字消息
type Text struct {
	Content string `xml:"Content" json:"Content"`
}

//ImageMessage 图片消息
type ImageMessage struct {
	XMLName xml.Name `xml:"xml"`
	Image   `xml:"Image" json:"Image"`
	BaseMessage
}

//Image 图片结构
type Image struct {
	MediaID string `xml:"MediaId" json:"MediaId"`
}

//VoiceMessage 语音消息
type VoiceMessage struct {
	XMLName xml.Name `xml:"xml"`
	BaseMessage
	Voice
}

// Voice 语言结构
type Voice struct {
	XMLName xml.Name `xml:"Voice" json:"Voice"`
	MediaID string   `xml:"MediaId" json:"MediaId"`
}

//VideoMessage 视频消息
type VideoMessage struct {
	XMLName xml.Name `xml:"xml"`
	BaseMessage
	Video `xml:"Video" json:"Video"`
}

//Video 视频
type Video struct {
	MediaID     string `xml:"MediaId" json:"MediaId"`
	Title       string `xml:"Title" json:"Title"`
	Description string `xml:"Description" json:"Description"`
}

//MusicMessage 音乐消息
type MusicMessage struct {
	XMLName xml.Name `xml:"xml"`
	BaseMessage
	Music
}

//Music 音乐
type Music struct {
	XMLName      xml.Name `xml:"Music"`
	Title        string   `xml:"Title" json:"Title"`
	Description  string   `xml:"Description" json:"Description"`
	MusicURL     string   `xml:"MusicUrl" json:"MusicUrl"`
	HQMusicURL   string   `xml:"HQMusicUrl" json:"HQMusicUrl"`
	ThumbMediaID string   `xml:"ThumbMediaId" json:"ThumbMediaId"`
}

//NewsMessage 图文消息结构体
type NewsMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ArticleCount int      `xml:"ArticleCount" json:"ArticleCount"`
	BaseMessage
	Article
}

// Article 图文列表
type Article struct {
	XMLNAME     xml.Name `xml:"Articles" json:"Articles"`
	ArticleItem []ArticleItem
}

// ArticleItem 文章
type ArticleItem struct {
	XMLNAME     xml.Name `xml:"itme"`
	Title       string   `xml:"Title" json:"Title"`
	Description string   `xml:"Description" json:"Description"`
	PicURL      string   `xml:"PicUrl" json:"PicUrl"`
	URL         string   `xml:"Url" json:"Url"`
}
