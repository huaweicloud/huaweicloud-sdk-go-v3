package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachManagedPolicyFromPermissionSetRequest Request Object
type DetachManagedPolicyFromPermissionSetRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM身份中心实例的全局唯一标识符（ID）。
	InstanceId string `json:"instance_id"`

	// 权限集的全局唯一标识符（ID）
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
