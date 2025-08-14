package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterMfaDeviceForUserResponse Response Object
type RegisterMfaDeviceForUserResponse struct {

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// 身份源中用户唯一标识符（ID）
	UserId *string `json:"user_id,omitempty"`

	// 注册MFA需要的一次性随机字符
	WorkFlow *string `json:"work_flow,omitempty"`

	// MFA注册重定向地址
	RedirectUrl    *string `json:"redirect_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterMfaDeviceForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterMfaDeviceForUserResponse struct{}"
	}

	return strings.Join([]string{"RegisterMfaDeviceForUserResponse", string(data)}, " ")
}
