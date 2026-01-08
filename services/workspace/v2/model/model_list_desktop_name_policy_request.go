package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopNamePolicyRequest Request Object
type ListDesktopNamePolicyRequest struct {

	// 是否包含用户名的桌面名称策略。 - true 包含 - false 不包含
	IsContainUser *bool `json:"is_contain_user,omitempty"`

	// 策略名称，由数字、字母、中文、下划线组成，必须以字母或下划线开头，长度范围为1~30个字符，支持模糊筛选。
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略id。
	PolicyId *string `json:"policy_id,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-100，默认值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopNamePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopNamePolicyRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopNamePolicyRequest", string(data)}, " ")
}
