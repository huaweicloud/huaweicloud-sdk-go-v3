package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTag tms标签
type TmsTag struct {

	// 键。最大长度128个unicode字符。
	Key string `json:"key"`

	// 值。每个值最大长度255个unicode字符。
	Value *string `json:"value,omitempty"`
}

func (o TmsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTag struct{}"
	}

	return strings.Join([]string{"TmsTag", string(data)}, " ")
}
