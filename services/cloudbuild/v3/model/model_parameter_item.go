package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParameterItem struct {
	// 参数名称

	Name string `json:"name"`
	// 参数值

	Value string `json:"value"`
}

func (o ParameterItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterItem struct{}"
	}

	return strings.Join([]string{"ParameterItem", string(data)}, " ")
}
