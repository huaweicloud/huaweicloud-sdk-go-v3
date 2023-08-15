package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriberInPic 子画面信息。
type SubscriberInPic struct {

	// 多画面中每个画面的编号。编号从1开始。 默认值为1。
	Index int32 `json:"index"`

	// 每个画面中与会者标识列表，从[[查询会议实时信息](https://support.huaweicloud.com/api-meeting/meeting_21_0029.html)](tag:hws)[[查询会议实时信息](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0029.html)](tag:hk)接口返回的pid中获取。
	Subscriber *[]string `json:"subscriber,omitempty"`

	// 是否为辅流。默认值为0。 - 0: 不是辅流 - 1: 是辅流
	IsAssistStream *int32 `json:"isAssistStream,omitempty"`
}

func (o SubscriberInPic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriberInPic struct{}"
	}

	return strings.Join([]string{"SubscriberInPic", string(data)}, " ")
}
