package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StacksAttribute struct {
	// 规格列表

	Specs *[]string `json:"specs,omitempty"`
	// 提示id

	Suggest *string `json:"suggest,omitempty"`
	// 提示信息

	SuggestTitle *string `json:"suggest_title,omitempty"`
	// 卷容量列表

	Volumes *[]string `json:"volumes,omitempty"`
}

func (o StacksAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StacksAttribute struct{}"
	}

	return strings.Join([]string{"StacksAttribute", string(data)}, " ")
}
