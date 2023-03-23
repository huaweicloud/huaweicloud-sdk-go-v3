package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourceSharePermissionsResponse struct {

	// 资源共享实例关联的共享资源权限信息列表。
	AssociatedPermissions *[]AssociatedPermission `json:"associated_permissions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResourceSharePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSharePermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceSharePermissionsResponse", string(data)}, " ")
}
