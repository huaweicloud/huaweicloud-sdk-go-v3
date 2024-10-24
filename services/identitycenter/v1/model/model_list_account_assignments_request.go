package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAssignmentsRequest Request Object
type ListAccountAssignmentsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM身份中心实例的全局唯一标识符（ID）。
	InstanceId string `json:"instance_id"`

	// 每个请求返回的最大结果数
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 账号的唯一身份标识
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
