package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceLogConfig struct {
	Instance *LogInstanceInfo `json:"instance"`

	// LTS日志配置明细。从未设置LTS日志流，不返回该字段。
	LtsConfigs []InstanceLogConfigDetail `json:"lts_configs"`
}

func (o InstanceLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceLogConfig struct{}"
	}

	return strings.Join([]string{"InstanceLogConfig", string(data)}, " ")
}
