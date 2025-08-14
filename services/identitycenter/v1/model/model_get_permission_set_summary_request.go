package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPermissionSetSummaryRequest Request Object
type GetPermissionSetSummaryRequest struct {

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o GetPermissionSetSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPermissionSetSummaryRequest struct{}"
	}

	return strings.Join([]string{"GetPermissionSetSummaryRequest", string(data)}, " ")
}
