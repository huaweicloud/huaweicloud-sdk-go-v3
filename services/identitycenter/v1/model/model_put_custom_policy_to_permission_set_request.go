package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutCustomPolicyToPermissionSetRequest Request Object
type PutCustomPolicyToPermissionSetRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 权限集的全局唯一标识符（ID）
	PermissionSetId string `json:"permission_set_id"`

	Body *PutCustomPolicyToPermissionSetReqBody `json:"body,omitempty"`
}

func (o PutCustomPolicyToPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCustomPolicyToPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"PutCustomPolicyToPermissionSetRequest", string(data)}, " ")
}
