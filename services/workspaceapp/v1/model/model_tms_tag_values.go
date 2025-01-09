package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTagValues tms标签
type TmsTagValues struct {

	// 键，最大长度128个unicode字符。
	Key *string `json:"key,omitempty"`

	// 值列表。
	Values *[]string `json:"values,omitempty"`
}

func (o TmsTagValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagValues struct{}"
	}

	return strings.Join([]string{"TmsTagValues", string(data)}, " ")
}
