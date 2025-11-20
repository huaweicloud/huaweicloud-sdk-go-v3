package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerCertificateResponse Response Object
type ShowServerCertificateResponse struct {

	// 唯一标识ID
	ServerCertificateId *string `json:"server_certificate_id,omitempty"`

	// **参数说明**：证书通用名
	CommonName *string `json:"common_name,omitempty"`

	// **参数说明**：证书内容，PEM格式
	CertificatePem *string `json:"certificate_pem,omitempty"`

	// 证书生效日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	EffectiveDate *string `json:"effective_date,omitempty"`

	// 证书失效日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 证书所有者。
	CertificateOwner *string `json:"certificate_owner,omitempty"`

	// 证书颁发者。
	CertificateIssuer *string `json:"certificate_issuer,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowServerCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowServerCertificateResponse", string(data)}, " ")
}
