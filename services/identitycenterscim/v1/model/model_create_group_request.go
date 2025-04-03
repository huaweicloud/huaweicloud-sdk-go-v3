package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupRequest Request Object
type CreateGroupRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	Body *CreateGroupReqBody `json:"body,omitempty"`
}

func (o CreateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateGroupRequest", string(data)}, " ")
}
