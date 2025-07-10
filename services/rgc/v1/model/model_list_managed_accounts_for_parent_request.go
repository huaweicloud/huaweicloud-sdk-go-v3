package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedAccountsForParentRequest Request Object
type ListManagedAccountsForParentRequest struct {

	// 注册OU ID。
	ManagedOrganizationalUnitId string `json:"managed_organizational_unit_id"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListManagedAccountsForParentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedAccountsForParentRequest struct{}"
	}

	return strings.Join([]string{"ListManagedAccountsForParentRequest", string(data)}, " ")
}
