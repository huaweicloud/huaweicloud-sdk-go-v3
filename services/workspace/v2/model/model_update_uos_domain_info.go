package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUosDomainInfo 更新域控的配置信息请求。
type UpdateUosDomainInfo struct {

	// 域管理员。
	Username string `json:"username"`

	// 域管理员密码。
	UserPassword string `json:"user_password"`

	// 域管平台地址。
	MainDcAddress string `json:"main_dc_address"`

	// 域管开放接口地址。
	OpenInterfaceAddress string `json:"open_interface_address"`

	// 域管开放接口域名。
	OpenInterfaceDomainName *string `json:"open_interface_domain_name,omitempty"`

	// 域管内部服务地址。
	InternalServiceAddress string `json:"internal_service_address"`

	// 客户端证书公钥。
	AppCert string `json:"app_cert"`

	// 客户端证书私钥。
	AppCertKey string `json:"app_cert_key"`

	// 服务端CA。
	OpenapiCaCert string `json:"openapi_ca_cert"`

	// 域控id。
	Id *string `json:"id,omitempty"`

	// 客户端证书公钥id。需更新客户端证书时，必传。
	AppCertId *string `json:"app_cert_id,omitempty"`

	// 服务端CA id。需更新服务端CA时，必传。
	OpenapiCaCertId *string `json:"openapi_ca_cert_id,omitempty"`
}

func (o UpdateUosDomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUosDomainInfo struct{}"
	}

	return strings.Join([]string{"UpdateUosDomainInfo", string(data)}, " ")
}
