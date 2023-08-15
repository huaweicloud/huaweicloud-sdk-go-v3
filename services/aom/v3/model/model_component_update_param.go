package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentUpdateParam struct {

	// 组件描述
	Description *string `json:"description,omitempty"`

	// 组件名称
	Name string `json:"name"`
}

func (o ComponentUpdateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentUpdateParam struct{}"
	}

	return strings.Join([]string{"ComponentUpdateParam", string(data)}, " ")
}
