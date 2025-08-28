package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulAffectImagesRequest Request Object
type ListVulAffectImagesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 查询已处理/未处理的镜像 **约束限制**: 不涉及 **取值范围**: - true：查询已处理的镜像 - false：查询未处理的镜像  **默认取值**: 不涉及
	IsHandled bool `json:"is_handled"`

	// **参数解释**: 漏洞id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	VulId string `json:"vul_id"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - app_vul：应用漏洞 - linux_vul：系统漏洞  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - local：本地镜像 - registry：仓库镜像 - cicd：CI/CD镜像 - cluster：集群镜像  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**: 危险程度 **约束限制**: 不涉及 **取值范围**: - Low：低危 - Medium：中危 - High：高危 - Critical：严重  **默认取值**: 不涉及
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 镜像id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像仓名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 镜像仓类型 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释**: 漏洞对应镜像的处理状态 **约束限制**: 不涉及 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：已忽略  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ListVulAffectImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulAffectImagesRequest struct{}"
	}

	return strings.Join([]string{"ListVulAffectImagesRequest", string(data)}, " ")
}
