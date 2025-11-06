package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupPermissionResourcesResponse Response Object
type ListGroupPermissionResourcesResponse struct {

	// **参数解释：** 是否使用项目权限。
	UseProjectPermission *bool `json:"use_project_permission,omitempty"`

	// **参数解释：** 资源列表。
	Resources      *[]PermissionResourcesDto `json:"resources,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListGroupPermissionResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupPermissionResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListGroupPermissionResourcesResponse", string(data)}, " ")
}
