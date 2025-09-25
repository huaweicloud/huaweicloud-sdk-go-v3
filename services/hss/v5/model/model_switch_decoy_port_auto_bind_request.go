package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchDecoyPortAutoBindRequest Request Object
type SwitchDecoyPortAutoBindRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 是否启用 **约束限制**: 不涉及 **取值范围**: -true：是。 -false：否。 **默认取值**: false
	Enable bool `json:"enable"`

	// **参数解释**: linux策略id **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	LinuxPolicyId *string `json:"linux_policy_id,omitempty"`

	// **参数解释**: windows策略id **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	WindowsPolicyId *string `json:"windows_policy_id,omitempty"`
}

func (o SwitchDecoyPortAutoBindRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchDecoyPortAutoBindRequest struct{}"
	}

	return strings.Join([]string{"SwitchDecoyPortAutoBindRequest", string(data)}, " ")
}
