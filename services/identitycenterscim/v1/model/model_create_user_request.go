package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserRequest Request Object
type CreateUserRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	Body *CreateUserReqBody `json:"body,omitempty"`
}

func (o CreateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserRequest struct{}"
	}

	return strings.Join([]string{"CreateUserRequest", string(data)}, " ")
}
