package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobPipelineInfoParamsItems struct {

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`
}

func (o JobPipelineInfoParamsItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobPipelineInfoParamsItems struct{}"
	}

	return strings.Join([]string{"JobPipelineInfoParamsItems", string(data)}, " ")
}
