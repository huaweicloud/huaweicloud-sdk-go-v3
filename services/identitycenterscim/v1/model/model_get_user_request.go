package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetUserRequest Request Object
type GetUserRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	// 用户的全局唯一标识符（ID）
	UserId string `json:"user_id"`
}

func (o GetUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUserRequest struct{}"
	}

	return strings.Join([]string{"GetUserRequest", string(data)}, " ")
}
