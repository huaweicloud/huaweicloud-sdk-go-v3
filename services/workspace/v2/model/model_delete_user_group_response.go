package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserGroupResponse Response Object
type DeleteUserGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserGroupResponse", string(data)}, " ")
}
