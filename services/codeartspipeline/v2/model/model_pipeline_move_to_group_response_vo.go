package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineMoveToGroupResponseVo PipelineMoveToGroupResponseVo
type PipelineMoveToGroupResponseVo struct {

	// 响应码 [\"failed\", \"success\"]
	Code string `json:"code"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线名
	PipelineName string `json:"pipeline_name"`
}

func (o PipelineMoveToGroupResponseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineMoveToGroupResponseVo struct{}"
	}

	return strings.Join([]string{"PipelineMoveToGroupResponseVo", string(data)}, " ")
}
