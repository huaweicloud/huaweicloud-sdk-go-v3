package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAssignmentsRequest Request Object
type ListAccountAssignmentsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// The identifier of the account from which to list the assignments.
	AccountId string `json:"account_id"`

	// The identifier of the permission set from which to list assignments.
	PermissionSetId *string `json:"permission_set_id,omitempty"`
}

func (o ListAccountAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentsRequest", string(data)}, " ")
}
