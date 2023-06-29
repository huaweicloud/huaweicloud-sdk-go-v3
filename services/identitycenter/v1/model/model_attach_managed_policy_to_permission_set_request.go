package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachManagedPolicyToPermissionSetRequest Request Object
type AttachManagedPolicyToPermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Body *AttachManagedPolicyToPermissionSetReqBody `json:"body,omitempty"`
}

func (o AttachManagedPolicyToPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachManagedPolicyToPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"AttachManagedPolicyToPermissionSetRequest", string(data)}, " ")
}
