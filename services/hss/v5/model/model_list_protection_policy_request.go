package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectionPolicyRequest Request Object
type ListProtectionPolicyRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 取值0-2000000 **默认取值**: 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 防护策略名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-255 **默认取值**: 不涉及
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 防护策略id **约束限制**: 不涉及 **取值范围**: 字符长度0-128 **默认取值**: 不涉及
	ProtectPolicyId *string `json:"protect_policy_id,omitempty"`

	// **参数解释**: 策略支持的操作系统 **约束限制**: 不涉及 **取值范围**: 包含如下：   - Windows : Windows系统   - Linux : Linux系统 **默认取值**: 不涉及
	OperatingSystem *string `json:"operating_system,omitempty"`
}

func (o ListProtectionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListProtectionPolicyRequest", string(data)}, " ")
}
