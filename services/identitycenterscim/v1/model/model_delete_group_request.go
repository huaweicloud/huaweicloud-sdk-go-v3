package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupRequest Request Object
type DeleteGroupRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	// 用户组的全局唯一标识符（ID）
	GroupId string `json:"group_id"`
}

func (o DeleteGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupRequest", string(data)}, " ")
}
