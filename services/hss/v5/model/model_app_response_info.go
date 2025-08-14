package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppResponseInfo 软件信息
type AppResponseInfo struct {

	// **参数解释**: Agent ID **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 软件名称 **取值范围**: 字符长度1-256位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 版本号 **取值范围**: 字符长度1-128位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 更新时间，最近一次更新的时间，用毫秒表示 **取值范围**: 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 最近扫描时间，用毫秒表示 **取值范围**: 最小值0，最大值9223372036854775807
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`

	// **参数解释**: 容器ID **取值范围**: 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器实例名称，只有容器类型的告警有 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`
}

func (o AppResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppResponseInfo struct{}"
	}

	return strings.Join([]string{"AppResponseInfo", string(data)}, " ")
}
