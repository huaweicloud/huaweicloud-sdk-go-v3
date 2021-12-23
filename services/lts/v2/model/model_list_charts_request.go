package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListChartsRequest struct {
	// 日志组ID

	LogGroupId string `json:"log_group_id"`
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
}

func (o ListChartsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChartsRequest struct{}"
	}

	return strings.Join([]string{"ListChartsRequest", string(data)}, " ")
}
