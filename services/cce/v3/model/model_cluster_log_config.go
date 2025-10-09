package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterLogConfig struct {

	// **参数解释**： 日志存储时长，单位（天） **取值范围**： 0-30
	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	// **参数解释**： 日志配置项详细信息 **约束限制**: 不涉及
	LogConfigs *[]ClusterLogConfigLogConfigs `json:"log_configs,omitempty"`
}

func (o ClusterLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterLogConfig struct{}"
	}

	return strings.Join([]string{"ClusterLogConfig", string(data)}, " ")
}
