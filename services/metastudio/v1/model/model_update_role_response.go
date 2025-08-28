package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoleResponse Response Object
type UpdateRoleResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRoleResponse", string(data)}, " ")
}
