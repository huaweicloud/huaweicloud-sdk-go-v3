package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRoleRequest Request Object
type UpdateUserRoleRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *UpdateUserRoleReq `json:"body,omitempty"`
}

func (o UpdateUserRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRoleRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRoleRequest", string(data)}, " ")
}
