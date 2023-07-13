package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedPoliciesInPermissionSetRequest Request Object
type ListManagedPoliciesInPermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListManagedPoliciesInPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedPoliciesInPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"ListManagedPoliciesInPermissionSetRequest", string(data)}, " ")
}
