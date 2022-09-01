package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataStruct struct {
	Behaviors *BehaviorsConfig `json:"behaviors,omitempty" xml:"behaviors"`

	// 物品参数。
	ItemAttrs *[]ItemAttrs `json:"item_attrs,omitempty" xml:"item_attrs"`

	// 用户参数。
	UserAttrs *[]UserAttrs `json:"user_attrs,omitempty" xml:"user_attrs"`

	UserDynamicAttr *UserDynamicAttr `json:"user_dynamic_attr,omitempty" xml:"user_dynamic_attr"`
}

func (o DataStruct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStruct struct{}"
	}

	return strings.Join([]string{"DataStruct", string(data)}, " ")
}
