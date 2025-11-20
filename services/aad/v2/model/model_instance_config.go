package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceConfig struct {

	// 实例级别标签
	Tags *[]string `json:"tags,omitempty"`

	// 冻结类型
	FreezeType *[]int32 `json:"freeze_type,omitempty"`
}

func (o InstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceConfig struct{}"
	}

	return strings.Join([]string{"InstanceConfig", string(data)}, " ")
}
