package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachCustomerManagedPolicyReferenceFromPermissionSetRequest Request Object
type DetachCustomerManagedPolicyReferenceFromPermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Body *DetachCustomerManagedPolicyReferenceFromPermissionSetReqBody `json:"body,omitempty"`
}

func (o DetachCustomerManagedPolicyReferenceFromPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachCustomerManagedPolicyReferenceFromPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"DetachCustomerManagedPolicyReferenceFromPermissionSetRequest", string(data)}, " ")
}
