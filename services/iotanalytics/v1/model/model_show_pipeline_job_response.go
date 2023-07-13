package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineJobResponse Response Object
type ShowPipelineJobResponse struct {

	// 管道作业详细配置，每个作业可选择不同的算子进行组合，各算子的使用方法详见：数据管道算子配置指南。
	PipelineConfig map[string]interface{} `json:"pipeline_config,omitempty"`

	PipelineInfo   *PipelineJobInfoDto `json:"pipeline_info,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowPipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineJobResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineJobResponse", string(data)}, " ")
}
