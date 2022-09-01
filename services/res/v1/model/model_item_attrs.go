package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ItemAttrs struct {

	// 数据类型。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 物品。
	Name *string `json:"name,omitempty" xml:"name"`

	// 其他用途。
	OtherUses *[]string `json:"other_uses,omitempty" xml:"other_uses"`
}

func (o ItemAttrs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemAttrs struct{}"
	}

	return strings.Join([]string{"ItemAttrs", string(data)}, " ")
}
