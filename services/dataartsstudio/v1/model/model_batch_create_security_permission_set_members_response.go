package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSecurityPermissionSetMembersResponse Response Object
type BatchCreateSecurityPermissionSetMembersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSecurityPermissionSetMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityPermissionSetMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityPermissionSetMembersResponse", string(data)}, " ")
}
