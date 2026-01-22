package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInsertHeaderConfig **参数解释**：要添加的请求头参数。  **约束限制**：不涉及
type UpdateInsertHeaderConfig struct {

	// **参数解释**：请求头参数名。  **约束限制**：不能是以下字符： connection、upgrade、content-length、transfer-encoding、keep-alive、te、host、cookie、remoteip、authority、x-forwarded-host、x-forwarded-for、x-forwarded-for-port、x-forwarded-tls-certificate-id、x-forwarded-tls-protocol、x-forwarded-tls-cipher、x-forwarded-elb-ip、x-forwarded-port、x-forwarded-elb-id、x-forwarded-elb-vip、x-real-ip、x-forwarded-proto、x-nuwa-trace-ne-in、x-nuwa-trace-ne-out。  **取值范围**：1-40个字符，字母a-z（不区分大小写）、数字，短划线-和下划线_。  **默认取值**：不涉及
	Key string `json:"key"`

	// **参数解释**：请求头参数类别。  **约束限制**：不涉及  **取值范围**： - USER_DEFINED：用户自定义请求头。 - REFERENCE_HEADER：引用请求头字段中的某一个字段。 - SYSTEM_DEFINED：表示请求头参数只能使用系统定义的一系列固定值，取值列表见value字段描述。  **默认取值**：不涉及
	ValueType string `json:"value_type"`

	// **参数解释**：请求头参数的值。  **约束限制**：当value_type为SYSTEM_DEFINED时，value只可从CLIENT-PORT, CLIENT-IP, ELB-PROTOCOL, ELB-ID,         ELB-PORT, ELB-EIP, ELB-VIP, TLS-PROTOCOL, TLS-CIPHER, CLIENTCERT-SUBJECTDN, CLIENTCERT-ISSUERDN,         CLIENTCERT-FINGERPRINT, CLIENTCERT-CLIENTVERIFY, CLIENTCERT-SERIALNUMBER,         CLIENTCERT, CLIENTCERT-CIPHERS, CLIENTCERT-END, TLS-ALPN-PROTOCOL, TLS-SNI,         TLS-JA3, TLS-JA4中取值。{TLS-PROTOCOL, TLS-CIPHER, CLIENTCERT-SUBJECTDN, CLIENTCERT-ISSUERDN,         CLIENTCERT-FINGERPRINT, CLIENTCERT-CLIENTVERIFY, CLIENTCERT-SERIALNUMBER,         CLIENTCERT, CLIENTCERT-CIPHERS, CLIENTCERT-END, TLS-ALPN-PROTOCOL, TLS-SNI,         TLS-JA3, TLS-JA4}只有HTTPS、quic协议类型的监听器支持  **取值范围**：1-128个字符，支持ascii码值32<=ch<=127范围内可打印字符，*和英文问号?。不能以空格开头或结尾，不允许设置值为：connection、upgrade、content-length、transfer-encoding、keep-alive、te、host、cookie、remoteip、authority、transparent、x-forwarded-host、x-forwarded-for、x-forwarded-for-port、x-forwarded-tls-certificate-id、 x-forwarded-tls-protocol、x-forwarded-tls-cipher、x-forwarded-port、x-real-ip、x-forwarded-proto、x-forwarded-elb-ip、 x-forwarded-elb-id、x-forwarded-elb-vip、x-nuwa-trace-ne-in、x-nuwa-trace-ne-out。  **默认取值**：不涉及
	Value string `json:"value"`
}

func (o UpdateInsertHeaderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInsertHeaderConfig struct{}"
	}

	return strings.Join([]string{"UpdateInsertHeaderConfig", string(data)}, " ")
}
