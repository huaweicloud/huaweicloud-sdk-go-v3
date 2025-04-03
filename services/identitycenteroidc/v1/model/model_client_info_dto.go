package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClientInfoDto struct {

	// 客户端可以请求授权的端点
	AuthorizationEndpoint string `json:"authorization_endpoint"`

	// 客户端应用唯一标识
	ClientId string `json:"client_id"`

	// 客户端标识符和客户端密钥的注册时间
	ClientIdIssuedAt int64 `json:"client_id_issued_at"`

	// 为客户端生成的秘密字符串。客户端将使用此字符串在后续调用中获得服务的身份验证
	ClientSecret string `json:"client_secret"`

	// 客户端标识符和客户端密钥失效的时间
	ClientSecretExpiresAt int64 `json:"client_secret_expires_at"`

	// 客户端可以在其中获取访问令牌的端点
	TokenEndpoint string `json:"token_endpoint"`

	// 服务器为客户端注册的作用域列表。后续授权访问令牌时，权限都应该限制在此作用域列表的子集范围内
	Scopes *[]string `json:"scopes,omitempty"`
}

func (o ClientInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientInfoDto struct{}"
	}

	return strings.Join([]string{"ClientInfoDto", string(data)}, " ")
}
