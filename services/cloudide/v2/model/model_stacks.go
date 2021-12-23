package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Stacks struct {
	ArmConfig *StacksConfig `json:"arm_config,omitempty"`

	Config *StacksConfig `json:"config,omitempty"`
	// 创建人

	Creator *string `json:"creator,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 是否可用

	Disable *bool `json:"disable,omitempty"`
	// id

	Id *int64 `json:"id,omitempty"`
	// 标签

	Label *string `json:"label,omitempty"`
	// 图标

	Logo *string `json:"logo,omitempty"`
	// 技术栈名称

	Name *string `json:"name,omitempty"`
	// 范围

	Scope *string `json:"scope,omitempty"`
	// 技术栈ID，通过技术栈管理ListStacksByTag接口获取。

	StackId *string `json:"stack_id,omitempty"`
	// tags

	Tags *[]string `json:"tags,omitempty"`
}

func (o Stacks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stacks struct{}"
	}

	return strings.Join([]string{"Stacks", string(data)}, " ")
}
