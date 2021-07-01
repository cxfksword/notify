package wework

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fastwego/wxwork/corporation"
	api "github.com/fastwego/wxwork/corporation/apis/message"
	"github.com/pkg/errors"
)

// WeWork struct holds necessary data to send wechat corporation message.
type WeWork struct {
	corpID     string
	agentID    string
	corpSecret string
	toUser     []string
	msgType    string

	url string
}

// New returns a new instance of a wechat corporation message service.
// For more information about wechat corporation api token:
//    -> https://work.weixin.qq.com/api/doc/90000/90135/90665
func New(corpID, agentID, corpSecret string) *WeWork {
	return &WeWork{
		corpID:     corpID,
		agentID:    agentID,
		corpSecret: corpSecret,
		toUser:     []string{},
		msgType:    "textcard",
	}
}

func (m *WeWork) SetUrl(url string) {
	m.url = url
}

// AddReceivers takes wechat corporation user IDs and adds them to the internal user ID list. The Send method will send
// a given message to all those users. To send to all user use "@all".
func (m *WeWork) AddReceivers(user ...string) {
	m.toUser = append(m.toUser, user...)
}

// Send takes a message subject and a message body and sends them to all previously set twitterIDs as a DM.
// See https://work.weixin.qq.com/api/doc/90000/90135/90236
func (m WeWork) Send(subject, message string) error {
	corp := corporation.New(corporation.Config{Corpid: m.corpID})
	app := corp.NewApp(corporation.AppConfig{
		AgentId: m.agentID,
		Secret:  m.corpSecret,
	})

	data := payload{
		AgentID: m.agentID,
		MsgType: m.msgType,
		ToUser:  strings.Join(m.toUser, "|"),
		Textcard: textCard{
			Title:       subject,
			Description: message,
			Url:         m.url,
		},
	}
	payloadJSON, _ := json.Marshal(data)

	resp, err := api.Send(app, payloadJSON)
	if err != nil {
		return errors.Wrapf(err, "failed to send message to wechat corporation users '%s'", data.ToUser)
	}

	var result apiResult
	err = json.Unmarshal(resp, &result)
	if err == nil && result.ErrCode != 0 {
		return errors.Wrapf(fmt.Errorf(result.ErrMsg), "failed to send message to wechat corporation users '%s'", data.ToUser)
	}

	return nil
}
