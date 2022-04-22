package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StacksTags struct {

	// 技术栈列表
	StackList *[]StackInfo `json:"stack_list,omitempty"`

	// 技术栈tag集合
	Tags *[]string `json:"tags,omitempty"`
}

func (o StacksTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StacksTags struct{}"
	}

	return strings.Join([]string{"StacksTags", string(data)}, " ")
}
