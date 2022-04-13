package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点类型详细
type Detail struct {
	// 属性单位

	Unit string `json:"unit"`
	// 属性类型

	Type *string `json:"type,omitempty"`
	// 属性值

	Value string `json:"value"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}
