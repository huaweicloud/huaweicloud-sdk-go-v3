package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterFactor struct {

	// 维度分组条件
	Key *string `json:"key,omitempty"`

	// 过滤器值
	Value *[]string `json:"value,omitempty"`
}

func (o FilterFactor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterFactor struct{}"
	}

	return strings.Join([]string{"FilterFactor", string(data)}, " ")
}
