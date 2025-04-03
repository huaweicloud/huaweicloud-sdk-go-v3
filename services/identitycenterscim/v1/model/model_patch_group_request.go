package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchGroupRequest Request Object
type PatchGroupRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	// 用户组的全局唯一标识符（ID）
	GroupId string `json:"group_id"`

	Body *PatchGroupReqBody `json:"body,omitempty"`
}

func (o PatchGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchGroupRequest struct{}"
	}

	return strings.Join([]string{"PatchGroupRequest", string(data)}, " ")
}
