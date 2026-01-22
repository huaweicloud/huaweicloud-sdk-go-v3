package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortProtocol **参数解释**： 设置Kafka实例的接入方式。PLAINTEXT表示明文接入，SASL_SSL[或者SASL_PLAINTEXT](tag:hws,hws_hk,hws_eu)表示密文接入。 **约束限制**： 内网访问不支持关闭，明文接入和密文接入至少开启一个。 [跨VPC访问的安全协议等于内网访问的安全协议，若内网同时开启了密文访问和明文访问，则跨VPC访问的安全协议会优先使用密文访问的安全协议。](tag:hws,hws_hk,hws_eu)
type PortProtocol struct {

	// **参数解释**： 是否开启内网明文访问连接方式。 **约束限制**： 不涉及。 **取值范围**： - true：开启内网明文访问连接方式，连接地址：ip:9092，访问协议PLAINTEXT。 - false：关闭内网明文访问。 **默认取值**： false。
	PrivatePlainEnable *bool `json:"private_plain_enable,omitempty"`

	// **参数解释**： 是否开启安全协议为SASL_SSL的内网密文接入方式。 **约束限制**： [private_sasl_ssl_enable和private_sasl_plaintext_enable不能同时为true。](tag:hws,hws_eu,hws_hk,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[不涉及。](tag:hcs,ocb,hws_ocb) **取值范围**： - true：开启安全协议为SASL_SSL的内网密文接入方式。          - false：关闭安全协议为SASL_SSL的内网接入方式。 **默认取值**： false。
	PrivateSaslSslEnable *bool `json:"private_sasl_ssl_enable,omitempty"`

	// **参数解释**： 是否开启安全协议为SASL_PLAINTEXT的内网密文接入方式。 **约束限制**： [private_sasl_plaintext_enable和private_sasl_ssl_enable不能同时为true。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[华为云Stack不支持此参数。](tag:hcs) **取值范围**： - true：开启安全协议为SASL_PLAINTEXT的内网密文接入方式，连接地址：ip:9093，访问协议SASL_PLAINTEXT。     - false：关闭安全协议为SASL_PLAINTEXT的内网密文接入方式。 **默认取值**： false。
	PrivateSaslPlaintextEnable *bool `json:"private_sasl_plaintext_enable,omitempty"`

	// **参数解释**： 是否开启公网明文访问连接方式。 **约束限制**： 开启公网明文接入前，需要先开启公网访问功能。 **取值范围**： - true：开启公网明文访问连接方式，连接地址：ip:9094，访问协议PLAINTEXT。     - false：关闭公网明文接入方式。 **默认取值**： false。
	PublicPlainEnable *bool `json:"public_plain_enable,omitempty"`

	// **参数解释**： 是否开启安全协议为SASL_SSL的公网密文接入。 **约束限制**： [public_sasl_ssl_enable和public_sasl_plaintext_enable不能同时为true。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[不涉及。](tag:hcs) 为true时，需要实例开启公网。 **取值范围**： - true：开启安全协议为SASL_SSL的公网密文接入方式，连接地址：ip:9095，访问协议：SASL_SSL。 - false：关闭安全协议为SASL_SSL的公网密文接入方式。 **默认取值**： false。
	PublicSaslSslEnable *bool `json:"public_sasl_ssl_enable,omitempty"`

	// **参数解释**： 是否开启安全协议为SASL_PLAINTEXT的公网密文接入方式。 **约束限制**： [public_sasl_plaintext_enable和public_sasl_ssl_enable不能同时为true。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[华为云Stack不支持此参数。](tag:hcs) 为true时，需要实例开启公网。 **取值范围**： - true：开启安全协议为SASL_PLAINTEXT的公网密文接入方式，连接地址：ip:9095，访问协议：SASL_PLAINTEXT。           - false：关闭安全协议为SASL_PLAINTEXT的公网密文接入方式。 **默认取值**： false。
	PublicSaslPlaintextEnable *bool `json:"public_sasl_plaintext_enable,omitempty"`
}

func (o PortProtocol) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortProtocol struct{}"
	}

	return strings.Join([]string{"PortProtocol", string(data)}, " ")
}
