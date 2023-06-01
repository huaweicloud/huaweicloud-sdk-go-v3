package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailOfEventRequest struct {

	// 追踪事件的uniqueId
	TraceId string `json:"trace_id"`
}

func (o ShowDetailOfEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventRequest", string(data)}, " ")
}
