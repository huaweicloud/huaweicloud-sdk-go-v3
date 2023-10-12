package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthCheckConfigurationExec 健康检查配置中执行命令检查。
type HealthCheckConfigurationExec struct {

	// shell语句。
	Command *[]interface{} `json:"command,omitempty"`
}

func (o HealthCheckConfigurationExec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCheckConfigurationExec struct{}"
	}

	return strings.Join([]string{"HealthCheckConfigurationExec", string(data)}, " ")
}
