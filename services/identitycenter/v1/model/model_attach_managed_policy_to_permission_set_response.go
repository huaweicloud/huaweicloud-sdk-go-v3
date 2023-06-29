package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachManagedPolicyToPermissionSetResponse Response Object
type AttachManagedPolicyToPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachManagedPolicyToPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachManagedPolicyToPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"AttachManagedPolicyToPermissionSetResponse", string(data)}, " ")
}
