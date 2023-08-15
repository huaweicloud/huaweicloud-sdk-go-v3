package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachCustomerManagedPolicyToPermissionSetRequest Request Object
type AttachCustomerManagedPolicyToPermissionSetRequest struct {
	PermissionSetId string `json:"permission_set_id"`

	InstanceId string `json:"instance_id"`

	Body *AttachCustomerManagedPolicyToPermissionSetReqBody `json:"body,omitempty"`
}

func (o AttachCustomerManagedPolicyToPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachCustomerManagedPolicyToPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"AttachCustomerManagedPolicyToPermissionSetRequest", string(data)}, " ")
}
