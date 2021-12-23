package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备接入类型。
type EdgeDeviceAuthInfo struct {
	// 鉴权类型。支持密钥认证接入(SECRET)和证书认证接入(CERTIFICATES)两种方式。使用密钥认证接入方式(SECRET)填写secret字段，使用证书认证接入方式(CERTIFICATES)填写fingerprint字段，不填写auth_type默认为密钥认证接入方式(SECRET)

	AuthType *string `json:"auth_type,omitempty"`
	// 设备密钥，认证类型使用密钥认证接入(SECRET)可填写该字段。注意：NB设备密钥由于协议特殊性，只支持十六进制密钥接入；修改设备、查询设备及查询设备列表接口不返回该参数。

	Secret *string `json:"secret,omitempty"`
	// 证书指纹，认证类型使用证书认证接入(CERTIFICATES)可填写该字段，注册设备时不填写该字段则取第一次设备接入时的证书指纹。注意：指纹只能为40位十六进制字符串或者64位十六进制字符串；修改设备、查询设备及查询设备列表接口不返回该参数。

	Fingerprint *string `json:"fingerprint,omitempty"`
	// 指设备是否通过安全协议方式接入，默认值为true。 - true：通过安全协议方式接入。 - false：通过非安全协议方式接入。

	SecureAccess *bool `json:"secure_access,omitempty"`
	// 设备验证码的有效时间，单位：秒，默认值：0 若设备在有效时间内未接入物联网平台并激活，则平台会删除该设备的注册信息。若设置为“0”，则表示设备验证码不会失效（建议填写为“0”）。注意：只有注册设备接口或者修改设备接口修改timeout时返回该参数。

	Timeout *int32 `json:"timeout,omitempty"`
}

func (o EdgeDeviceAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeDeviceAuthInfo struct{}"
	}

	return strings.Join([]string{"EdgeDeviceAuthInfo", string(data)}, " ")
}
