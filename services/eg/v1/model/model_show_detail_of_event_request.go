package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfEventRequest Request Object
type ShowDetailOfEventRequest struct {

	// 追踪事件的uniqueId
	TraceId string `json:"trace_id"`

	// 指定查询的事件通道ID
	ChannelId *string `json:"channel_id,omitempty"`
}

func (o ShowDetailOfEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventRequest", string(data)}, " ")
}
