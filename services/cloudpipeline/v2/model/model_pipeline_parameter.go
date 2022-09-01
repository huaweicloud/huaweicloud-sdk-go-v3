package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineParameter struct {

	// 参数名称
	Name string `json:"name" xml:"name"`

	// 参数值
	Value string `json:"value" xml:"value"`
}

func (o PipelineParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineParameter struct{}"
	}

	return strings.Join([]string{"PipelineParameter", string(data)}, " ")
}
