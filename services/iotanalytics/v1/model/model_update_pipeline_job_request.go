package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePipelineJobRequest struct {
	// 是否需要校验配置是否正确

	Check *bool `json:"check,omitempty"`
	// 管道ID

	PipelineId string `json:"pipeline_id"`
	// 管道作业详细配置，每个作业可选择不同的算子进行组合，各算子的使用方法详见：数据管道算子配置指南。

	Body map[string]interface{} `json:"body,omitempty"`
}

func (o UpdatePipelineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineJobRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipelineJobRequest", string(data)}, " ")
}
