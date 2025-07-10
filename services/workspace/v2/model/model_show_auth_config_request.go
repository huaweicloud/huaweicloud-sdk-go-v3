package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthConfigRequest Request Object
type ShowAuthConfigRequest struct {

	// 认证类型。LOCAL_PASSWORD：本地密码认证模式，KERBEROS：Windows AD认证模式，LDAP：第三方LDAP模式，CLIENT_TOKEN：金审UKEY客户端Token认证模式，OAUTH2：第三方单点登录模式。
	AuthType *string `json:"auth_type,omitempty"`
}

func (o ShowAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthConfigRequest", string(data)}, " ")
}
