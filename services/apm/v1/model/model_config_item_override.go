package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigItemOverride 实际生效值
type ConfigItemOverride struct {

	// 环境标签ID
	EnvTagId *int64 `json:"env_tag_id,omitempty"`

	// 环境标签名
	EnvTagName *string `json:"env_tag_name,omitempty"`

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ConfigItemOverride) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigItemOverride struct{}"
	}

	return strings.Join([]string{"ConfigItemOverride", string(data)}, " ")
}
