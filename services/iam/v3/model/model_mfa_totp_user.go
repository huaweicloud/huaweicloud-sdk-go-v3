package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type MfaTotpUser struct {
	// 已开启虚拟MFA方式的登录保护的IAM用户ID。

	Id string `json:"id"`
	// 虚拟MFA验证码，在MFA应用程序中获取动态验证码，获取方法请参见：[如何获取虚拟MFA验证码](https://support.huaweicloud.com/iam_faq/iam_01_0001.html)。

	Passcode string `json:"passcode"`
}

func (o MfaTotpUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaTotpUser struct{}"
	}

	return strings.Join([]string{"MfaTotpUser", string(data)}, " ")
}
