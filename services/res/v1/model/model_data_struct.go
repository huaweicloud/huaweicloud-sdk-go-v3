package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataStruct struct {
	Behaviors *BehaviorsConfig `json:"behaviors,omitempty"`

	// 物品参数。
	ItemAttrs *[]ItemAttrs `json:"item_attrs,omitempty"`

	// 用户参数。
	UserAttrs *[]UserAttrs `json:"user_attrs,omitempty"`

	UserDynamicAttr *UserDynamicAttr `json:"user_dynamic_attr,omitempty"`
}

func (o DataStruct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStruct struct{}"
	}

	return strings.Join([]string{"DataStruct", string(data)}, " ")
}
