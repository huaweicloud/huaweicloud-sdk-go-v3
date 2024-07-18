package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClientCaCertificateResponse Response Object
type CheckClientCaCertificateResponse struct {

	// 证书名
	Name *string `json:"name,omitempty"`

	// 颁发者
	Issuer *string `json:"issuer,omitempty"`

	// 主体
	Subject *string `json:"subject,omitempty"`

	// 序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 过期时间
	ExpirationTime *sdktime.SdkTime `json:"expiration_time,omitempty"`

	// 客户端 CA 证书签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o CheckClientCaCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClientCaCertificateResponse struct{}"
	}

	return strings.Join([]string{"CheckClientCaCertificateResponse", string(data)}, " ")
}
