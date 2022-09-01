package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 系统标签对象。
type SysTags struct {

	// 标签名。
	Key string `json:"key" xml:"key"`

	// 标签值。
	Value string `json:"value" xml:"value"`
}

func (o SysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTags struct{}"
	}

	return strings.Join([]string{"SysTags", string(data)}, " ")
}
