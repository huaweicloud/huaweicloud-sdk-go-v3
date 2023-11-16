package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewExtensionOutputs struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o NewExtensionOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionOutputs struct{}"
	}

	return strings.Join([]string{"NewExtensionOutputs", string(data)}, " ")
}
