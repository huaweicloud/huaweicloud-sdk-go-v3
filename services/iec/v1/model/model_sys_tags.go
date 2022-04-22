package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 系统标签。
type SysTags struct {

	// 系统标签的Key值。
	Key *string `json:"key,omitempty"`

	// 系统标签的value值。
	Value *string `json:"value,omitempty"`
}

func (o SysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTags struct{}"
	}

	return strings.Join([]string{"SysTags", string(data)}, " ")
}
