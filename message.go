package wechat

import (
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

//Text 回复微信文字消息
func (m Message) Text(content Text) TextMessage {
	t := TextMessage{}
	t.Text = content
	t.ToUserName = m.OpenID
	t.FromUserName = m.AppID
	t.CreateTime = time.Now().Unix()
	t.MsgType = MsgTypeText

	return t
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

//Image 图片消息
func (m Message) Image(image Image) ImageMessage {
	t := ImageMessage{}
	t.Image = image
	t.ToUserName = m.OpenID
	t.FromUserName = m.AppID
	t.CreateTime = time.Now().Unix()
	t.MsgType = MsgTypeImage

	return t
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

//Voice 返回语音消息
func (m Message) Voice(voice Voice) VoiceMessage {
	t := VoiceMessage{}
	t.Voice = voice
	t.ToUserName = m.OpenID
	t.FromUserName = m.AppID
	t.CreateTime = time.Now().Unix()
	t.MsgType = MsgTypeImage

	return t
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

//Video 视频消息返回
func (m Message) Video(video Video) VideoMessage {
	t := VideoMessage{}
	t.Video = video
	t.ToUserName = m.OpenID
	t.FromUserName = m.AppID
	t.CreateTime = time.Now().Unix()
	t.MsgType = MsgTypeVedio

	return t
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

//Music 音乐消息返回
func (m Message) Music(music Music) MusicMessage {
	t := MusicMessage{}
	t.Music = music
	t.ToUserName = m.OpenID
	t.FromUserName = m.AppID
	t.CreateTime = time.Now().Unix()
	t.MsgType = MsgTypeMusic

	return t
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

//News 图文消息
func (m Message) News(Article []ArticleItem) NewsMessage {
	t := NewsMessage{}
	t.Article.ArticleItem = Article
	t.ArticleCount = len(Article)
	t.ToUserName = m.OpenID
	t.FromUserName = m.AppID
	t.CreateTime = time.Now().Unix()
	t.MsgType = MsgTypeNews

	return t
}
