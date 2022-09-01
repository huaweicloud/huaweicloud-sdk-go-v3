package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeystoneUpdateUserOption struct {

	// IAM用户所属账号ID。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 新IAM用户名，长度1~64之间，只能包含如下字符：大小写字母、空格、数字或特殊字符（-_.）且不能以数字开头
	Name *string `json:"name,omitempty" xml:"name"`

	// IAM用户密码。 - 系统默认密码最小长度为6位字符，在6-32位之间支持用户自定义密码长度。 - 至少包含以下四种字符中的两种： 大写字母、小写字母、数字和特殊字符。 - 不能包含手机号和邮箱。 - 必须满足账户设置中密码策略的要求。 - 新密码不能与当前密码相同。
	Password *string `json:"password,omitempty" xml:"password"`

	// 是否启用IAM用户。true为启用，false为停用，默认为true。
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// IAM用户新描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// IAM用户密码状态。true:需要修改密码,false:正常。
	PwdStatus *bool `json:"pwd_status,omitempty" xml:"pwd_status"`
}

func (o KeystoneUpdateUserOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserOption struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserOption", string(data)}, " ")
}
