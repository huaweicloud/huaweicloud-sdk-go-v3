package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachCustomerManagedPolicyToPermissionSetResponse Response Object
type AttachCustomerManagedPolicyToPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachCustomerManagedPolicyToPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachCustomerManagedPolicyToPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"AttachCustomerManagedPolicyToPermissionSetResponse", string(data)}, " ")
}
