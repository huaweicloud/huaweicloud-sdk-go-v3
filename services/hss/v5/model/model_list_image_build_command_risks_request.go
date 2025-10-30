package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageBuildCommandRisksRequest Request Object
type ListImageBuildCommandRisksRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**:   - cicd：cicd镜像   - registry：仓库镜像 字符长度1-32  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**： 镜像id **约束限制**： 不涉及 **取值范围**： 字符长度0-128位 **默认取值**： 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 风险名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-512位 **默认取值**： 不涉及
	RiskName *string `json:"risk_name,omitempty"`

	// **参数解释**: 风险程度 **约束限制**: 不涉及 **取值范围**:   - critical ：严重   - high ：高危   - medium ：中危   - low ：低危 字符长度1-64  **默认取值**: 不涉及
	RiskLevel *string `json:"risk_level,omitempty"`
}

func (o ListImageBuildCommandRisksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageBuildCommandRisksRequest struct{}"
	}

	return strings.Join([]string{"ListImageBuildCommandRisksRequest", string(data)}, " ")
}
