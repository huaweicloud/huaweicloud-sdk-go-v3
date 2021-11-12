package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StacksTag struct {
	// 技术栈列表

	StackList *[]Stacks `json:"stack_list,omitempty"`
	// 技术栈tag集合

	Tags *[]string `json:"tags,omitempty"`
}

func (o StacksTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StacksTag struct{}"
	}

	return strings.Join([]string{"StacksTag", string(data)}, " ")
}
