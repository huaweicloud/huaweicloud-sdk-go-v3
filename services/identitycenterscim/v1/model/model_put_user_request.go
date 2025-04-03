package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutUserRequest Request Object
type PutUserRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	// 用户的全局唯一标识符（ID）
	UserId string `json:"user_id"`

	Body *UpdateUserReqBody `json:"body,omitempty"`
}

func (o PutUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutUserRequest struct{}"
	}

	return strings.Join([]string{"PutUserRequest", string(data)}, " ")
}
