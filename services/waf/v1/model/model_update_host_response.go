package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type UpdateHostResponse struct {
	// 域名id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 创建的云模式防护域名

	Hostname *string `json:"hostname,omitempty"`
	// 账户id

	Domainid *string `json:"domainid,omitempty"`
	// cname前缀

	AccessCode *string `json:"access_code,omitempty"`
	// 后端协议类型

	Protocol *string `json:"protocol,omitempty"`
	// 源站信息

	Server *[]CloudWafServer `json:"server,omitempty"`
	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名，通过查询证书列表接口（ListCertificates）接口获取证书id

	Certificatename *string `json:"certificatename,omitempty"`
	// 是否开启了代理

	Proxy *bool `json:"proxy,omitempty"`
	// 锁定状态,默认为0

	Locked *int32 `json:"locked,omitempty"`
	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测

	ProtectStatus *int32 `json:"protect_status,omitempty"`
	// 接入状态

	AccessStatus *int32 `json:"access_status,omitempty"`
	// 创建防护域名的时间

	Timestamp *int64 `json:"timestamp,omitempty"`
	// ssl协议版本

	Tls *UpdateHostResponseTls `json:"tls,omitempty"`
	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM    cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH    cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM

	Cipher *UpdateHostResponseCipher `json:"cipher,omitempty"`
	// 企业项目ID

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`
	// 域名名称

	WebTag *bool `json:"web_tag,omitempty"`

	Flag *Flag `json:"flag,omitempty"`
	// 是否使用独享ip

	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`
	// 域名描述

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostResponse", string(data)}, " ")
}

type UpdateHostResponseTls struct {
	value string
}

type UpdateHostResponseTlsEnum struct {
	TLS_V1_0 UpdateHostResponseTls
	TLS_V1_1 UpdateHostResponseTls
	TLS_V1_2 UpdateHostResponseTls
	TLS_V1_3 UpdateHostResponseTls
}

func GetUpdateHostResponseTlsEnum() UpdateHostResponseTlsEnum {
	return UpdateHostResponseTlsEnum{
		TLS_V1_0: UpdateHostResponseTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: UpdateHostResponseTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: UpdateHostResponseTls{
			value: "TLS v1.2",
		},
		TLS_V1_3: UpdateHostResponseTls{
			value: "TLS v1.3",
		},
	}
}

func (c UpdateHostResponseTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHostResponseTls) UnmarshalJSON(b []byte) error {
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

type UpdateHostResponseCipher struct {
	value string
}

type UpdateHostResponseCipherEnum struct {
	CIPHER_1       UpdateHostResponseCipher
	CIPHER_2       UpdateHostResponseCipher
	CIPHER_3       UpdateHostResponseCipher
	CIPHER_4       UpdateHostResponseCipher
	CIPHER_DEFAULT UpdateHostResponseCipher
}

func GetUpdateHostResponseCipherEnum() UpdateHostResponseCipherEnum {
	return UpdateHostResponseCipherEnum{
		CIPHER_1: UpdateHostResponseCipher{
			value: "cipher_1",
		},
		CIPHER_2: UpdateHostResponseCipher{
			value: "cipher_2",
		},
		CIPHER_3: UpdateHostResponseCipher{
			value: "cipher_3",
		},
		CIPHER_4: UpdateHostResponseCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: UpdateHostResponseCipher{
			value: "cipher_default",
		},
	}
}

func (c UpdateHostResponseCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHostResponseCipher) UnmarshalJSON(b []byte) error {
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
