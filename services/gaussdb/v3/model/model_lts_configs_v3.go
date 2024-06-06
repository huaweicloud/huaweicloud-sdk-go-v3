package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsConfigsV3 struct {

	// LTS配置列表
	LtsConfigs *[]LtsConfig `json:"lts_configs,omitempty"`

	Instance *InstanceLtsBasicInfo `json:"instance,omitempty"`
}

func (o LtsConfigsV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfigsV3 struct{}"
	}

	return strings.Join([]string{"LtsConfigsV3", string(data)}, " ")
}
