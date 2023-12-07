package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityPermissionSetMembersResponse Response Object
type BatchDeleteSecurityPermissionSetMembersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityPermissionSetMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityPermissionSetMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityPermissionSetMembersResponse", string(data)}, " ")
}
