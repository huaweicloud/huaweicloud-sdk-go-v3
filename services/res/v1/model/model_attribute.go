package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Attribute struct {

	// 名称。
	Name string `json:"name" xml:"name"`

	// 数据类型。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 其他用途。
	OtherUses *[]string `json:"other_uses,omitempty" xml:"other_uses"`
}

func (o Attribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attribute struct{}"
	}

	return strings.Join([]string{"Attribute", string(data)}, " ")
}
