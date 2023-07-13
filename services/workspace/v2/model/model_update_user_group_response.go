package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserGroupResponse Response Object
type UpdateUserGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserGroupResponse", string(data)}, " ")
}
