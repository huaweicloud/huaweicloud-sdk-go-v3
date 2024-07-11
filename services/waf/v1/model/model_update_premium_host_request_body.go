package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePremiumHostRequestBody 修改独享模式域名的请求
type UpdatePremiumHostRequestBody struct {

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificatename *string `json:"certificatename,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]PremiumWafServer `json:"server,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对于低于最低TLS版本的请求，将无法正常访问网站
	Tls *UpdatePremiumHostRequestBodyTls `json:"tls,omitempty"`

	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_5：加密算法为AES128-SHA:AES256-SHA:AES128-SHA256:AES256-SHA256:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!EXPORT:!DES:!MD5:!PSK:!RC4:!DHE:@STRENGTH    cipher_6：加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *UpdatePremiumHostRequestBodyCipher `json:"cipher,omitempty"`

	// 独享模式特殊域名模式（仅特殊模式需要，如elb）
	Mode *string `json:"mode,omitempty"`

	// 预留参数，用于后期设计冻结域名，解锁域名功能，目前暂不支持
	Locked *int32 `json:"locked,omitempty"`

	// 域名防护状态：  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 时间戳
	Timestamp *int32 `json:"timestamp,omitempty"`

	// 特殊模式域名所属独享引擎组（仅特殊模式需要，如elb）
	PoolIds *[]string `json:"pool_ids,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`

	CircuitBreaker *CircuitBreaker `json:"circuit_breaker,omitempty"`

	TimeoutConfig *TimeoutConfig `json:"timeout_config,omitempty"`

	Flag *HostFlag `json:"flag,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`
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

type UpdatePremiumHostRequestBodyCipher struct {
	value string
}

type UpdatePremiumHostRequestBodyCipherEnum struct {
	CIPHER_1       UpdatePremiumHostRequestBodyCipher
	CIPHER_2       UpdatePremiumHostRequestBodyCipher
	CIPHER_3       UpdatePremiumHostRequestBodyCipher
	CIPHER_4       UpdatePremiumHostRequestBodyCipher
	CIPHER_5       UpdatePremiumHostRequestBodyCipher
	CIPHER_6       UpdatePremiumHostRequestBodyCipher
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
		CIPHER_5: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_5",
		},
		CIPHER_6: UpdatePremiumHostRequestBodyCipher{
			value: "cipher_6",
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
