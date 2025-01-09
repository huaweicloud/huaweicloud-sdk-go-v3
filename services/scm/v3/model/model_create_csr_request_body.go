package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCsrRequestBody struct {

	// 自定义CSR名称。
	Name string `json:"name"`

	// CSR绑定的域名。如果您想在提交证书申请时使用该CSR，必须确保证书绑定域名包含此处设置的域名。
	DomainName string `json:"domain_name"`

	// CSR绑定的附加域名。
	Sans *string `json:"sans,omitempty"`

	// 私钥算法。取值如下： - RSA_2048 - RSA_3072 - RSA_4096 - EC_P256 - EC_P384 - SM2
	PrivateKeyAlgo string `json:"private_key_algo"`

	// CSR用途。取值如下： - PERSONAL：个人证书 - ENTERPRISE：企业证书
	Usage string `json:"usage"`

	// 国家，当“usage”取值为“ENTERPRISE”时，本参数必填。取值示例：CN。
	CompanyCountry *string `json:"company_country,omitempty"`

	// 省份，当“usage”取值为“ENTERPRISE”时，本参数必填。取值示例：北京市。
	CompanyProvince *string `json:"company_province,omitempty"`

	// 城市，用途为企业证书场景下必填。当“usage”取值为“ENTERPRISE”时，本参数必填。取值示例：北京市。
	CompanyCity *string `json:"company_city,omitempty"`

	// 公司名称，用途为企业证书场景下必填。当“usage”取值为“ENTERPRISE”时，本参数必填。
	CompanyName *string `json:"company_name,omitempty"`
}

func (o CreateCsrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCsrRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCsrRequestBody", string(data)}, " ")
}
