package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchUserRequest Request Object
type PatchUserRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	// 用户的全局唯一标识符（ID）
	UserId string `json:"user_id"`

	Body *PatchUserReqBody `json:"body,omitempty"`
}

func (o PatchUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchUserRequest struct{}"
	}

	return strings.Join([]string{"PatchUserRequest", string(data)}, " ")
}
