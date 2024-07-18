package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryClientCaCertificateBody 客户端CA证书
type QueryClientCaCertificateBody struct {

	// ID
	Id *string `json:"id,omitempty"`

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

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o QueryClientCaCertificateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryClientCaCertificateBody struct{}"
	}

	return strings.Join([]string{"QueryClientCaCertificateBody", string(data)}, " ")
}
