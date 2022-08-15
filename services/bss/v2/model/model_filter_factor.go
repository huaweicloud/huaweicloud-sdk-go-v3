package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterFactor struct {

	// |参数名称：维度分组条件| |参数约束及描述：维度分组条件|
	Key *string `json:"key,omitempty"`

	// |参数名称：过滤器值| |参数约束及描述：过滤器值|
	Value *[]string `json:"value,omitempty"`
}

func (o FilterFactor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterFactor struct{}"
	}

	return strings.Join([]string{"FilterFactor", string(data)}, " ")
}
