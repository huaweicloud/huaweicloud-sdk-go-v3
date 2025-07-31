package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PwdPolicyInfoResponseInfo 服务器的口令复杂度策略。建议设置最小口令长度不小于8，同时包含大写字母、小写字母、数字和特殊字符。
type PwdPolicyInfoResponseInfo struct {

	// **参数解释**: 主机id **取值范围**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP（私有IP），为兼容用户使用，不删除此字段 **取值范围**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 服务器私有IP **取值范围**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器公网IP **取值范围**: 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 口令最小长度的设置是否符合要求 **取值范围**: - true：符合要求 - false：不符合要求
	MinLength *bool `json:"min_length,omitempty"`

	// **参数解释**: 大写字母的设置是否符合要求 **取值范围**: - true：符合要求 - false：不符合要求
	UppercaseLetter *bool `json:"uppercase_letter,omitempty"`

	// **参数解释**: 小写字母的设置是否符合要求 **取值范围**: - true：符合要求 - false：不符合要求
	LowercaseLetter *bool `json:"lowercase_letter,omitempty"`

	// **参数解释**: 数字的设置是否符合要求，符合为true，不符合为false **取值范围**: - true：符合要求 - false：不符合要求
	Number *bool `json:"number,omitempty"`

	// **参数解释**: 特殊字符的设置是否符合要求，符合为true，不符合为false **取值范围**: - true：符合要求 - false：不符合要求
	SpecialCharacter *bool `json:"special_character,omitempty"`

	// **参数解释**: 复杂口令策略中定义的口令最小长度 **取值范围**: 8 - 26
	MinLengthNum *int32 `json:"min_length_num,omitempty"`

	// **参数解释**: 复杂口令策略中定义的最少包含的大写字母数 **取值范围**: 0 - 10
	MinUppercaseLetter *int32 `json:"min_uppercase_letter,omitempty"`

	// **参数解释**: 复杂口令策略中定义的最少包含的小写字母数 **取值范围**: 0 - 10
	MinLowercaseLetter *int32 `json:"min_lowercase_letter,omitempty"`

	// **参数解释**: 复杂口令策略中定义的最少包含的数字数 **取值范围**: 0 - 10
	MinNumber *int32 `json:"min_number,omitempty"`

	// **参数解释**: 复杂口令策略中定义的最少包含的特殊字母数 **取值范围**: 0 - 10
	MinSpecialCharacter *int32 `json:"min_special_character,omitempty"`

	// **参数解释**: 最近扫描时间 **取值范围**: 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 修改建议 **取值范围**: 不涉及
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o PwdPolicyInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdPolicyInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"PwdPolicyInfoResponseInfo", string(data)}, " ")
}
