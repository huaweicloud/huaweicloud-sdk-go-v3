package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionSetsRequest Request Object
type ListPermissionSetsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM身份中心实例的全局唯一标识符（ID）。
	InstanceId string `json:"instance_id"`

	// 每个请求返回的最大结果数
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
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
