package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchValidateUserGroupPermissionsRequest Request Object
type BatchValidateUserGroupPermissionsRequest struct {

	// **参数解释：** 代码组权限列表。
	Body *[]GroupPermissionsDto `json:"body,omitempty"`
}

func (o BatchValidateUserGroupPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateUserGroupPermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchValidateUserGroupPermissionsRequest", string(data)}, " ")
}
