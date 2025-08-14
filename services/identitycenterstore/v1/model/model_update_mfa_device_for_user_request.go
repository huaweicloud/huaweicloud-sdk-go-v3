package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMfaDeviceForUserRequest Request Object
type UpdateMfaDeviceForUserRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 身份源中IdentityCenter用户的全局唯一标识符（ID）
	UserId string `json:"user_id"`

	// MFA设备ID
	DeviceId string `json:"device_id"`

	Body *UpdateMfaDeviceForUserReqBody `json:"body,omitempty"`
}

func (o UpdateMfaDeviceForUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMfaDeviceForUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateMfaDeviceForUserRequest", string(data)}, " ")
}
