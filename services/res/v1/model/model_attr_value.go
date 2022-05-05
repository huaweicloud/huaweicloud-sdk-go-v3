package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AttrValue struct {

	// 属性名称。
	Name string `json:"name"`

	// 属性值。
	Value string `json:"value"`
}

func (o AttrValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttrValue struct{}"
	}

	return strings.Join([]string{"AttrValue", string(data)}, " ")
}
