package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostVulInfo struct {

	// **参数解释**: 漏洞名称 **取值范围**: 字符范围0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞ID **取值范围**: 字符范围0-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞标签列表 **取值范围**: 最小值0，最大值2147483647
	LabelList *[]string `json:"label_list,omitempty"`

	// **参数解释**: 修复紧急度 **取值范围**: - immediate_repair  : 尽快修复 - delay_repair      : 延后修复 - not_needed_repair : 暂可不修复
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 最近扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 漏洞类型 **取值范围**: - linux_vul   : linux漏洞 - windows_vul : windows漏洞 - web_cms     : Web-CMS漏洞 - app_vul     : 应用漏洞 - urgent_vul  : 应急漏洞
	Type *string `json:"type,omitempty"`

	// **参数解释**: 服务器上受该漏洞影响的软件列表 **取值范围**: 最小值0，最大值2147483647
	AppList *[]HostVulInfoAppList `json:"app_list,omitempty"`

	// **参数解释**: 危险程度 **取值范围**: - Critical : 漏洞cvss评分大于等于9；对应控制台页面的高危 - High     : 漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium   : 漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low      : 漏洞cvss评分小于4；对应控制台页面的低危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 解决方案 **取值范围**: 字符范围0-65534位
	SolutionDetail *string `json:"solution_detail,omitempty"`

	// **参数解释**: URL链接 **取值范围**: 字符范围0-2083位
	Url *string `json:"url,omitempty"`

	// **参数解释**: 漏洞描述 **取值范围**: 字符范围0-65534位
	Description *string `json:"description,omitempty"`

	// **参数解释**: 修复命令行 **取值范围**: 字符范围1-256位
	RepairCmd *string `json:"repair_cmd,omitempty"`

	// **参数解释**: 漏洞状态 **取值范围**: - vul_status_unfix            : 未处理 - vul_status_ignored          : 已忽略 - vul_status_verified         : 验证中 - vul_status_fixing           : 修复中 - vul_status_fixed            : 修复成功 - vul_status_reboot           : 修复成功待重启 - vul_status_failed           : 修复失败 - vul_status_fix_after_reboot : 请重启主机再次修复
	Status *string `json:"status,omitempty"`

	// **参数解释**: HSS全网修复该漏洞的次数 **取值范围**: 最小值0，最大值1000000
	RepairSuccessNum *int32 `json:"repair_success_num,omitempty"`

	// **参数解释**: CVE列表 **取值范围**: 最小值1，最大值10000
	CveList *[]HostVulInfoCveList `json:"cve_list,omitempty"`

	// **参数解释**: 是否影响业务 **取值范围**: - true  : 影响业务 - false : 不影响业务
	IsAffectBusiness *bool `json:"is_affect_business,omitempty"`

	// **参数解释**: 首次扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// **参数解释**: 软件名称 **取值范围**: 字符长度0-256位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符长度0-256位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 软件路径 **取值范围**: 字符长度0-512位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 主机配额 **取值范围**: 字符长度0-128位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 是否可以回滚到修复漏洞时创建的备份 **取值范围**: - true  : 可以回滚 - false : 不可以回滚
	SupportRestore *bool `json:"support_restore,omitempty"`

	// **参数解释**: 该漏洞不可进行的操作类型列表 **取值范围**: 最小值1，最大值10000
	DisabledOperateTypes *[]HostVulInfoDisabledOperateTypes `json:"disabled_operate_types,omitempty"`

	// **参数解释**: 修复优先级 **取值范围**: - Critical : 紧急  - High     : 高  - Medium   : 中  - Low      : 低
	RepairPriority *string `json:"repair_priority,omitempty"`
}

func (o HostVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulInfo struct{}"
	}

	return strings.Join([]string{"HostVulInfo", string(data)}, " ")
}
