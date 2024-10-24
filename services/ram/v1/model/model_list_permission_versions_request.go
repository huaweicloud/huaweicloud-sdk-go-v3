package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionVersionsRequest Request Object
type ListPermissionVersionsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

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
