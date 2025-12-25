package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DpeCompareFunctionDetail 分类映射比较函数详情
type DpeCompareFunctionDetail struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// NORMAL(通用方法), STRING(字符串方法), NUMBER(数字方法), DATE(时间方法)
	Classify *string `json:"classify,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 示例
	Example *string `json:"example,omitempty"`

	// 比较方法Key
	Comparefunc *string `json:"comparefunc,omitempty"`
}

func (o DpeCompareFunctionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DpeCompareFunctionDetail struct{}"
	}

	return strings.Join([]string{"DpeCompareFunctionDetail", string(data)}, " ")
}
