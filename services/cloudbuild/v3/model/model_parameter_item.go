package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParameterItem struct {

	// 参数名称
	Name string `json:"name" xml:"name"`

	// 参数值
	Value string `json:"value" xml:"value"`
}

func (o ParameterItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterItem struct{}"
	}

	return strings.Join([]string{"ParameterItem", string(data)}, " ")
}
