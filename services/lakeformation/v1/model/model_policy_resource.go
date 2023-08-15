package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyResource struct {

	// 是否排除
	IsExcludes *bool `json:"is_excludes,omitempty"`

	// 是否递归
	IsRecursive *bool `json:"is_recursive,omitempty"`

	// 值
	Values *[]string `json:"values,omitempty"`
}

func (o PolicyResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyResource struct{}"
	}

	return strings.Join([]string{"PolicyResource", string(data)}, " ")
}
