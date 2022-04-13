package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineParameter struct {
	// 参数名称

	Name string `json:"name"`
	// 参数值

	Value string `json:"value"`
}

func (o PipelineParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineParameter struct{}"
	}

	return strings.Join([]string{"PipelineParameter", string(data)}, " ")
}
