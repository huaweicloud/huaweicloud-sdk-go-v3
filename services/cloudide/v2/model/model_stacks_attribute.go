package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StacksAttribute struct {

	// 规格列表
	Specs *[]string `json:"specs,omitempty" xml:"specs"`

	// 提示id
	Suggest *string `json:"suggest,omitempty" xml:"suggest"`

	// 提示信息
	SuggestTitle *string `json:"suggest_title,omitempty" xml:"suggest_title"`

	// 卷容量列表
	Volumes *[]string `json:"volumes,omitempty" xml:"volumes"`
}

func (o StacksAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StacksAttribute struct{}"
	}

	return strings.Join([]string{"StacksAttribute", string(data)}, " ")
}
