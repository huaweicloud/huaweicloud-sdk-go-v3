package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAssignmentsRequest Request Object
type ListAccountAssignmentsRequest struct {
	InstanceId string `json:"instance_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	// 指定账户的唯一身份标识.
	AccountId string `json:"account_id"`

	// 指定权限集的唯一身份标识.
	PermissionSetId *string `json:"permission_set_id,omitempty"`
}

func (o ListAccountAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentsRequest", string(data)}, " ")
}
