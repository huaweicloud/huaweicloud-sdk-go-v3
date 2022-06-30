package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 修改独享模式域名的请求
type UpdatePremiumHostRequestBody struct {

	// 是否使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// https证书id，通过查询证书列表接口（ListCertificates）接口获取证书id
	Certificateid *string `json:"certificateid,omitempty"`

	// https证书名称，通过查询证书列表接口（ListCertificates）接口获取证书id
	Certificatename *string `json:"certificatename,omitempty"`

	// 支持最低的TLS版本
	Tls *UpdatePremiumHostRequestBodyTls `json:"tls,omitempty"`

	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM    cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH    cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *UpdatePremiumHostRequestBodyCipher `json:"cipher,omitempty"`

	// 独享模式特殊域名模式（仅特殊模式需要，如elb）
	Mode *string `json:"mode,omitempty"`

	// 是否锁定
	Locked *int32 `json:"locked,omitempty"`

	// 防护状态
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 接入状态
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 时间戳
	Timestamp *int32 `json:"timestamp,omitempty"`

	// 域名关联的组ID（仅特殊模式需要，如elb）
	PoolIds *[]string `json:"pool_ids,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`

	// 域名特殊标识
	Flag map[string]string `json:"flag,omitempty"`

	// 可扩展字段
	Extend map[string]string `json:"extend,omitempty"`

	CircuitBreaker *CircuitBreaker `json:"circuit_breaker,omitempty"`

	TimeoutConfig *TimeoutConfig `json:"timeout_config,omitempty"`
}

func (o UpdatePremiumHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostRequestBody", string(data)}, " ")
}

type UpdatePremiumHostRequestBodyTls struct {
	value string
}

type UpdatePremiumHostRequestBodyTlsEnum struct {
	TLS_V1_0 UpdatePremiumHostRequestBodyTls
	TLS_V1_1 UpdatePremiumHostRequestBodyTls
	TLS_V1_2 UpdatePremiumHostRequestBodyTls
	TLS_V1_3 UpdatePremiumHostRequestBodyTls
}

func GetUpdatePremiumHostRequestBodyTlsEnum() UpdatePremiumHostRequestBodyTlsEnum {
	return UpdatePremiumHostRequestBodyTlsEnum{
		TLS_V1_0: UpdatePremiumHostRequestBodyTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: UpdatePremiumHostRequestBodyTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: UpdatePremiumHostRequestBodyTls{
			value: "TLS v1.2",
		},
		TLS_V1_3: UpdatePremiumHostRequestBodyTls{
			value: "TLS v1.3",
		},
	}
}

func (c UpdatePremiumHostRequestBodyTls) Value() string {
	return c.value
}

func (c UpdatePremiumHostRequestBodyTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumHostRequestBodyTls) UnmarshalJSON(b []byte) error {
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

type UpdatePremiumHostRequestBodyCipher struct {
	value string
}

type UpdatePremiumHostRequestBodyCipherEnum struct {
	CIPHER_1       UpdatePremiumHostRequestBodyCipher
	CIPHER_2       UpdatePremiumHostRequestBodyCipher
	CIPHER_3       UpdatePremiumHostRequestBodyCipher
	CIPHER_4       UpdatePremiumHostRequestBodyCipher
	CIPHER_DEFAULT UpdatePremiumHostRequestBodyCipher
}

func GetUpdatePremiumHostRequestBodyCipherEnum() UpdatePremiumHostRequestBodyCipherEnum {
	return UpdatePremiumHostRequestBodyCipherEnum{
		CIPHER_1: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_1",
		},
		CIPHER_2: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_2",
		},
		CIPHER_3: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_3",
		},
		CIPHER_4: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_default",
		},
	}
}

func (c UpdatePremiumHostRequestBodyCipher) Value() string {
	return c.value
}

func (c UpdatePremiumHostRequestBodyCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumHostRequestBodyCipher) UnmarshalJSON(b []byte) error {
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
