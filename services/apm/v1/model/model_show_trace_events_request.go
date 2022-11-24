package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTraceEventsRequest struct {

	// 调用链ID。
	TraceId string `json:"trace_id"`
}

func (o ShowTraceEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTraceEventsRequest struct{}"
	}

	return strings.Join([]string{"ShowTraceEventsRequest", string(data)}, " ")
}
