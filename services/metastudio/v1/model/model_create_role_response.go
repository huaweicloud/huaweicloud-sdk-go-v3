package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoleResponse Response Object
type CreateRoleResponse struct {

	// 角色ID。
	RoleId *string `json:"role_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoleResponse struct{}"
	}

	return strings.Join([]string{"CreateRoleResponse", string(data)}, " ")
}
