package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionVersionsResponse Response Object
type ListPermissionVersionsResponse struct {

	// 共享资源权限的详细信息列表。
	Permissions *[]PermissionSummary `json:"permissions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPermissionVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionVersionsResponse", string(data)}, " ")
}
