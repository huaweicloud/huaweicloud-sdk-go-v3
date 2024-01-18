package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsForAccountRequest Request Object
type ListControlsForAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListControlsForAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsForAccountRequest struct{}"
	}

	return strings.Join([]string{"ListControlsForAccountRequest", string(data)}, " ")
}
