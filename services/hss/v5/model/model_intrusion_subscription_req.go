package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntrusionSubscriptionReq 入侵订阅信息
type IntrusionSubscriptionReq struct {

	// **参数解释**： 风险账号是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	RiskyAccountStatus *bool `json:"risky_account_status,omitempty"`

	// **参数解释**： 异地登录是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	AbnormalLoginStatus *bool `json:"abnormal_login_status,omitempty"`

	// **参数解释**： 恶意程序是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	MaliciousProgramStatus *bool `json:"malicious_program_status,omitempty"`

	// **参数解释**： 网站后门是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	WebshellStatus *bool `json:"webshell_status,omitempty"`

	// **参数解释**： 账号破解是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	AccountCrackStatus *bool `json:"account_crack_status,omitempty"`

	// **参数解释**： 关键文件变更是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	CriticalFileChangeStatus *bool `json:"critical_file_change_status,omitempty"`
}

func (o IntrusionSubscriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntrusionSubscriptionReq struct{}"
	}

	return strings.Join([]string{"IntrusionSubscriptionReq", string(data)}, " ")
}
