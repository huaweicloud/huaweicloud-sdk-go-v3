package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 修改云模式域名的请求体
type UpdateHostRequestBody struct {

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificatename *string `json:"certificatename,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]CloudWafServer `json:"server,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对于低于最低TLS版本的请求，将无法正常访问网站
	Tls *UpdateHostRequestBodyTls `json:"tls,omitempty"`

	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *UpdateHostRequestBodyCipher `json:"cipher,omitempty"`

	// 是否支持http2   - true：表示支持http2   - false：表示不支持http2
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 是否开启IPv6防护，仅专业版（原企业版）和铂金版（原旗舰版）支持IPv6防护。   - true：开启IPv6防护   - false：关闭IPV6防护
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// 是否使用独享ip   - true：使用独享ip   - false：不实用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 套餐付费模式，目前只支持prePaid预付款模式
	PaidType *string `json:"paid_type,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty"`

	CircuitBreaker *CircuitBreaker `json:"circuit_breaker,omitempty"`

	TimeoutConfig *TimeoutConfig `json:"timeout_config,omitempty"`
}

func (o UpdateHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostRequestBody", string(data)}, " ")
}

type UpdateHostRequestBodyTls struct {
	value string
}

type UpdateHostRequestBodyTlsEnum struct {
	TLS_V1_0 UpdateHostRequestBodyTls
	TLS_V1_1 UpdateHostRequestBodyTls
	TLS_V1_2 UpdateHostRequestBodyTls
	TLS_V1_3 UpdateHostRequestBodyTls
}

func GetUpdateHostRequestBodyTlsEnum() UpdateHostRequestBodyTlsEnum {
	return UpdateHostRequestBodyTlsEnum{
		TLS_V1_0: UpdateHostRequestBodyTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: UpdateHostRequestBodyTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: UpdateHostRequestBodyTls{
			value: "TLS v1.2",
		},
		TLS_V1_3: UpdateHostRequestBodyTls{
			value: "TLS v1.3",
		},
	}
}

func (c UpdateHostRequestBodyTls) Value() string {
	return c.value
}

func (c UpdateHostRequestBodyTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHostRequestBodyTls) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateHostRequestBodyCipher struct {
	value string
}

type UpdateHostRequestBodyCipherEnum struct {
	CIPHER_1       UpdateHostRequestBodyCipher
	CIPHER_2       UpdateHostRequestBodyCipher
	CIPHER_3       UpdateHostRequestBodyCipher
	CIPHER_4       UpdateHostRequestBodyCipher
	CIPHER_DEFAULT UpdateHostRequestBodyCipher
}

func GetUpdateHostRequestBodyCipherEnum() UpdateHostRequestBodyCipherEnum {
	return UpdateHostRequestBodyCipherEnum{
		CIPHER_1: UpdateHostRequestBodyCipher{
			value: "cipher_1",
		},
		CIPHER_2: UpdateHostRequestBodyCipher{
			value: "cipher_2",
		},
		CIPHER_3: UpdateHostRequestBodyCipher{
			value: "cipher_3",
		},
		CIPHER_4: UpdateHostRequestBodyCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: UpdateHostRequestBodyCipher{
			value: "cipher_default",
		},
	}
}

func (c UpdateHostRequestBodyCipher) Value() string {
	return c.value
}

func (c UpdateHostRequestBodyCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHostRequestBodyCipher) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
