package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateUserRoleResponse struct {

	// 用户id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRoleResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserRoleResponse", string(data)}, " ")
}
