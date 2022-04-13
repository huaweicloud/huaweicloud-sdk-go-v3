package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateLoginProject struct {
	// IAM用户是否开启登录保护，开启为\"true\"，未开启为\"false\"。

	Enabled bool `json:"enabled"`
	// IAM用户登录验证方式。手机验证为“sms”,邮箱验证为“email”,MFA验证为“vmfa”。

	VerificationMethod string `json:"verification_method"`
}

func (o UpdateLoginProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginProject struct{}"
	}

	return strings.Join([]string{"UpdateLoginProject", string(data)}, " ")
}
