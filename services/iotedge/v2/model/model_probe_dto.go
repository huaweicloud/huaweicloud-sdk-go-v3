package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProbeDto struct {

	// 执行探测的命令行命令
	ExecCommand *string `json:"exec_command,omitempty"`

	TcpSocket *TcpSocketDto `json:"tcp_socket,omitempty"`

	HttpGet *HttpGetDto `json:"http_get,omitempty"`

	// 表示从工作负载启动后从多久开始探测
	InitialDelaySeconds int32 `json:"initial_delay_seconds"`

	// 表示探测超时时间
	TimeoutSeconds int32 `json:"timeout_seconds"`

	// 检查周期
	PeriodSeconds *int32 `json:"period_seconds,omitempty"`

	// 失败多少次算不健康
	FailureThreshold *int32 `json:"failure_threshold,omitempty"`
}

func (o ProbeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProbeDto struct{}"
	}

	return strings.Join([]string{"ProbeDto", string(data)}, " ")
}
