package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthCheckConfigurationTcpSocket 健康检查配置中TCP端口检查信息。
type HealthCheckConfigurationTcpSocket struct {

	// 端口。
	Port *int32 `json:"port,omitempty"`
}

func (o HealthCheckConfigurationTcpSocket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCheckConfigurationTcpSocket struct{}"
	}

	return strings.Join([]string{"HealthCheckConfigurationTcpSocket", string(data)}, " ")
}
