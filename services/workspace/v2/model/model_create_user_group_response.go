package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserGroupResponse Response Object
type CreateUserGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateUserGroupResponse", string(data)}, " ")
}
