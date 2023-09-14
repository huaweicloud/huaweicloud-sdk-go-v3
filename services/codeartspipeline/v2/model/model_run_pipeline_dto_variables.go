package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunPipelineDtoVariables struct {

	// 参数名
	Name string `json:"name"`

	// 参数值
	Value string `json:"value"`
}

func (o RunPipelineDtoVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoVariables struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoVariables", string(data)}, " ")
}
