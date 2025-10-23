package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchValidateUserGroupPermissionsResponse Response Object
type BatchValidateUserGroupPermissionsResponse struct {

	// **参数解释：** 代码组权限列表。
	Body           *[]UserGroupPermissionDto `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchValidateUserGroupPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateUserGroupPermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchValidateUserGroupPermissionsResponse", string(data)}, " ")
}
