package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerResponseServerCertificate 服务端证书
type ShowServerResponseServerCertificate struct {

	// 证书 ID，CCM 服务中的certificate_id，证书在CCM中被删除后，该ID为空
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

	// 证书签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`
}

func (o ShowServerResponseServerCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerResponseServerCertificate struct{}"
	}

	return strings.Join([]string{"ShowServerResponseServerCertificate", string(data)}, " ")
}
