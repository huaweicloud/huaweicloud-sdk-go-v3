package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHaConfigurationRequest Request Object
type UpdateHaConfigurationRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	Body *UpdateHaConfigurationReqBody `json:"body,omitempty"`
}

func (o UpdateHaConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHaConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateHaConfigurationRequest", string(data)}, " ")
}
