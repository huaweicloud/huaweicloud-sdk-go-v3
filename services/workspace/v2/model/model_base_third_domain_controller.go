package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseThirdDomainController struct {

	// 域管理员。
	Username *string `json:"username,omitempty"`

	// 域管理员密码。
	UserPassword *string `json:"user_password,omitempty"`

	// 域管平台地址。
	MainDcAddress *string `json:"main_dc_address,omitempty"`

	// 域管开放接口地址。
	OpenInterfaceAddress *string `json:"open_interface_address,omitempty"`

	// 域管开放接口域名。
	OpenInterfaceDomainName *string `json:"open_interface_domain_name,omitempty"`

	// 域管内部服务地址。
	InternalServiceAddress *string `json:"internal_service_address,omitempty"`

	// 客户端证书公钥。
	AppCert *string `json:"app_cert,omitempty"`

	// 客户端证书私钥。
	AppCertKey *string `json:"app_cert_key,omitempty"`

	// 服务端CA。
	OpenapiCaCert *string `json:"openapi_ca_cert,omitempty"`
}

func (o BaseThirdDomainController) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseThirdDomainController struct{}"
	}

	return strings.Join([]string{"BaseThirdDomainController", string(data)}, " ")
}
