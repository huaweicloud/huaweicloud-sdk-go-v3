package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PortProtocolsEntity struct {

	// **参数解释**： 实例是否支持内网PLAINTEXT访问接入方式。 **取值范围**： - true：实例支持内网PLAINTEXT访问方式接入方式。 - false：实例不支持内网PLAINTEXT访问接入方式。
	PrivatePlainEnable *bool `json:"private_plain_enable,omitempty"`

	// **参数解释**： Kafka内网PLAINTEXT接入方式连接地址。 **取值范围**： 不涉及
	PrivatePlainAddress *string `json:"private_plain_address,omitempty"`

	// **参数解释**： 内网明文连接域名。 **取值范围**： 不涉及
	PrivatePlainDomainName *string `json:"private_plain_domain_name,omitempty"`

	// **参数解释**： 实例是否支持内网SASL_SSL访问接入方式。 **取值范围**： - true：实例支持内网SASL_SSL访问方式接入方式。 - false：实例不支持内网SASL_SSL访问接入方式。
	PrivateSaslSslEnable *bool `json:"private_sasl_ssl_enable,omitempty"`

	// **参数解释**： Kafka内网SASL_SSL接入方式连接地址。 **取值范围**： 不涉及
	PrivateSaslSslAddress *string `json:"private_sasl_ssl_address,omitempty"`

	// **参数解释**： 内网SASL_SSL连接域名。 **取值范围**： 不涉及
	PrivateSaslSslDomainName *string `json:"private_sasl_ssl_domain_name,omitempty"`

	// **参数解释**： 实例是否支持内网SASL_PLAINTEXT访问接入方式。 **取值范围**： - true：实例支持内网SASL_PLAINTEXT访问方式接入方式。 - false：实例不支持内网SASL_PLAINTEXT访问接入方式。
	PrivateSaslPlaintextEnable *bool `json:"private_sasl_plaintext_enable,omitempty"`

	// **参数解释**： Kafka内网SASL_PLAINTEXT接入方式连接地址。 **取值范围**： 不涉及
	PrivateSaslPlaintextAddress *string `json:"private_sasl_plaintext_address,omitempty"`

	// **参数解释**： 内网SASL_PLAINTEXT连接域名。 **取值范围**： 不涉及
	PrivateSaslPlaintextDomainName *string `json:"private_sasl_plaintext_domain_name,omitempty"`

	// **参数解释**： 实例是否支持公网PLAINTEXT访问接入方式。 **取值范围**： - true：实例支持公网PLAINTEXT访问方式接入方式。 - false：实例不支持公网PLAINTEXT访问接入方式。
	PublicPlainEnable *bool `json:"public_plain_enable,omitempty"`

	// **参数解释**： Kafka公网PLAINTEXT接入方式连接地址。 **取值范围**： 不涉及
	PublicPlainAddress *string `json:"public_plain_address,omitempty"`

	// **参数解释**： 公网明文连接域名。 **取值范围**： 不涉及
	PublicPlainDomainName *string `json:"public_plain_domain_name,omitempty"`

	// **参数解释**： 实例是否支持公网SASL_SSL访问接入方式。 **取值范围**： - true：实例支持内网SASL_SSL访问方式接入方式。 - false：实例不支持公网SASL_SSL访问接入方式。
	PublicSaslSslEnable *bool `json:"public_sasl_ssl_enable,omitempty"`

	// **参数解释**： Kafka公网SASL_SSL接入方式连接地址。 **取值范围**： 不涉及
	PublicSaslSslAddress *string `json:"public_sasl_ssl_address,omitempty"`

	// **参数解释**： 公网SASL_SSL连接域名。 **取值范围**： 不涉及
	PublicSaslSslDomainName *string `json:"public_sasl_ssl_domain_name,omitempty"`

	// **参数解释**： 实例是否支持公网SASL_PLAINTEXT访问接入方式。 **取值范围**： - true：实例支持公网SASL_PLAINTEXT访问方式接入方式。 - false：实例不支持公网SASL_PLAINTEXT访问接入方式。
	PublicSaslPlaintextEnable *bool `json:"public_sasl_plaintext_enable,omitempty"`

	// **参数解释**： Kafka公网SASL_PLAINTEXT接入方式连接地址。 **取值范围**： 不涉及
	PublicSaslPlaintextAddress *string `json:"public_sasl_plaintext_address,omitempty"`

	// **参数解释**： 公网SASL_PLAINTEXT连接域名。 **取值范围**： 不涉及
	PublicSaslPlaintextDomainName *string `json:"public_sasl_plaintext_domain_name,omitempty"`
}

func (o PortProtocolsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortProtocolsEntity struct{}"
	}

	return strings.Join([]string{"PortProtocolsEntity", string(data)}, " ")
}
