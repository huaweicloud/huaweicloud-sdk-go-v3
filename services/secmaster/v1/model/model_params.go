package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Params 分类映射函数参数信息
type Params struct {

	// 分类映射id
	Key *string `json:"key,omitempty"`

	// 分类映射id
	Value *string `json:"value,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 示例
	Example *string `json:"example,omitempty"`

	// 是否必填
	Mandatory *bool `json:"mandatory,omitempty"`
}

func (o Params) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Params struct{}"
	}

	return strings.Join([]string{"Params", string(data)}, " ")
}
