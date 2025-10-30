package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UrgentVulInfo 应急漏洞信息
type UrgentVulInfo struct {

	// **参数解释**： 漏洞名称 **取值范围**： 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 漏洞ID **取值范围**： 字符长度0-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**： 漏洞标签 **取值范围**: 最小值0，最大值2147483647
	LabelList *[]string `json:"label_list,omitempty"`

	HostsNum *VulnerabilityHostNumberInfo `json:"hosts_num,omitempty"`

	// **参数解释**： 最近扫描时间 **取值范围**： 字符长度0-9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 漏洞披露时间 **取值范围**： 字符长度0-9223372036854775807
	PublishTime *int64 `json:"publish_time,omitempty"`

	// **参数解释**： 解决方案 **取值范围**： 字符长度0-65534位
	SolutionDetail *string `json:"solution_detail,omitempty"`

	// **参数解释**： 漏洞描述 **取值范围**： 字符长度0-64位
	Description *string `json:"description,omitempty"`

	// **参数解释**： 漏洞扫描状态 **约束限制**: 不涉及 **取值范围**： - never_scan : 未扫描 - scanning : 扫描中 - finished : 扫描完成  **默认取值**: 不涉及
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**： 危险程度 **约束限制**: 不涉及 **取值范围**： - Critical：漏洞cvss评分大于等于9；对应控制台页面的高危 - High：漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium：漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low：漏洞cvss评分小于4；对应控制台页面的低危  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**： 处于扫描中状态的主机数量 **取值范围**： 字符长度0-2147483647
	ScanningHostNum *int32 `json:"scanning_host_num,omitempty"`

	// **参数解释**： 已扫描成功的主机数量 **取值范围**： 字符长度0-2147483647
	SuccessHostNum *int32 `json:"success_host_num,omitempty"`

	// **参数解释**： 已扫描失败的主机数量 **取值范围**： 字符长度0-2147483647
	FailedHostNum *int32 `json:"failed_host_num,omitempty"`
}

func (o UrgentVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrgentVulInfo struct{}"
	}

	return strings.Join([]string{"UrgentVulInfo", string(data)}, " ")
}
