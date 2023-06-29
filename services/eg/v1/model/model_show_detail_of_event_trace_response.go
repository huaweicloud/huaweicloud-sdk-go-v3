package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfEventTraceResponse Response Object
type ShowDetailOfEventTraceResponse struct {

	// 事件ID
	EventId *string `json:"eventId,omitempty"`

	// 事件源
	EventSource *string `json:"eventSource,omitempty"`

	// 事件类型
	EventType *string `json:"eventType,omitempty"`

	// 接收时间
	ReceiveTime *string `json:"receiveTime,omitempty"`

	// 通道ID
	ChannelId *string `json:"channelId,omitempty"`

	// 通道名称
	ChannelName *string `json:"channelName,omitempty"`

	// 事件投递列表
	DeliverList    *[]DeliverItem `json:"deliverList,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDetailOfEventTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventTraceResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventTraceResponse", string(data)}, " ")
}
