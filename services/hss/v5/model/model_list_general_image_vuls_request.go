package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeneralImageVulsRequest Request Object
type ListGeneralImageVulsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞 - app_vul：应用漏洞  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 漏洞的处理状态 **约束限制**: 不涉及 **取值范围**:： - unhandled：未处理 - handled：已处理  **默认取值**: 不涉及
	HandleStatus string `json:"handle_status"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - local：本地镜像 - registry：仓库镜像 - cicd：CI/CD镜像 - cluster：集群镜像  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**: 漏洞的风险程度 **约束限制**: 不涉及 **取值范围**: - Critical：严重 - High：高危 - Medium：中危 - Low：低危  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 漏洞id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞编号 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: 漏洞标签列表 **取值范围**: 最小值0，最大值10
	LabelList *[]string `json:"label_list,omitempty"`

	// **参数解释**: 漏洞当前状态 **约束限制**: 不涉及 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：已忽略 **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ListGeneralImageVulsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeneralImageVulsRequest struct{}"
	}

	return strings.Join([]string{"ListGeneralImageVulsRequest", string(data)}, " ")
}
