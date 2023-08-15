package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachManagedPolicyFromPermissionSetRequest Request Object
type DetachManagedPolicyFromPermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Body *DetachManagedPolicyFromPermissionSetReqBody `json:"body,omitempty"`
}

func (o DetachManagedPolicyFromPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachManagedPolicyFromPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"DetachManagedPolicyFromPermissionSetRequest", string(data)}, " ")
}
