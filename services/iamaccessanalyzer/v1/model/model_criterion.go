package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Criterion 要在查找筛选器中使用的条件。最多只能有一个运算符。
type Criterion struct {

	// 要匹配筛选器的“包含”运算符。
	Contains *[]string `json:"contains,omitempty"`

	// 要匹配筛选器的“等于”运算符。
	Eq *[]string `json:"eq,omitempty"`

	// 要匹配筛选器的“存在”运算符。
	Exists *bool `json:"exists,omitempty"`

	// 要匹配筛选器的“不等于”运算符。
	Neq *[]string `json:"neq,omitempty"`
}

func (o Criterion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Criterion struct{}"
	}

	return strings.Join([]string{"Criterion", string(data)}, " ")
}
