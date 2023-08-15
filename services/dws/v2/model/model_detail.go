package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Detail 节点类型详细
type Detail struct {

	// 属性类型。
	Type *string `json:"type,omitempty"`

	// 属性值。
	Value string `json:"value"`

	// 属性单位。
	Unit string `json:"unit"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}
