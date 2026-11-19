package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntrusionInfoRes 入侵信息
type IntrusionInfoRes struct {

	// **参数解释**： 风险账号列表 **取值范围**： 0-200个RiskyAccountInfo对象
	RiskyAccountInfoList *[]interface{} `json:"risky_account_info_list,omitempty"`

	// **参数解释**： 异地登录列表 **取值范围**： 0-200个AbnormalLoginInfo对象
	AbnormalLoginInfoList *[]interface{} `json:"abnormal_login_info_list,omitempty"`

	// **参数解释**： 恶意程序列表 **取值范围**： 0-200个MaliciousProgramInfo对象
	MaliciousProgramInfoList *[]interface{} `json:"malicious_program_info_list,omitempty"`

	// **参数解释**： 网站后门列表 **取值范围**： 0-200个WebshellInfo对象
	WebshellInfoList *[]interface{} `json:"webshell_info_list,omitempty"`

	// **参数解释**： 账号破解列表 **取值范围**： 0-200个AccountCrackInfo对象
	AccountCrackInfoList *[]interface{} `json:"account_crack_info_list,omitempty"`

	// **参数解释**： 关键文件变更列表 **取值范围**： 0-200个CriticalFileChangeInfo对象
	CriticalFileChangeInfoList *[]interface{} `json:"critical_file_change_info_list,omitempty"`
}

func (o IntrusionInfoRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntrusionInfoRes struct{}"
	}

	return strings.Join([]string{"IntrusionInfoRes", string(data)}, " ")
}
