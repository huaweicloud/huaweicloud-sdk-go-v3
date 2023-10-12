package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthCheckConfigurationLivenessProbe 组件健康检查配置的存活探针配置。
type HealthCheckConfigurationLivenessProbe struct {

	// 检测周期，单位为s。
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`

	// 延迟时间，单位为s。
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`

	// 超时时间，单位为s。
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`

	// 成功阈值。
	SuccessThreshold *int32 `json:"successThreshold,omitempty"`

	// 最大失败次数。
	FailureThreshold *int32 `json:"failureThreshold,omitempty"`

	HttpGet *HealthCheckConfigurationHttpGet `json:"httpGet,omitempty"`

	TcpSocket *HealthCheckConfigurationTcpSocket `json:"tcpSocket,omitempty"`

	Exec *HealthCheckConfigurationExec `json:"exec,omitempty"`
}

func (o HealthCheckConfigurationLivenessProbe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCheckConfigurationLivenessProbe struct{}"
	}

	return strings.Join([]string{"HealthCheckConfigurationLivenessProbe", string(data)}, " ")
}
