package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBearerTokenResponse Response Object
type CreateBearerTokenResponse struct {

	// 创建时间
	CreationTime float32 `json:"creation_time,omitempty"`

	// 过期时间
	ExpirationTime float32 `json:"expiration_time,omitempty"`

	// 访问令牌
	Token *string `json:"token,omitempty"`

	// 访问令牌的全局唯一标识符（ID）
	TokenId        *string `json:"token_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBearerTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBearerTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateBearerTokenResponse", string(data)}, " ")
}
