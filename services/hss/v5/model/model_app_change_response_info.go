package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppChangeResponseInfo 软件变动历史信息
type AppChangeResponseInfo struct {

	// **参数解释**: Agent ID **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 变更类型 **取值范围**: - add：新建 - delete：删除 - modify：修改
	VariationType *string `json:"variation_type,omitempty"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 软件名称 **取值范围**: 字符长度1-256位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 版本号 **取值范围**: 字符长度1-128位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 软件更新时间，单位毫秒 **取值范围**: 最小值0，最大值10000
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 最近扫描时间，单位毫秒 **取值范围**: 最小值0，最大值10000
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AppChangeResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppChangeResponseInfo struct{}"
	}

	return strings.Join([]string{"AppChangeResponseInfo", string(data)}, " ")
}
