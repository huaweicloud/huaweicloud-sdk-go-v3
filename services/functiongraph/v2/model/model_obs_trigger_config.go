package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsTriggerConfig 触发器结构体
type ObsTriggerConfig struct {

	// OBS桶名（trigger_type为OBS时配置）
	Bucket string `json:"bucket"`

	// OBS事件列表（trigger_type为OBS时配置）
	Events []string `json:"events"`

	// 对象名前缀（trigger_type为OBS时配置）
	Prefix string `json:"prefix"`

	// 对象名后缀（trigger_type为OBS时配置）
	Suffix string `json:"suffix"`
}

func (o ObsTriggerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsTriggerConfig struct{}"
	}

	return strings.Join([]string{"ObsTriggerConfig", string(data)}, " ")
}
