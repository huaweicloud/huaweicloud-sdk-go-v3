package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceLtsLogConfigResult struct {

	// **参数解释**: 实例基本信息。
	Instance *interface{} `json:"instance,omitempty"`

	// **参数解释**: LTS相关信息。
	LtsConfigs *[]LtsLogConfigResult `json:"lts_configs,omitempty"`
}

func (o InstanceLtsLogConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceLtsLogConfigResult struct{}"
	}

	return strings.Join([]string{"InstanceLtsLogConfigResult", string(data)}, " ")
}
