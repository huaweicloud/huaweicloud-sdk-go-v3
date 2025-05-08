package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortProtocolsEntity 实例支持的连接方式及其连接地址。
type PortProtocolsEntity struct {

	// 实例是否支持内网PLAINTEXT访问接入方式。  - true：实例支持内网PLAINTEXT访问方式接入方式。  - false：实例不支持内网PLAINTEXT访问接入方式。
	PrivatePlainEnable *bool `json:"private_plain_enable,omitempty"`

	// kafka内网PLAINTEXT接入方式连接地址。
	PrivatePlainAddress *string `json:"private_plain_address,omitempty"`

	// 内网明文连接域名
	PrivatePlainDomainName *string `json:"private_plain_domain_name,omitempty"`

	// 实例是否支持内网SASL_SSL访问接入方式。  - true：实例支持内网SASL_SSL访问方式接入方式。  - false：实例不支持内网SASL_SSL访问接入方式。
	PrivateSaslSslEnable *bool `json:"private_sasl_ssl_enable,omitempty"`

	// kafka内网SASL_SSL接入方式连接地址。
	PrivateSaslSslAddress *string `json:"private_sasl_ssl_address,omitempty"`

	// 内网SASL_SSL连接域名
	PrivateSaslSslDomainName *string `json:"private_sasl_ssl_domain_name,omitempty"`

	// 实例是否支持内网SASL_PLAINTEXT访问接入方式。  - true，实例支持内网SASL_PLAINTEXT访问方式接入方式。  - false，实例不支持内网SASL_PLAINTEXT访问接入方式。
	PrivateSaslPlaintextEnable *bool `json:"private_sasl_plaintext_enable,omitempty"`

	// kafka内网SASL_PLAINTEXT接入方式连接地址。
	PrivateSaslPlaintextAddress *string `json:"private_sasl_plaintext_address,omitempty"`

	// 内网SASL_PLAINTEXT连接域名
	PrivateSaslPlaintextDomainName *string `json:"private_sasl_plaintext_domain_name,omitempty"`

	// 实例是否支持公网PLAINTEXT访问接入方式。  - true，实例支持公网PLAINTEXT访问方式接入方式。  - false，实例不支持公网PLAINTEXT访问接入方式。
	PublicPlainEnable *bool `json:"public_plain_enable,omitempty"`

	// kafka公网PLAINTEXT接入方式连接地址。
	PublicPlainAddress *string `json:"public_plain_address,omitempty"`

	// 公网明文连接域名
	PublicPlainDomainName *string `json:"public_plain_domain_name,omitempty"`

	// 实例是否支持公网SASL_SSL访问接入方式。  - true，实例支持内网SASL_SSL访问方式接入方式。  - false，实例不支持公网SASL_SSL访问接入方式。
	PublicSaslSslEnable *bool `json:"public_sasl_ssl_enable,omitempty"`

	// kafka公网SASL_SSL接入方式连接地址。
	PublicSaslSslAddress *string `json:"public_sasl_ssl_address,omitempty"`

	// 公网SASL_SSL连接域名
	PublicSaslSslDomainName *string `json:"public_sasl_ssl_domain_name,omitempty"`

	// 实例是否支持公网SASL_PLAINTEXT访问接入方式。  - true，实例支持公网SASL_PLAINTEXT访问方式接入方式。  - false，实例不支持公网SASL_PLAINTEXT访问接入方式。
	PublicSaslPlaintextEnable *bool `json:"public_sasl_plaintext_enable,omitempty"`

	// kafka公网SASL_PLAINTEXT接入方式连接地址。
	PublicSaslPlaintextAddress *string `json:"public_sasl_plaintext_address,omitempty"`

	// 公网SASL_PLAINTEXT连接域名
	PublicSaslPlaintextDomainName *string `json:"public_sasl_plaintext_domain_name,omitempty"`
}

func (o PortProtocolsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortProtocolsEntity struct{}"
	}

	return strings.Join([]string{"PortProtocolsEntity", string(data)}, " ")
}
