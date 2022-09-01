package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点类型详细
type Detail struct {

	// 属性类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 属性值。
	Value string `json:"value" xml:"value"`

	// 属性单位。
	Unit string `json:"unit" xml:"unit"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}
