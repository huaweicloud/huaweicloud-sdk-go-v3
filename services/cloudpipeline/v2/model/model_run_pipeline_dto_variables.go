package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunPipelineDtoVariables struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`
}

func (o RunPipelineDtoVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoVariables struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoVariables", string(data)}, " ")
}
