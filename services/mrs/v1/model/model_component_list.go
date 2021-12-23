package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentList struct {
	// 组件名称

	ComponentName string `json:"component_name"`
}

func (o ComponentList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentList struct{}"
	}

	return strings.Join([]string{"ComponentList", string(data)}, " ")
}
