package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsRequest Request Object
type ListGroupsRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`
}

func (o ListGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupsRequest", string(data)}, " ")
}
