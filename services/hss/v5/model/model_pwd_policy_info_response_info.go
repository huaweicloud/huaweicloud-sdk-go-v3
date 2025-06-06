package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PwdPolicyInfoResponseInfo 服务器的口令复杂度策略。建议设置最小口令长度不小于8，同时包含大写字母、小写字母、数字和特殊字符。
type PwdPolicyInfoResponseInfo struct {

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器IP（私有IP），为兼容用户使用，不删除此字段
	HostIp *string `json:"host_ip,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 口令最小长度的设置是否符合要求，符合为true，不符合为false
	MinLength *bool `json:"min_length,omitempty"`

	// 大写字母的设置是否符合要求，符合为true，不符合为false
	UppercaseLetter *bool `json:"uppercase_letter,omitempty"`

	// 小写字母的设置是否符合要求，符合为true，不符合为false
	LowercaseLetter *bool `json:"lowercase_letter,omitempty"`

	// 数字的设置是否符合要求，符合为true，不符合为false
	Number *bool `json:"number,omitempty"`

	// 特殊字符的设置是否符合要求，符合为true，不符合为false
	SpecialCharacter *bool `json:"special_character,omitempty"`

	// 复杂口令策略中定义的口令最小长度
	MinLengthNum *int32 `json:"min_length_num,omitempty"`

	// 复杂口令策略中定义的最少包含的大写字母数
	MinUppercaseLetter *int32 `json:"min_uppercase_letter,omitempty"`

	// 复杂口令策略中定义的最少包含的小写字母数
	MinLowercaseLetter *int32 `json:"min_lowercase_letter,omitempty"`

	// 复杂口令策略中定义的最少包含的数字数
	MinNumber *int32 `json:"min_number,omitempty"`

	// 复杂口令策略中定义的最少包含的特殊字母数
	MinSpecialCharacter *int32 `json:"min_special_character,omitempty"`

	// 最近扫描时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 修改建议
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o PwdPolicyInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdPolicyInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"PwdPolicyInfoResponseInfo", string(data)}, " ")
}
