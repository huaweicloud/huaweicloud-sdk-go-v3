package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作保护策略
type ProtectPolicyOption struct {

	// 是否开启操作保护，开启为\"true\"，未开启为\"false\"。
	OperationProtection bool `json:"operation_protection" xml:"operation_protection"`

	AllowUser *AllowUserBody `json:"allow_user,omitempty" xml:"allow_user"`

	// 操作保护验证指定手机号码。示例：0086-123456789。
	Mobile *string `json:"mobile,omitempty" xml:"mobile"`

	// 是否指定人员验证。on为指定人员验证，必须填写scene参数。off为操作员验证。
	AdminCheck *string `json:"admin_check,omitempty" xml:"admin_check"`

	// 操作保护验证指定邮件地址。示例：example@email.com。
	Email *string `json:"email,omitempty" xml:"email"`

	// 操作保护指定人员验证方式，admin_check为on时，必须填写。包括mobile、email。
	Scene *string `json:"scene,omitempty" xml:"scene"`
}

func (o ProtectPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectPolicyOption struct{}"
	}

	return strings.Join([]string{"ProtectPolicyOption", string(data)}, " ")
}
