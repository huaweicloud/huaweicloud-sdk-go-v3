package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedAccountsRequest Request Object
type ListManagedAccountsRequest struct {

	// 启用的控制策略信息。
	ControlId *string `json:"control_id,omitempty"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListManagedAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedAccountsRequest struct{}"
	}

	return strings.Join([]string{"ListManagedAccountsRequest", string(data)}, " ")
}
