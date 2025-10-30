package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinuxVulDetailInfo struct {

	// **参数解释**: 漏洞编号 **取值范围**: 字符长度0-255
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: cve漏洞名称 **取值范围**: 字符长度0-512
	CveName *string `json:"cve_name,omitempty"`

	// **参数解释**: 攻击矢量 **取值范围**: 字符长度0-255
	CvssVector *string `json:"cvss_vector,omitempty"`

	// **参数解释**: cve修复建议 **取值范围**: 字符长度0-4096
	CveSolution *string `json:"cve_solution,omitempty"`

	// **参数解释**: cve漏洞危害 **取值范围**: 字符长度0-128
	CveAffect *string `json:"cve_affect,omitempty"`

	// **参数解释**: cve漏洞危害的描述信息 **取值范围**: 字符长度0-4096
	CveAffectDescription *string `json:"cve_affect_description,omitempty"`

	// **参数解释**: cve漏洞类型 **取值范围**: 字符长度0-128
	CveType *string `json:"cve_type,omitempty"`

	// **参数解释**: cve漏洞类型的描述信息 **取值范围**: 字符长度0-4096
	CveTypeDescription *string `json:"cve_type_description,omitempty"`

	// **参数解释**: cve危险程度 **取值范围**: - Low：低危 - Medium：中危 - High：高危
	CveLevel *string `json:"cve_level,omitempty"`

	// **参数解释**: 漏洞分值 **取值范围**: 最小值0，最大值10
	Cvss *float32 `json:"cvss,omitempty"`

	// **参数解释**: cvss评分版本 **取值范围**: 字符长度0-32
	CvssVersion *string `json:"cvss_version,omitempty"`

	// **参数解释**: 漏洞描述 **取值范围**: 字符长度0-1024
	Description *string `json:"description,omitempty"`

	// **参数解释**: 披露时间 **取值范围**: 最小值0，最大值2^63-1
	PublicTime *int64 `json:"public_time,omitempty"`

	// **参数解释**: cnvd编号 **取值范围**: 字符长度0-32
	CnvdId *string `json:"cnvd_id,omitempty"`

	// **参数解释**: cnnvd编号 **取值范围**: 字符长度0-32
	CnnvdId *string `json:"cnnvd_id,omitempty"`

	HostsNum *VulnerabilityHostNumberInfo `json:"hosts_num,omitempty"`
}

func (o LinuxVulDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinuxVulDetailInfo struct{}"
	}

	return strings.Join([]string{"LinuxVulDetailInfo", string(data)}, " ")
}
