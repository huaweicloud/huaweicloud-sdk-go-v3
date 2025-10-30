package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebCmsVulDetailInfo struct {

	// **参数解释**: 漏洞补丁编号 **取值范围**: 字符长度0-256
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 软件名称 **取值范围**: 字符长度0-32
	App *string `json:"app,omitempty"`

	// **参数解释**: 中文名称 **取值范围**: 字符长度0-128
	NameZh *string `json:"name_zh,omitempty"`

	// **参数解释**: 英文名称 **取值范围**: 字符长度0-128
	NameEn *string `json:"name_en,omitempty"`

	// **参数解释**: 披露时间 **取值范围**: 最小值0，最大值2^63-1
	PublicTime *int64 `json:"public_time,omitempty"`

	// **参数解释**: 漏洞标签中文名称 **取值范围**: 字符长度0-64
	VulLabelZh *string `json:"vulLabel_zh,omitempty"`

	// **参数解释**: 漏洞标签英文名称 **取值范围**: 字符长度0-64
	VulLabelEn *string `json:"vulLabel_en,omitempty"`

	// **参数解释**: 修复必要性 **取值范围**: - 1：高 - 2：中 - 3：低
	RepairNecessity *int32 `json:"repair_necessity,omitempty"`

	// **参数解释**: 修复必要性 **取值范围**: - Low：低危 - Medium：中危 - High：高危 - Critical：紧急
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: cve漏洞中文描述 **取值范围**: 字符长度0-1024
	DescriptionZh *string `json:"description_zh,omitempty"`

	// **参数解释**: cve漏洞英文描述 **取值范围**: 字符长度0-1024
	DescriptionEn *string `json:"description_en,omitempty"`

	// **参数解释**: cve漏洞中文修复建议 **取值范围**: 字符长度0-1024
	SolutionZh *string `json:"solution_zh,omitempty"`

	// **参数解释**: cve漏洞英文修复建议 **取值范围**: 字符长度0-1024
	SolutionEn *string `json:"solution_en,omitempty"`

	// **参数解释**: 漏洞编号 **取值范围**: 字符长度0-255
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: cve分数 **取值范围**: 最小值0，最大值10
	CveScore *float32 `json:"cve_score,omitempty"`

	// **参数解释**: cnvd编号 **取值范围**: 字符长度0-32
	CnvdId *string `json:"cnvd_id,omitempty"`

	// **参数解释**: cnnvd编号 **取值范围**: 字符长度0-32
	CnnvdId *string `json:"cnnvd_id,omitempty"`

	// **参数解释**: bugtraq编号 **取值范围**: 字符长度0-32
	BugtraqId *string `json:"bugtraq_id,omitempty"`

	// **参数解释**: 后缀路径 **取值范围**: 字符长度0-128
	SuffixPath *string `json:"suffix_path,omitempty"`

	// **参数解释**: md5 **取值范围**: 字符长度0-32
	Md5 *string `json:"md5,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 更新时间 **取值范围**: 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 漏洞标签中文名称 **取值范围**: 字符长度0-64
	TagsZh *string `json:"tags_zh,omitempty"`

	// **参数解释**: 漏洞标签英文名称 **取值范围**: 字符长度0-64
	TagsEn *string `json:"tags_en,omitempty"`

	// **参数解释**: 补丁地址 **取值范围**: 字符长度0-512
	PatchUrl *string `json:"patch_url,omitempty"`

	HostsNum *VulnerabilityHostNumberInfo `json:"hosts_num,omitempty"`

	// **参数解释**: cve危险程度 **取值范围**: - Low：低危 - Medium：中危 - High：高危
	CveLevel *string `json:"cve_level,omitempty"`

	// **参数解释**: 漏洞分值 **取值范围**: 最小值0，最大值10
	Cvss *float32 `json:"cvss,omitempty"`

	// **参数解释**: cvss评分版本 **取值范围**: 字符长度0-32
	CvssVersion *string `json:"cvss_version,omitempty"`

	// **参数解释**: 漏洞描述 **取值范围**: 字符长度0-1024
	Description *string `json:"description,omitempty"`

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
}

func (o WebCmsVulDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebCmsVulDetailInfo struct{}"
	}

	return strings.Join([]string{"WebCmsVulDetailInfo", string(data)}, " ")
}
