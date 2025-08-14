package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulInfo 漏洞列表
type VulInfo struct {

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞ID **取值范围**: 字符长度0-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞标签列表 **取值范围**: 不涉及
	LabelList *[]string `json:"label_list,omitempty"`

	// **参数解释**: 漏洞修复的必要性 **取值范围**: - Critical：漏洞cvss评分大于等于9；对应控制台页面的高危 - High：漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium：漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low：漏洞cvss评分小于4；对应控制台页面的低危
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 漏洞风险级别 **取值范围**: - Critical：漏洞cvss评分大于等于9；对应控制台页面的高危 - High：漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium：漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low：漏洞cvss评分小于4；对应控制台页面的低危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 受影响服务器台数 **取值范围**: 取值0-2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 未处理主机台数：除已忽略和已修复的主机数量 **取值范围**: 取值0-2147483647
	UnhandleHostNum *int32 `json:"unhandle_host_num,omitempty"`

	// **参数解释**: 最近扫描时间，时间戳单位：毫秒 **取值范围**: 取值0-9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 修复漏洞的指导意见 **取值范围**: 字符长度0-65534位
	SolutionDetail *string `json:"solution_detail,omitempty"`

	// **参数解释**: URL链接 **取值范围**: 字符长度0-2083位
	Url *string `json:"url,omitempty"`

	// **参数解释**: 漏洞描述 **取值范围**: 字符长度0-65534位
	Description *string `json:"description,omitempty"`

	// **参数解释**: 漏洞类型 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞
	Type *string `json:"type,omitempty"`

	// **参数解释**: 可处置该漏洞的主机ID列表 **取值范围**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释**: 漏洞关联的cve信息列表 **取值范围**: 不涉及
	CveList *[]VulInfoCveList `json:"cve_list,omitempty"`

	// **参数解释**: 补丁地址 **取值范围**: 字符长度0-512位
	PatchUrl *string `json:"patch_url,omitempty"`

	// **参数解释**: 修复优先级 **取值范围**: - Critical：紧急 - High：高 - Medium：中 - Low：低
	RepairPriority *string `json:"repair_priority,omitempty"`

	HostsNum *VulnerabilityHostNumberInfo `json:"hosts_num,omitempty"`

	// **参数解释**: 修复成功次数 **取值范围**: 取值0-1000000
	RepairSuccessNum *int32 `json:"repair_success_num,omitempty"`

	// **参数解释**: 修复数量 **取值范围**: 取值0-1000000
	FixedNum *int64 `json:"fixed_num,omitempty"`

	// **参数解释**: 忽略数量 **取值范围**: 取值0-1000000
	IgnoredNum *int64 `json:"ignored_num,omitempty"`

	// **参数解释**: 验证数量 **取值范围**: 取值0-1000000
	VerifyNum *int32 `json:"verify_num,omitempty"`

	// **参数解释**: 修复优先级，每个修复优先级对应的主机数量 **取值范围**: 不涉及
	RepairPriorityList *[]RepairPriorityListInfo `json:"repair_priority_list,omitempty"`
}

func (o VulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulInfo struct{}"
	}

	return strings.Join([]string{"VulInfo", string(data)}, " ")
}
