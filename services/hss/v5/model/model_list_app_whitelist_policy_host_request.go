package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistPolicyHostRequest Request Object
type ListAppWhitelistPolicyHostRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **策略学习状态**： 策略学习状态 **约束限制**: 不涉及 **取值范围**: - effecting：学习完成，策略生效 - learned：学习完成，待确认 - learning：学习中 - pause：暂停 - abnormal：学习异常  **默认取值**: 不涉及
	LearningStatus *string `json:"learning_status,omitempty"`

	// **策略学习状态**： 策略应用状态 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否  **默认取值**: 不涉及
	ApplyStatus *bool `json:"apply_status,omitempty"`

	// **参数解释**： 资产重要性 **约束限制**： 不涉及 **取值范围**： - important：重要资产。 - common：一般资产。 - test：测试资产。  **默认取值**： 无
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 操作系统类型。 **约束限制**: 不涉及 **取值范围**: - Linux：Linux操作系统。 - Windows：Windows操作系统。  **默认取值**: 无
	OsType *string `json:"os_type,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 服务器弹性IP地址。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`
}

func (o ListAppWhitelistPolicyHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistPolicyHostRequest struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistPolicyHostRequest", string(data)}, " ")
}
