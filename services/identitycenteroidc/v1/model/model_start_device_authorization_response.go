package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDeviceAuthorizationResponse Response Object
type StartDeviceAuthorizationResponse struct {

	// 设备在轮询会话令牌时使用的设备码
	DeviceCode *string `json:"device_code,omitempty"`

	// 设备码失效时间（以秒为单位）
	ExpiresIn *int32 `json:"expires_in,omitempty"`

	// 指示轮询会话时，客户端在两次尝试之间必须等待的秒数
	Interval *int32 `json:"interval,omitempty"`

	// 一次性用户验证码。授权正在使用的设备时需要此操作
	UserCode *string `json:"user_code,omitempty"`

	// 使用一次性用户验证码授权设备的验证页面的URI
	VerificationUri *string `json:"verification_uri,omitempty"`

	// 客户端可用于自动启动浏览器的备用URL。此过程跳过用户访问验证页面并输入代码的手动步骤
	VerificationUriComplete *string `json:"verification_uri_complete,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o StartDeviceAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDeviceAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"StartDeviceAuthorizationResponse", string(data)}, " ")
}
