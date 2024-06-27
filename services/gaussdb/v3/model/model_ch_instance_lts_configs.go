package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChInstanceLtsConfigs struct {

	// 实例LTS配置信息。
	LtsConfigs *[]ChLtsConfigs `json:"lts_configs,omitempty"`

	Instance *ChInstanceLtsConfigsInstance `json:"instance,omitempty"`
}

func (o ChInstanceLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstanceLtsConfigs struct{}"
	}

	return strings.Join([]string{"ChInstanceLtsConfigs", string(data)}, " ")
}
