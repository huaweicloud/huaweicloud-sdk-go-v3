package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttrValueRules
type AttrValueRules struct {

	// 被推荐对象的属性-值配置。
	AttrValuesA *[]AttrValue `json:"attr_values_a,omitempty"`

	// 待推荐对象的属性-值配置。
	AttrValuesB []AttrValue `json:"attr_values_b"`
}

func (o AttrValueRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttrValueRules struct{}"
	}

	return strings.Join([]string{"AttrValueRules", string(data)}, " ")
}
