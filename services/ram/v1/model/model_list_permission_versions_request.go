package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionVersionsRequest Request Object
type ListPermissionVersionsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 共享资源权限的ID。
	PermissionId string `json:"permission_id"`
}

func (o ListPermissionVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionVersionsRequest", string(data)}, " ")
}
