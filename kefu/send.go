package kefu

import (
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/util"
)

const (
	sendMessageURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

// Kefu 客服
type Kefu struct {
	*context.Context

	openID string
}

// New 实例化
func New(context *context.Context) *Kefu {
	kefu := new(Kefu)
	kefu.Context = context
	return kefu
}

// Send 发送客服消息
func (r *Kefu) Send(openID string) *Kefu {
	r.openID = openID
	return r
}

// Text 发送文字
func (r *Kefu) Text(msg string) (err error) {
	return r.send("text", msg)
}

// Image 发送图片
func (r *Kefu) Image(mediaID string) (err error) {
	return r.send("image", mediaID)
}

// Voice 发送语音
func (r *Kefu) Voice(mediaID string) (err error) {
	return r.send("voice", mediaID)
}

// News 发送图文
func (r *Kefu) News(articles []*message.Article) (err error) {
	var news []map[string]string
	for _, v := range articles {
		news = append(news, map[string]string{"title": v.Title, "description": v.Description, "url": v.URL, "picurl": v.PicURL})
	}
	return r.send("news", news)
}

// MpNews 发送图文
func (r *Kefu) MpNews(mediaID string) (err error) {
	return r.send("mpnews", mediaID)
}

func (r *Kefu) send(typ string, msg interface{}) (err error) {
	if r.openID == "" {
		return fmt.Errorf("open_id empty")
	}

	body := map[string]interface{}{
		"touser":  r.openID,
		"msgtype": typ,
	}
	switch typ {
	case "text":
		body[typ] = map[string]interface{}{"content": msg}
	case "image", "voice", "mpnews":
		body[typ] = map[string]interface{}{"media_id": msg}
	case "news":
		body[typ] = map[string]interface{}{"articles": msg}
	default:
		return fmt.Errorf("invalid type: " + typ)
	}

	var accessToken string
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(sendMessageURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, body)
	if err != nil {
		return
	}

	return util.DecodeWithCommonError(response, "SendKefuMessage")
}
