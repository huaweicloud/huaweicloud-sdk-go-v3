package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TokenInfoDto struct {

	// 用于访问分配给用户的IAM身份中心资源的不透明令牌
	AccessToken *string `json:"access_token,omitempty"`

	// 访问令牌的过期时间（以秒为单位）
	ExpiresIn *int32 `json:"expires_in,omitempty"`

	// 用于表明用户身份的不透明令牌
	IdToken *string `json:"id_token,omitempty"`

	// 刷新令牌，此令牌可在访问令牌过期后获取新的访问令牌
	RefreshToken *string `json:"refresh_token,omitempty"`

	// 用于通知客户端返回的令牌是访问令牌，目前为BearerToken
	TokenType *string `json:"token_type,omitempty"`
}

func (o TokenInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenInfoDto struct{}"
	}

	return strings.Join([]string{"TokenInfoDto", string(data)}, " ")
}
