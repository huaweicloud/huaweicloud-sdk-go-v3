package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralImageVulsResponseInfo 所有镜像漏洞详情
type GeneralImageVulsResponseInfo struct {

	// **参数解释**： 漏洞名称 **取值范围**： 字符长度0-128位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 漏洞ID **取值范围**： 字符长度0-128位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**： 漏洞类型 **约束限制**: 不涉及 **取值范围**： - linux_vul：linux漏洞 - app_vul：应用漏洞 **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 镜像类型 **约束限制**: 不涉及 **取值范围**： - local：本地镜像 - registry：仓库镜像 - cicd：CI/CD镜像 - cluster：集群镜像 **默认取值**: 不涉及
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**： 漏洞标签列表 **取值范围**： 最小值0，最大值10
	LabelList *[]string `json:"label_list,omitempty"`

	// **参数解释**： 漏洞的风险程度 **约束限制**: 不涉及 **取值范围**： - Critical：严重 - High：高危 - Medium：中危 - Low：低危 **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**： 受影响镜像总数 **取值范围**： 字符长度0-2147483547位
	ImageNum *int32 `json:"image_num,omitempty"`

	// **参数解释**： CVE列表 **取值范围**： 最小值0，最大值10000
	CveList *[]GeneralImageVulsResponseInfoCveList `json:"cve_list,omitempty"`

	// **参数解释**： 镜像最大CVSS分值 **取值范围**： 字符长度0-10位
	MaxCvssScore *float32 `json:"max_cvss_score,omitempty"`

	// **参数解释**： 最近扫描时间，时间单位：毫秒（ms） **取值范围**： 字符长度0-4070880000000位
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 漏洞描述 **取值范围**： 字符长度0-128位
	Description *string `json:"description,omitempty"`

	// **参数解释**： 漏洞修复参考链接 **取值范围**： 字符长度0-128位
	Url *string `json:"url,omitempty"`

	// **参数解释**： 修复建议 **取值范围**： 字符长度0-128位
	SolutionDetail *string `json:"solution_detail,omitempty"`

	// **参数解释**： 受影响集群总数 **取值范围**： 字符长度0-2147483547位
	ClusterNum *int32 `json:"cluster_num,omitempty"`
}

func (o GeneralImageVulsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralImageVulsResponseInfo struct{}"
	}

	return strings.Join([]string{"GeneralImageVulsResponseInfo", string(data)}, " ")
}
