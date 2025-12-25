package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageVulnerabilitiesRequest Request Object
type ListImageVulnerabilitiesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - private_image：私有镜像仓库 - shared_image：共享镜像仓库 - instance_image：企业镜像 - cicd：cicd镜像 - harbor：Harbor仓库镜像 - jfrog：Jfrog仓库镜像  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**: 镜像id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ImageId string `json:"image_id"`

	// **参数解释**： 企业仓库实例ID，SWR企业版可以使用该参数 **约束限制**： 不涉及 **取值范围**： 字符长度0-128位 **默认取值**： 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 组织名称 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	Namespace string `json:"namespace"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128 **默认取值**: 不涉及
	ImageName string `json:"image_name"`

	// **参数解释**: 镜像版本名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	TagName string `json:"tag_name"`

	// **参数解释**: 危险程度 **约束限制**: 不涉及 **取值范围**: - immediate_repair：高危 - delay_repair：中危 - not_needed_repair：低危  **默认取值**: 不涉及
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 漏洞ID（支持模糊查询） **约束限制**: 不涉及 **取值范围**: 字符长度0-128 **默认取值**: 不涉及
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 软件名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞 - app_vul：应用漏洞  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: 处置状态 **约束限制**: 不涉及 **取值范围**: - unhandled：未处理 - handled：已处理  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`
}

func (o ListImageVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListImageVulnerabilitiesRequest", string(data)}, " ")
}
