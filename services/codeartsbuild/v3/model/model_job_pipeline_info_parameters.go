package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobPipelineInfoParameters 任务流水线参数
type JobPipelineInfoParameters struct {

	// 地域
	Region *string `json:"region,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 参数值
	Params *[]JobPipelineInfoParamsItems `json:"params,omitempty"`
}

func (o JobPipelineInfoParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobPipelineInfoParameters struct{}"
	}

	return strings.Join([]string{"JobPipelineInfoParameters", string(data)}, " ")
}
