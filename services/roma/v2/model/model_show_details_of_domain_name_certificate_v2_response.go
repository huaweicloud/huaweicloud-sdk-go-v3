package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDetailsOfDomainNameCertificateV2Response struct {

	// 证书域名
	CommonName *string `json:"common_name,omitempty" xml:"common_name"`

	// SAN域名
	San *[]string `json:"san,omitempty" xml:"san"`

	// 证书版本
	Version *int32 `json:"version,omitempty" xml:"version"`

	// 公司、组织
	Organization *[]string `json:"organization,omitempty" xml:"organization"`

	// 部门
	OrganizationalUnit *[]string `json:"organizational_unit,omitempty" xml:"organizational_unit"`

	// 城市
	Locality *[]string `json:"locality,omitempty" xml:"locality"`

	// 省份
	State *[]string `json:"state,omitempty" xml:"state"`

	// 国家
	Country *[]string `json:"country,omitempty" xml:"country"`

	// 证书有效期起始时间
	NotBefore *sdktime.SdkTime `json:"not_before,omitempty" xml:"not_before"`

	// 证书有效期截止时间
	NotAfter *sdktime.SdkTime `json:"not_after,omitempty" xml:"not_after"`

	// 序列号
	SerialNumber *string `json:"serial_number,omitempty" xml:"serial_number"`

	// 颁发者
	Issuer *[]string `json:"issuer,omitempty" xml:"issuer"`

	// 签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty" xml:"signature_algorithm"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowDetailsOfDomainNameCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfDomainNameCertificateV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfDomainNameCertificateV2Response", string(data)}, " ")
}
