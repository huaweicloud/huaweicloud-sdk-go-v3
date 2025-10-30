package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageBaselineStatisticRequest Request Object
type ShowImageBaselineStatisticRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - private_image：私有镜像仓库。 - shared_image：共享镜像仓库。 - local_image : 本地镜像。 - instance_image：企业镜像。 - registry : 仓库镜像。 - local : 本地镜像，用于查询全局数据。  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**: 组织名称（没有镜像相关信息时，表示查询所有镜像） **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 企业仓库实例ID，swr共享版无需使用该参数 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 镜像id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ImageId *string `json:"image_id,omitempty"`
}

func (o ShowImageBaselineStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageBaselineStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowImageBaselineStatisticRequest", string(data)}, " ")
}
