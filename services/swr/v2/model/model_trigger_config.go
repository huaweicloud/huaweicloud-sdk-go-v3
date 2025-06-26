package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TriggerConfig struct {

	// 触发类型，镜像签名、老化规则只支持manual(手动)、scheduled(定时+手动)；同步策略支持manual(手动)、scheduled(定时+手动)、event_based(事件触发+手动)
	Type string `json:"type"`

	TriggerSettings *TriggerSetting `json:"trigger_settings,omitempty"`
}

func (o TriggerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerConfig struct{}"
	}

	return strings.Join([]string{"TriggerConfig", string(data)}, " ")
}
