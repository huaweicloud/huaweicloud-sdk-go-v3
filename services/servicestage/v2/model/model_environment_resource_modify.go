package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvironmentResourceModify struct {
	// 添加基础资源。

	AddBaseResources *[]Resource `json:"add_base_resources,omitempty"`
	// 添加其他资源。

	AddOptionalResources *[]Resource `json:"add_optional_resources,omitempty"`
	// 移除资源。

	RemoveResources *[]Resource `json:"remove_resources,omitempty"`
}

func (o EnvironmentResourceModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentResourceModify struct{}"
	}

	return strings.Join([]string{"EnvironmentResourceModify", string(data)}, " ")
}
