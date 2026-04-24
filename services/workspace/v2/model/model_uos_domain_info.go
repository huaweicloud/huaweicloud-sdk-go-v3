package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UosDomainInfo 查询域控的配置信息响应。
type UosDomainInfo struct {

	// 认证配置id。
	AuthConfigId *string `json:"auth_config_id,omitempty"`

	// 域控id。
	Id *string `json:"id,omitempty"`

	Type *DomainType `json:"type,omitempty"`

	// 统信域控名称。
	DomainName *string `json:"domain_name,omitempty"`

	// 域管理员。
	Username *string `json:"username,omitempty"`

	// 域管平台地址。
	MainDcAddress *string `json:"main_dc_address,omitempty"`

	// 域管开放接口地址。
	OpenInterfaceAddress *string `json:"open_interface_address,omitempty"`

	// 域管开放接口域名。
	OpenInterfaceDomainName *string `json:"open_interface_domain_name,omitempty"`

	// 域管内部服务地址。
	InternalServiceAddress *string `json:"internal_service_address,omitempty"`

	// 客户端证书公钥id。
	AppCertId *string `json:"app_cert_id,omitempty"`

	// 客户端证书公钥有效期起始时间。
	AppCertStartTime *string `json:"app_cert_start_time,omitempty"`

	// 客户端证书公钥有效期结束时间。
	AppCertEndTime *string `json:"app_cert_end_time,omitempty"`

	// 服务端CA id。
	OpenapiCaCertId *string `json:"openapi_ca_cert_id,omitempty"`

	// 服务端CA有效期起始时间。
	OpenapiCaCertStartTime *string `json:"openapi_ca_cert_start_time,omitempty"`

	// 服务端CA有效期结束时间。
	OpenapiCaCertEndTime *string `json:"openapi_ca_cert_end_time,omitempty"`
}

func (o UosDomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UosDomainInfo struct{}"
	}

	return strings.Join([]string{"UosDomainInfo", string(data)}, " ")
}
