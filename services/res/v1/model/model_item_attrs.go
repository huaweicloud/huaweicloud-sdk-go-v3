package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ItemAttrs struct {

	// 数据类型。
	DataType *string `json:"data_type,omitempty"`

	// 物品。
	Name *string `json:"name,omitempty"`

	// 其他用途。
	OtherUses *[]string `json:"other_uses,omitempty"`
}

func (o ItemAttrs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemAttrs struct{}"
	}

	return strings.Join([]string{"ItemAttrs", string(data)}, " ")
}
