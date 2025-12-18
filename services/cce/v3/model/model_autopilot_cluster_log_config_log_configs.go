package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutopilotClusterLogConfigLogConfigs struct {

	// 日志类型
	Name *string `json:"name,omitempty"`

	// 是否采集
	Enable *bool `json:"enable,omitempty"`
}

func (o AutopilotClusterLogConfigLogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutopilotClusterLogConfigLogConfigs struct{}"
	}

	return strings.Join([]string{"AutopilotClusterLogConfigLogConfigs", string(data)}, " ")
}
