package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachCustomerManagedPolicyReferenceFromPermissionSetReqBody DetachCustomerManagedPolicyReferenceFromPermissionSet请求体
type DetachCustomerManagedPolicyReferenceFromPermissionSetReqBody struct {
	CustomerManagedPolicyReference *CustomerManagedPolicyReferenceDto `json:"customer_managed_policy_reference"`
}

func (o DetachCustomerManagedPolicyReferenceFromPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachCustomerManagedPolicyReferenceFromPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"DetachCustomerManagedPolicyReferenceFromPermissionSetReqBody", string(data)}, " ")
}
