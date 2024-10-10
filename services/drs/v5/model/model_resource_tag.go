package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 标签信息体。
type ResourceTag struct {

	// 标签键。 约束：最大长度36，只能包含字母、数字、下划线、中划线和中文。
	Key *string `json:"key,omitempty"`

	// 标签值。标签的值可以包含任意语种字母、数字、空格和_ . : / = + - @。
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
