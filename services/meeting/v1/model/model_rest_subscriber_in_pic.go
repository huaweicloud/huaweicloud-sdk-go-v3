package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子画面信息
type RestSubscriberInPic struct {
	// 多画面中每个画面的编号，编号从1开始

	Index *int32 `json:"index,omitempty"`
	// 每个画面中会话标识，即Call_ID，通过会议状态通知获取

	Subscriber *[]string `json:"subscriber,omitempty"`
	// 是否为辅流 0： 不是辅流 1： 是辅流

	IsAssistStream *int32 `json:"isAssistStream,omitempty"`
}

func (o RestSubscriberInPic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSubscriberInPic struct{}"
	}

	return strings.Join([]string{"RestSubscriberInPic", string(data)}, " ")
}
