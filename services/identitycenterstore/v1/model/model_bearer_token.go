package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BearerToken struct {

	// 创建时间
	CreationTime float32 `json:"creation_time"`

	// 过期时间
	ExpirationTime float32 `json:"expiration_time"`

	// 访问令牌全局唯一标识（ID）
	TokenId string `json:"token_id"`
}

func (o BearerToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BearerToken struct{}"
	}

	return strings.Join([]string{"BearerToken", string(data)}, " ")
}
