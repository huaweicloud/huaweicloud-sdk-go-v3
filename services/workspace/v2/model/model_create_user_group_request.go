package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserGroupRequest Request Object
type CreateUserGroupRequest struct {
	Body *CreateUserGroupReq `json:"body,omitempty"`
}

func (o CreateUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateUserGroupRequest", string(data)}, " ")
}
