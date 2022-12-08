package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对于多个key相同的value，聚合成为key/values
type AggTag struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 相同键的值列表
	Values *[]string `json:"values,omitempty"`
}

func (o AggTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggTag struct{}"
	}

	return strings.Join([]string{"AggTag", string(data)}, " ")
}
