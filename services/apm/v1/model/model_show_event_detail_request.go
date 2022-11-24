package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEventDetailRequest struct {

	// trace id。
	TraceId string `json:"trace_id"`

	// span id。
	SpanId string `json:"span_id"`

	// event id。
	EventId string `json:"event_id"`

	// 环境id。
	EnvId int64 `json:"env_id"`
}

func (o ShowEventDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEventDetailRequest", string(data)}, " ")
}
