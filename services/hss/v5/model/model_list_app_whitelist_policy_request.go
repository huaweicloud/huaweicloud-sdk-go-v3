package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistPolicyRequest Request Object
type ListAppWhitelistPolicyRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 进程白名单策略类型 **约束限制**: 不涉及 **取值范围**: - allow：允许指定/授权进程运行 - block：阻止潜在恶意软件运行  **默认取值**: 不涉及
	PolicyType *string `json:"policy_type,omitempty"`

	// **参数解释**： 策略学习状态 **约束限制**: 不涉及 **取值范围**: - effecting：学习完成，策略生效 - learned：学习完成，待确认 - learning：学习中 - pause：暂停 - abnormal：学习异常  **默认取值**: 不涉及
	LearningStatus *string `json:"learning_status,omitempty"`

	// **参数解释**： 是否开启阻断 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否  **默认取值**: 不涉及
	Intercept *bool `json:"intercept,omitempty"`
}

func (o ListAppWhitelistPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistPolicyRequest", string(data)}, " ")
}
