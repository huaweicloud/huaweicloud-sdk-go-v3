package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendMessagesReq struct {
	// 发送消息成功后，是否返回Message ID，默认为false，设置为true时，返回参数才有Message ID。

	ReturnId *bool `json:"return_id,omitempty"`
	// 消息列表。

	Messages []SendMessageEntity `json:"messages"`
}

func (o SendMessagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessagesReq struct{}"
	}

	return strings.Join([]string{"SendMessagesReq", string(data)}, " ")
}
