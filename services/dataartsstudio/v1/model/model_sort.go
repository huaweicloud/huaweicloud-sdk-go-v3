package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Sort 排序条件
type Sort struct {

	// 属性
	Attribute *string `json:"attribute,omitempty"`

	// 排序枚举值，取值范围DESC、ASC， 默认值ASC
	Order *string `json:"order,omitempty"`
}

func (o Sort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sort struct{}"
	}

	return strings.Join([]string{"Sort", string(data)}, " ")
}
