package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckRulesInfoRequest Request Object
type ListCheckRulesInfoRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 类型 **约束限制**: 不涉及 **取值范围**: - image : 镜像  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - cicd：cicd镜像 - registry：仓库镜像  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**: 组织名称 **约束限制**: 不涉及 **取值范围**: 字符串大小0-65535 **默认取值**: 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符串大小0-65535 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本名称 **约束限制**: 不涉及 **取值范围**: 字符串大小0-65535 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 企业仓库实例ID，swr共享版无需使用该参数 **约束限制**: 不涉及 **取值范围**: 字符串大小0-128 **默认取值**: 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 镜像id **约束限制**： 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 检测结果 **约束限制**: 不涉及 **取值范围**:   - pass：通过   - failed：未通过 **默认取值**: 不涉及
	ScanResult *string `json:"scan_result,omitempty"`
}

func (o ListCheckRulesInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckRulesInfoRequest struct{}"
	}

	return strings.Join([]string{"ListCheckRulesInfoRequest", string(data)}, " ")
}
