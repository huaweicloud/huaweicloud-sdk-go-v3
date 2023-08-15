package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfEventTraceRequest Request Object
type ShowDetailOfEventTraceRequest struct {

	// 事件唯一标识
	TraceId string `json:"trace_id"`

	// 通道id
	ChannelId string `json:"channel_id"`
}

func (o ShowDetailOfEventTraceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventTraceRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventTraceRequest", string(data)}, " ")
}
