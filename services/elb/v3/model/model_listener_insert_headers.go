package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListenerInsertHeaders struct {

	// **参数解释**：X-Forwarded-ELB-IP设为true可以将ELB实例的eip地址从报文的http头中带到后端服务器。
	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`

	// **参数解释**：X-Forwarded-Port设为true可以将ELB实例的监听端口从报文的http头中带到后端服务器。
	XForwardedPort *bool `json:"X-Forwarded-Port,omitempty"`

	// **参数解释**：X-Forwarded-For-Port设为true可以将客户端的源端口从报文的http头中带到后端服务器。
	XForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`

	// **参数解释**：X-Forwarded-Host设为true可以将客户请求头的X-Forwarded-Host设置为请求头的Host带到后端服务器。
	XForwardedHost *bool `json:"X-Forwarded-Host,omitempty"`

	// **参数解释**：X-Forwarded-Proto设为true可以将负载均衡器实例的监听协议通过报文的http头带到后端服务器。
	XForwardedProto *bool `json:"X-Forwarded-Proto,omitempty"`

	// **参数解释**：X-Real-IP设为true可以将客户端的IP通过报文的http头带到后端服务器。
	XRealIP *bool `json:"X-Real-IP,omitempty"`

	// **参数解释**：X-Forwarded-ELB-ID设为true可以将负载均衡器实例的ID通过报文的http头带到后端服务器。
	XForwardedELBID *bool `json:"X-Forwarded-ELB-ID,omitempty"`

	// **参数解释**：X-Forwarded-TLS-Certificate-ID设为true可以将负载均衡器实例的证书ID通过报文的http头带到后端服务器。
	XForwardedTLSCertificateID *bool `json:"X-Forwarded-TLS-Certificate-ID,omitempty"`

	// **参数解释**：X-Forwarded-TLS-Protocol设为true可以将负载均衡器实例的算法协议通过报文的http头带到后端服务器。
	XForwardedTLSProtocol *bool `json:"X-Forwarded-TLS-Protocol,omitempty"`

	// **参数解释**：X-Forwarded-TLS-Cipher设为true可以将负载均衡器实例的算法套件通过报文的http头带到后端服务器。
	XForwardedTLSCipher *bool `json:"X-Forwarded-TLS-Cipher,omitempty"`

	// **参数解释**：自定义X-Forwarded-TLS-Protocol头字段名称。  **约束限制**：只有当 X-Forwarded-TLS-Protocol 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedTLSProtocolAlias *string `json:"X-Forwarded-TLS-Protocol-alias,omitempty"`

	// **参数解释**：自定义X-Forwarded-TLS-Cipher头字段名称。  **约束限制**：只有当 X-Forwarded-TLS-Cipher 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedTLSCipherAlias *string `json:"X-Forwarded-TLS-Cipher-alias,omitempty"`

	// **参数解释**：处理X-Forwarded-For头字段的模式： 配置 append，将请求发送至后端服务之前把最后一跳 IP 加入X-Forwarded-For头字段； 配置 remove，请求发送至后端服务之前删除X-Forwarded-For标头，无论请求是否携带X-Forwarded-For头字段； 配置preserve，保留请求中已有的X-Forwarded-For标头；  **约束限制**：仅HTTP、HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段；  **取值范围**：append、remove、preserve  **默认取值**：append
	XForwardedForProcessingMode *ListenerInsertHeadersXForwardedForProcessingMode `json:"X-Forwarded-For-Processing-Mode,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-subjectdn头字段获取访问负载均衡实例客户端证书的所有者信息。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertSubjectdnEnable *bool `json:"X-Forwarded-Clientcert-subjectdn-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-subjectdn头字段名称。  **约束限制**：只有当 X-Forwarded-Clientcert-subjectdn-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertSubjectdnAlias *string `json:"X-Forwarded-Clientcert-subjectdn-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-issuerdn头字段获取访问负载均衡实例客户端证书的发行者信息。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertIssuerdnEnable *bool `json:"X-Forwarded-Clientcert-issuerdn-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-issuerdn头字段名称。  **约束限制**：只有当 X-Forwarded-Clientcert-issuerdn-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertIssuerdnAlias *string `json:"X-Forwarded-Clientcert-issuerdn-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-fingerprint头字段获取访问负载均衡实例客户端证书的指纹取值。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertFingerprintEnable *bool `json:"X-Forwarded-Clientcert-fingerprint-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-fingerprint头字段名称。  **约束限制**：只有当 X-Forwarded-Clientcert-fingerprint-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertFingerprintAlias *string `json:"X-Forwarded-Clientcert-fingerprint-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-clientverify头字段获取对访问负载均衡实例客户端证书的校验结果。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertClientverifyEnable *bool `json:"X-Forwarded-Clientcert-clientverify-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-clientverify头字段名称。  **约束限制**：只有当 X-Forwarded-Clientcert-clientverify-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertClientverifyAlias *string `json:"X-Forwarded-Clientcert-clientverify-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-serialnumber 头字段获取客户端证书的序列号信息。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertSerialnumberEnable *bool `json:"X-Forwarded-Clientcert-serialnumber-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-serialnumber头字段名称。  **约束限制**：只有当X-Forwarded-Clientcert-serialnumber-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertSerialnumberAlias *string `json:"X-Forwarded-Clientcert-serialnumber-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert 头字段获取客户端证书的内容。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertEnable *bool `json:"X-Forwarded-Clientcert-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert头字段名称。  **约束限制**：只有当X-Forwarded-Clientcert-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC、TLS协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertAlias *string `json:"X-Forwarded-Clientcert-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-ciphers 头字段获取客户端支持的TLS加密套件。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段。  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertCiphersEnable *bool `json:"X-Forwarded-Clientcert-ciphers-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-ciphers头字段名称。  **约束限制**：只有当X-Forwarded-Clientcert-ciphers-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertCiphersAlias *string `json:"X-Forwarded-Clientcert-ciphers-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Clientcert-end 头字段获取客户端证书的结束日期。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段。  **取值范围**：true、false  **默认取值**：false
	XForwardedClientcertEndEnable *bool `json:"X-Forwarded-Clientcert-end-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Clientcert-end头字段名称。  **约束限制**：只有当X-Forwarded-Clientcert-end-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC、TLS协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedClientcertEndAlias *string `json:"X-Forwarded-Clientcert-end-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Tls-Alpn-Protocol 头字段获取客户端和服务器之间ALPN（Application-Layer Protocol Negotiation）协商的应用层协议。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段。  **取值范围**：true、false  **默认取值**：false
	XForwardedTlsAlpnProtocolEnable *bool `json:"X-Forwarded-Tls-Alpn-Protocol-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Tls-Alpn-Protocol头字段名称。  **约束限制**：只有当X-Forwarded-Tls-Alpn-Protocol-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedTlsAlpnProtocolAlias *string `json:"X-Forwarded-Tls-Alpn-Protocol-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Tls-Sni 头字段获取客户端访问的sni证书的域名。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段。  **取值范围**：true、false  **默认取值**：false
	XForwardedTlsSniEnable *bool `json:"X-Forwarded-Tls-Sni-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Tls-Sni头字段名称。  **约束限制**：只有当X-Forwarded-Tls-Sni-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedTlsSniAlias *string `json:"X-Forwarded-Tls-Sni-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Tls-Ja3头字段获取访问负载均衡实例客户端的ja3指纹。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段。  **取值范围**：true、false  **默认取值**：false
	XForwardedTlsJa3Enable *bool `json:"X-Forwarded-Tls-Ja3-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Tls-Ja3头字段名称。  **约束限制**：只有当X-Forwarded-Tls-Ja3-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedTlsJa3Alias *string `json:"X-Forwarded-Tls-Ja3-alias,omitempty"`

	// **参数解释**：是否通过X-Forwarded-Tls-Ja4头字段获取访问负载均衡实例客户端的ja4指纹。  **约束限制**：仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段。  **取值范围**：true、false  **默认取值**：false
	XForwardedTlsJa4Enable *bool `json:"X-Forwarded-Tls-Ja4-enable,omitempty"`

	// **参数解释**：自定义X-Forwarded-Tls-Ja4头字段名称。  **约束限制**：只有当X-Forwarded-Tls-Ja4-enable 的值为 true 时，此值才会生效。仅HTTPS、TERMINATED_HTTPS、QUIC协议的监听器支持该字段  **取值范围**：1~40 个字符。支持字母 a~z、短划线（-）、下划线（_）和数字。  **默认取值**：无
	XForwardedTlsJa4Alias *string `json:"X-Forwarded-Tls-Ja4-alias,omitempty"`
}

func (o ListenerInsertHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerInsertHeaders struct{}"
	}

	return strings.Join([]string{"ListenerInsertHeaders", string(data)}, " ")
}

type ListenerInsertHeadersXForwardedForProcessingMode struct {
	value string
}

type ListenerInsertHeadersXForwardedForProcessingModeEnum struct {
	APPENDREMOVEPRESERVE ListenerInsertHeadersXForwardedForProcessingMode
}

func GetListenerInsertHeadersXForwardedForProcessingModeEnum() ListenerInsertHeadersXForwardedForProcessingModeEnum {
	return ListenerInsertHeadersXForwardedForProcessingModeEnum{
		APPENDREMOVEPRESERVE: ListenerInsertHeadersXForwardedForProcessingMode{
			value: "append、remove、preserve",
		},
	}
}

func (c ListenerInsertHeadersXForwardedForProcessingMode) Value() string {
	return c.value
}

func (c ListenerInsertHeadersXForwardedForProcessingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerInsertHeadersXForwardedForProcessingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
