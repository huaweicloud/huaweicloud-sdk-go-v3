package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntrusionSubscriptionRes 入侵检测订阅的项目
type IntrusionSubscriptionRes struct {

	// **参数解释**: 风险账号 **取值范围**: - true：开启 - false：关闭
	RiskyAccountStatus *bool `json:"risky_account_status,omitempty"`

	// **参数解释**: 异地登录 **取值范围**: - true：开启 - false：关闭
	AbnormalLoginStatus *bool `json:"abnormal_login_status,omitempty"`

	// **参数解释**: 恶意程序 **取值范围**: - true：开启 - false：关闭
	MaliciousProgramStatus *bool `json:"malicious_program_status,omitempty"`

	// **参数解释**: 网站后门 **取值范围**: - true：开启 - false：关闭
	WebshellStatus *bool `json:"webshell_status,omitempty"`

	// **参数解释**: 账号破解 **取值范围**: - true：开启 - false：关闭
	AccountCrackStatus *bool `json:"account_crack_status,omitempty"`

	// **参数解释**: 关键文件变更 **取值范围**: - true：开启 - false：关闭
	CriticalFileChangeStatus *bool `json:"critical_file_change_status,omitempty"`
}

func (o IntrusionSubscriptionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntrusionSubscriptionRes struct{}"
	}

	return strings.Join([]string{"IntrusionSubscriptionRes", string(data)}, " ")
}
