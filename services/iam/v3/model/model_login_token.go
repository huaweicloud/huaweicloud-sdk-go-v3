package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type LoginToken struct {
	// 账号ID。

	DomainId string `json:"domain_id"`
	// logintoken的过期时间，默认10min。

	ExpiresAt string `json:"expires_at"`
	// 认证方法。当认证用户为华为云用户时，该字段内容为“token”，当认证用户为自定义代理用户时，该字段内容为“federation_proxy”。

	Method string `json:"method"`
	// 用户ID。

	UserId string `json:"user_id"`
	// 用户名。

	UserName string `json:"user_name"`
	// 会话ID。

	SessionId string `json:"session_id"`
	// 自定义代理用户ID。

	SessionUserId *string `json:"session_user_id,omitempty"`
	// 自定义代理用户名。 > - [通过委托获取临时访问密钥和securitytoken](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=CreateTemporaryAccessKeyByAgency)且请求体中填写session_user.name参数时，会返回该字段。该字段的值即为session_user.name所填写的值。

	SessionName *string `json:"session_name,omitempty"`

	AssumedBy *LoginTokenAssumedBy `json:"assumed_by,omitempty"`
}

func (o LoginToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginToken struct{}"
	}

	return strings.Join([]string{"LoginToken", string(data)}, " ")
}
