package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSubscriberInPic 子画面信息。
type RestSubscriberInPic struct {

	// 多画面中每个画面的编号。编号从1开始。默认值为“1”。
	Index *int32 `json:"index,omitempty"`

	// 每个画面中与会者号码，从[[查询会议实时信息](https://support.huaweicloud.com/api-meeting/meeting_21_0029.html)](tag:hws)[[查询会议实时信息](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0029.html)](tag:hk)接口返回participants中的phone中获取。
	Subscriber *[]string `json:"subscriber,omitempty"`

	// 是否为辅流。默认值为0。 * 0： 不是辅流 * 1： 是辅流
	IsAssistStream *int32 `json:"isAssistStream,omitempty"`
}

func (o RestSubscriberInPic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSubscriberInPic struct{}"
	}

	return strings.Join([]string{"RestSubscriberInPic", string(data)}, " ")
}
