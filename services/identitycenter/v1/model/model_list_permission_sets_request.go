package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionSetsRequest Request Object
type ListPermissionSetsRequest struct {
	InstanceId string `json:"instance_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	// 权限集的全局唯一标识符（ID）
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 权限集urn
	PermissionUrn *string `json:"permission_urn,omitempty"`

	// 权限集名称
	Name *string `json:"name,omitempty"`
}

func (o ListPermissionSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionSetsRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionSetsRequest", string(data)}, " ")
}
