package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeystoneUpdatePasswordOption struct {
	// IAM用户的新密码。 - 系统默认密码最小长度为6位字符，在6-32位之间支持用户自定义密码长度。 - 至少包含以下四种字符中的两种： 大写字母、小写字母、数字和特殊字符。 - 不能包含手机号和邮箱。 - 必须满足用户所属账号的[密码策略](https://support.huaweicloud.com/usermanual-iam/iam_01_0607.html)要求。 - 新密码不能与当前密码相同。

	Password string `json:"password"`
	// IAM用户的原密码。

	OriginalPassword string `json:"original_password"`
}

func (o KeystoneUpdatePasswordOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdatePasswordOption struct{}"
	}

	return strings.Join([]string{"KeystoneUpdatePasswordOption", string(data)}, " ")
}
