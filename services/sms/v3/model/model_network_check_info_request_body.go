package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkCheckInfoRequestBody struct {

	// 域名连通性
	DomainConnectivity *bool `json:"domain_connectivity,omitempty"`

	// 目的端连通性
	DestinationConnectivity *bool `json:"destination_connectivity,omitempty"`

	// 网络时延
	NetworkDelay float64 `json:"network_delay"`

	// 网络抖动
	NetworkJitter float64 `json:"network_jitter"`

	// 带宽
	MigrationSpeed float64 `json:"migration_speed"`

	// 丢包
	LossPercentage float64 `json:"loss_percentage"`

	// CPU占用
	CpuUsage float64 `json:"cpu_usage"`

	// 内存占用
	MemUsage float64 `json:"mem_usage"`

	// 评估结果
	EvaluationResult string `json:"evaluation_result"`
}

func (o NetworkCheckInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkCheckInfoRequestBody struct{}"
	}

	return strings.Join([]string{"NetworkCheckInfoRequestBody", string(data)}, " ")
}
