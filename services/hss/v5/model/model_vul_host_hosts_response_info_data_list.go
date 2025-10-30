package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulHostHostsResponseInfoDataList struct {

	// **参数解释**: 主机id **取值范围**: 字符长度0-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: agent id **取值范围**: 字符长度0-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度0-64位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 地域 **取值范围**: 字符长度0-32位
	RegionName *string `json:"region_name,omitempty"`

	// **参数解释**: 服务器公网ip **取值范围**: 字符长度0-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器私网ip **取值范围**: 字符长度0-4096位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器组id **取值范围**: 字符长度0-128位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度0-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 操作系统类型 **取值范围**: - Linux ：Linux系统 - Windows ：Windows系统
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 资产重要性 **取值范围**: 字符长度0-32位
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 最近扫描时间 **取值范围**: 最小值0，最大值2^63-1
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 主机风险 **取值范围**: 字符长度0-32位
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 主机风险分数 **取值范围**: 最小值0，最大值100
	Score *int32 `json:"score,omitempty"`

	// **参数解释**: 主机配额 **取值范围**: 字符长度0-64位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 主机配额 **取值范围**: - unhandled：待处理 - handled：已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 各个漏洞修复优先级下的漏洞数量
	VulNumWithRepairPriorityList *[]VulHostHostsResponseInfoVulNumWithRepairPriorityList `json:"vul_num_with_repair_priority_list,omitempty"`

	VulIdsInfo *VulHostHostsResponseInfoVulIdsInfo `json:"vul_ids_info,omitempty"`
}

func (o VulHostHostsResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostHostsResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulHostHostsResponseInfoDataList", string(data)}, " ")
}
