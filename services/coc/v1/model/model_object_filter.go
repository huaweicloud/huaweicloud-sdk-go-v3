package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectFilter struct {

	// 操作符 in/like/startwith/endwith/=/!=/>/<等
	Operator string `json:"operator"`

	// 字段名称
	Field string `json:"field"`

	// 搜索值
	Values []string `json:"values"`
}

func (o ObjectFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectFilter struct{}"
	}

	return strings.Join([]string{"ObjectFilter", string(data)}, " ")
}
