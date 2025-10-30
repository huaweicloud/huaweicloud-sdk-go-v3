package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WindowsVulDetailInfo struct {

	// **参数解释**: CVE ID **取值范围**: 字符范围1-32位
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: cve漏洞名称 **取值范围**: 字符范围0-512位
	CveName *string `json:"cve_name,omitempty"`

	// **参数解释**: cvss评分版本 **取值范围**: 字符范围0-32位
	CvssVersion *string `json:"cvss_version,omitempty"`

	// **参数解释**: 攻击矢量 **取值范围**: 字符范围0-128位
	CvssVector *string `json:"cvss_vector,omitempty"`

	// **参数解释**: cve修复建议 **取值范围**: 字符范围0-128位
	CveSolution *string `json:"cve_solution,omitempty"`

	// **参数解释**: cve漏洞危害 **取值范围**: 字符范围0-128位
	CveAffect *string `json:"cve_affect,omitempty"`

	// **参数解释**: cve漏洞危害的描述信息 **取值范围**: 字符范围0-4096位
	CveAffectDescription *string `json:"cve_affect_description,omitempty"`

	// **参数解释**: cve漏洞类型 **取值范围**: 字符范围0-128位
	CveType *string `json:"cve_type,omitempty"`

	// **参数解释**: cve漏洞类型的描述信息 **取值范围**: 字符范围0-4096位
	CveTypeDescription *string `json:"cve_type_description,omitempty"`

	// **参数解释**: 危险程度 **取值范围**: - Low：低危险程度 - Medium：中危险程度 - High：高危险程度
	CveLevel *string `json:"cve_level,omitempty"`

	// **参数解释**: cvss评分 **取值范围**: 最小值0，最大值10
	Cvss *float32 `json:"cvss,omitempty"`

	// **参数解释**: 漏洞披露时间 **取值范围**: 最小值0，最大值9223372036854775807
	PublicTime *int64 `json:"public_time,omitempty"`

	// **参数解释**: cve漏洞描述 **取值范围**: 字符范围0-65535位
	Description *string `json:"description,omitempty"`

	HostsNum *VulnerabilityHostNumberInfo `json:"hosts_num,omitempty"`
}

func (o WindowsVulDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WindowsVulDetailInfo struct{}"
	}

	return strings.Join([]string{"WindowsVulDetailInfo", string(data)}, " ")
}
