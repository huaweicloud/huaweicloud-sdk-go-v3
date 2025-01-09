package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCsrResponse Response Object
type CreateCsrResponse struct {

	// CSR的ID。
	Id *string `json:"id,omitempty"`

	// 自定义CSR名称。
	Name *string `json:"name,omitempty"`

	// 域名。
	DomainName *string `json:"domain_name,omitempty"`

	// CSR绑定的附加域名。
	Sans *string `json:"sans,omitempty"`

	// 密钥算法的类型。取值如下 - RSA_2048 - RSA_3072 - RSA_4096 - EC_P256 - EC_P384 - SM2
	PrivateKeyAlgo *string `json:"private_key_algo,omitempty"`

	// CSR用途。取值如下： - PERSONAL：个人证书 - ENTERPRISE：企业证书
	Usage *string `json:"usage,omitempty"`

	// 国家，当“usage”取值为“ENTERPRISE”时，本参数必填。
	CompanyCountry *string `json:"company_country,omitempty"`

	// 省份，当“usage”取值为“ENTERPRISE”时，本参数必填。
	CompanyProvince *string `json:"company_province,omitempty"`

	// 城市，当“usage”取值为“ENTERPRISE”时，本参数必填。
	CompanyCity *string `json:"company_city,omitempty"`

	// 公司名称，当“usage”取值为“ENTERPRISE”时，本参数必填。
	CompanyName    *string `json:"company_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCsrResponse struct{}"
	}

	return strings.Join([]string{"CreateCsrResponse", string(data)}, " ")
}
