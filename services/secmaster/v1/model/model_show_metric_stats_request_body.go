package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMetricStatsRequestBody struct {

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 结束时间
	EndTimestamp int64 `json:"end_timestamp"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 开始时间
	StartTimestamp int64 `json:"start_timestamp"`
}

func (o ShowMetricStatsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricStatsRequestBody struct{}"
	}

	return strings.Join([]string{"ShowMetricStatsRequestBody", string(data)}, " ")
}
