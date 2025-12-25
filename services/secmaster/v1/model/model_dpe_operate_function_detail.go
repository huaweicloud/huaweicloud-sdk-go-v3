package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DpeOperateFunctionDetail 分类映射比较函数信息
type DpeOperateFunctionDetail struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// NORMAL(通用方法), STRING(字符串方法), NUMBER(数字方法), DATE(时间方法)
	Classify *string `json:"classify,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 示例
	Example *string `json:"example,omitempty"`

	// 转换方法Key
	Operationfunc *string `json:"operationfunc,omitempty"`

	// 参数集合
	Params *[]Params `json:"params,omitempty"`
}

func (o DpeOperateFunctionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DpeOperateFunctionDetail struct{}"
	}

	return strings.Join([]string{"DpeOperateFunctionDetail", string(data)}, " ")
}
