package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulHostInfo **参数解释**: 软件漏洞列表
type VulHostInfo struct {

	// 受漏洞影响的服务器id
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 危险程度 **取值范围**: - Critical : 漏洞cvss评分大于等于9；对应控制台页面的高危 - High     : 漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium   : 漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low      : 漏洞cvss评分小于4；对应控制台页面的低危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 受影响主机名称 **取值范围**: 字符范围1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 受影响主机ip **取值范围**: 字符范围1-256位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 主机对应的agent id **取值范围**: 字符范围1-128位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 主机绑定的配额版本 **取值范围**: 字符范围1-128位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 漏洞cve总数 **取值范围**: 最小值0，最大值10000
	CveNum *int32 `json:"cve_num,omitempty"`

	// **参数解释**: 漏洞对应的cve id列表 **取值范围**: 最小值1，最大值10000
	CveIdList *[]string `json:"cve_id_list,omitempty"`

	// **参数解释**: 漏洞状态 **取值范围**: - vul_status_unfix            : 未处理 - vul_status_ignored          : 已忽略 - vul_status_verified         : 验证中 - vul_status_fixing           : 修复中 - vul_status_fixed            : 修复成功 - vul_status_reboot           : 修复成功待重启 - vul_status_failed           : 修复失败 - vul_status_fix_after_reboot : 请重启主机再次修复
	Status *string `json:"status,omitempty"`

	// **参数解释**: 修复漏洞需要执行的命令行（只有Linux漏洞有该字段） **取值范围**: 字符范围1-256位
	RepairCmd *string `json:"repair_cmd,omitempty"`

	// **参数解释**: 应用软件的路径（只有应用漏洞有该字段） **取值范围**: 字符范围1-512位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 地域 **取值范围**: 字符范围0-128位
	RegionName *string `json:"region_name,omitempty"`

	// **参数解释**: 服务器公网ip **取值范围**: 字符范围0-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器私网ip **取值范围**: 字符范围0-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器组id **取值范围**: 字符范围0-128位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符范围0-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 操作系统 **取值范围**: - Linux ：linux系统 - Windows：windows系统
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 资产重要性 **取值范围**: - important : 重要资产 - common    : 一般资产 - test      : 测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 是否影响业务 **取值范围**: - true  : 影响业务 - false : 不影响业务
	IsAffectBusiness *bool `json:"is_affect_business,omitempty"`

	// **参数解释**: 首次扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// **参数解释**: 扫描时间，时间戳单位：毫秒 **取值范围**: 最小值0，最大值9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 是否可以回滚到修复漏洞时创建的备份 **取值范围**: - true  : 可以回滚 - false : 不可以回滚
	SupportRestore *bool `json:"support_restore,omitempty"`

	// **参数解释**: 漏洞在当前主机上不可进行的操作类型列表 **取值范围**: 最小值1，最大值10000
	DisabledOperateTypes *[]HostVulInfoDisabledOperateTypes `json:"disabled_operate_types,omitempty"`

	// **参数解释**: 修复优先级 **取值范围**: - Critical : 紧急 - High     : 高 - Medium   : 中 - Low      : 低
	RepairPriority *string `json:"repair_priority,omitempty"`
}

func (o VulHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostInfo struct{}"
	}

	return strings.Join([]string{"VulHostInfo", string(data)}, " ")
}
