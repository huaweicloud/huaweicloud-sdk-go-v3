package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortProtocol 设置Kafka实例的接入方式。PLAINTEXT表示明文接入，SASL_SSL或者SASL_PLAINTEEXT表示密文接入。 内网访问不支持关闭，明文接入和密文接入至少开启一个。 跨VPC访问的安全协议等于内网访问的安全协议，若内网同时开启了密文访问和明文访问，则跨VPC访问的安全协议会优先使用密文访问的安全协议。
type PortProtocol struct {

	// 是否开启内网明文访问连接方式。  取值范围：   - true: 开启内网明文访问连接方式，连接地址：ip:9092，访问协议PLAINTEXT。   - false: 关闭内网明文访问。  默认为false。
	PrivatePlainEnable *bool `json:"private_plain_enable,omitempty"`

	// 是否开启安全协议为SASL_SSL的内网密文接入方式。  取值范围：   - true: 开启安全协议为SASL_SSL的内网密文接入方式。           private_sasl_ssl_enable和private_sasl_plaintext_enable不能同时为true。   - false: 关闭安全协议为SASL_SSL的内网接入方式。  默认为false。
	PrivateSaslSslEnable *bool `json:"private_sasl_ssl_enable,omitempty"`

	// 是否开启安全协议为SASL_PLAINTEXT的内网密文接入方式。  取值范围：   - true: 开启安全协议为SASL_PLAINTEXT的内网密文接入方式，连接地址：ip:9093，访问协议SASL_PLAINTEXT。           private_sasl_plaintext_enable和private_sasl_ssl_enable不能同时为true。   - false: 关闭安全协议为SASL_PLAINTEXT的内网密文接入方式。  默认为false。
	PrivateSaslPlaintextEnable *bool `json:"private_sasl_plaintext_enable,omitempty"`

	// 是否开启公网明文访问连接方式。  取值范围：   - true: 开启公网明文访问连接方式，连接地址：ip:9094，访问协议PLAINTEXT。           开启公网明文接入前，需要先开启公网访问功能。   - false: 关闭公网明文接入方式。  默认为false。
	PublicPlainEnable *bool `json:"public_plain_enable,omitempty"`

	// 是否开启安全协议为SASL_SSL的公网密文接入。  取值范围：   - true: 开启安全协议为SASL_SSL的公网密文接入方式，连接地址：ip:9095，访问协议：SASL_SSL。           public_sasl_ssl_enable和public_sasl_plaintext_enable不能同时为true。           为true时，需要实例开启公网。   - false: 关闭安全协议为SASL_SSL的公网密文接入方式。  默认为false。
	PublicSaslSslEnable *bool `json:"public_sasl_ssl_enable,omitempty"`

	// 是否开启安全协议为SASL_PLAINTEXT的公网密文接入方式。  取值范围：   - true: 开启安全协议为SASL_PLAINTEXT的公网密文接入方式，连接地址：ip:9095，访问协议：SASL_PLAINTEXT。           public_sasl_plaintext_enable和public_sasl_ssl_enable不能同时为true。           为true时，需要实例开启公网。   - false: 关闭安全协议为SASL_PLAINTEXT的公网密文接入方式。  默认为false。
	PublicSaslPlaintextEnable *bool `json:"public_sasl_plaintext_enable,omitempty"`
}

func (o PortProtocol) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortProtocol struct{}"
	}

	return strings.Join([]string{"PortProtocol", string(data)}, " ")
}
