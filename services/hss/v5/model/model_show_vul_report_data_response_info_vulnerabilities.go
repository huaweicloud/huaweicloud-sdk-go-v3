package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVulReportDataResponseInfoVulnerabilities struct {

	// **参数解释**： 漏洞名称 **取值范围**： 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 最近扫描时间 **取值范围**： 最小值0，最大值9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 漏洞类型 **取值范围**： - linux_vul：linux漏洞。 - windows_vul：windows漏洞。 - web_cms：Web-CMS漏洞。 - app_vul：应用漏洞。
	Type *string `json:"type,omitempty"`

	// CVE列表
	CveList *[]ShowVulReportDataResponseInfoCveList `json:"cve_list,omitempty"`

	// **参数解释**： 修复优先级 **取值范围**： - Critical：紧急。 - High：高。 - Medium：中。 - Low：低。
	RepairPriority *string `json:"repair_priority,omitempty"`

	// **参数解释**： 影响主机数量 **取值范围**： 最小值0，最大值20000
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 主机ID列表(数组长度跟host_num对不上时，主机数量以host_num为准) **取值范围**: 最小值0，最大值20000
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ShowVulReportDataResponseInfoVulnerabilities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulReportDataResponseInfoVulnerabilities struct{}"
	}

	return strings.Join([]string{"ShowVulReportDataResponseInfoVulnerabilities", string(data)}, " ")
}
