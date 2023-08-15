package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionsResponse Response Object
type ListPermissionsResponse struct {

	// 共享资源权限的详细信息列表。
	Permissions *[]PermissionSummary `json:"permissions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionsResponse", string(data)}, " ")
}
