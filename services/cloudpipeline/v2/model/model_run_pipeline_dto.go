package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 运行流水线请求体
type RunPipelineDto struct {

	// 使用的源
	Sources *[]RunPipelineDtoSources `json:"sources,omitempty"`

	// 运行描述
	Description *string `json:"description,omitempty"`

	// 使用的自定义参数
	Variables *[]RunPipelineDtoVariables `json:"variables,omitempty"`

	// 选择的任务
	ChooseJobs *[]string `json:"choose_jobs,omitempty"`
}

func (o RunPipelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDto struct{}"
	}

	return strings.Join([]string{"RunPipelineDto", string(data)}, " ")
}
