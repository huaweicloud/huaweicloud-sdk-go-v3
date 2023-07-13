package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerManagedPolicyReferencesInPermissionSetRequest Request Object
type ListCustomerManagedPolicyReferencesInPermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListCustomerManagedPolicyReferencesInPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerManagedPolicyReferencesInPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerManagedPolicyReferencesInPermissionSetRequest", string(data)}, " ")
}
