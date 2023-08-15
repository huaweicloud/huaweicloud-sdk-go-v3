package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachCustomerManagedPolicyReferenceFromPermissionSetResponse Response Object
type DetachCustomerManagedPolicyReferenceFromPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachCustomerManagedPolicyReferenceFromPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachCustomerManagedPolicyReferenceFromPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DetachCustomerManagedPolicyReferenceFromPermissionSetResponse", string(data)}, " ")
}
