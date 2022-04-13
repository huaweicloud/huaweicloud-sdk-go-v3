package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ShowPremiumHostResponse struct {
	// 域名id

	Id *string `json:"id,omitempty"`
	// 防护域名初始绑定的策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id

	Policyid *string `json:"policyid,omitempty"`
	// 创建的云模式防护域名

	Hostname *string `json:"hostname,omitempty"`
	// 用户Domain ID

	Domainid *string `json:"domainid,omitempty"`
	// 用户的project_id

	ProjectId *string `json:"project_id,omitempty"`
	// cname前缀

	AccessCode *string `json:"access_code,omitempty"`
	// http协议类型

	Protocol *string `json:"protocol,omitempty"`
	// 源站信息

	Server *[]PremiumWafServer `json:"server,omitempty"`
	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数

	Certificatename *string `json:"certificatename,omitempty"`
	// 支持最低的TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本

	Tls *ShowPremiumHostResponseTls `json:"tls,omitempty"`
	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM    cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH    cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM

	Cipher *ShowPremiumHostResponseCipher `json:"cipher,omitempty"`
	// 是否开启了代理

	Proxy *bool `json:"proxy,omitempty"`
	// 锁定状态

	Locked *int32 `json:"locked,omitempty"`
	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测

	ProtectStatus *int32 `json:"protect_status,omitempty"`
	// 接入状态

	AccessStatus *int32 `json:"access_status,omitempty"`
	// 创建防护域名的时间

	Timestamp *int64 `json:"timestamp,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`
	// 可扩展属性

	Extend map[string]string `json:"extend,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`
	// 域名特殊标记

	Flag map[string]string `json:"flag,omitempty"`
	// 独享模式特殊域名模式（仅特殊模式需要，如elb）

	Mode *string `json:"mode,omitempty"`
	// 域名关联的组ID（仅特殊模式需要，如elb）

	PoolIds        *[]string `json:"pool_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowPremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPremiumHostResponse struct{}"
	}

	return strings.Join([]string{"ShowPremiumHostResponse", string(data)}, " ")
}

type ShowPremiumHostResponseTls struct {
	value string
}

type ShowPremiumHostResponseTlsEnum struct {
	TLS_V1_0 ShowPremiumHostResponseTls
	TLS_V1_1 ShowPremiumHostResponseTls
	TLS_V1_2 ShowPremiumHostResponseTls
	TLS_V1_3 ShowPremiumHostResponseTls
}

func GetShowPremiumHostResponseTlsEnum() ShowPremiumHostResponseTlsEnum {
	return ShowPremiumHostResponseTlsEnum{
		TLS_V1_0: ShowPremiumHostResponseTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: ShowPremiumHostResponseTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: ShowPremiumHostResponseTls{
			value: "TLS v1.2",
		},
		TLS_V1_3: ShowPremiumHostResponseTls{
			value: "TLS v1.3",
		},
	}
}

func (c ShowPremiumHostResponseTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPremiumHostResponseTls) UnmarshalJSON(b []byte) error {
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

type ShowPremiumHostResponseCipher struct {
	value string
}

type ShowPremiumHostResponseCipherEnum struct {
	CIPHER_1       ShowPremiumHostResponseCipher
	CIPHER_2       ShowPremiumHostResponseCipher
	CIPHER_3       ShowPremiumHostResponseCipher
	CIPHER_4       ShowPremiumHostResponseCipher
	CIPHER_DEFAULT ShowPremiumHostResponseCipher
}

func GetShowPremiumHostResponseCipherEnum() ShowPremiumHostResponseCipherEnum {
	return ShowPremiumHostResponseCipherEnum{
		CIPHER_1: ShowPremiumHostResponseCipher{
			value: "cipher_1",
		},
		CIPHER_2: ShowPremiumHostResponseCipher{
			value: "cipher_2",
		},
		CIPHER_3: ShowPremiumHostResponseCipher{
			value: "cipher_3",
		},
		CIPHER_4: ShowPremiumHostResponseCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: ShowPremiumHostResponseCipher{
			value: "cipher_default",
		},
	}
}

func (c ShowPremiumHostResponseCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPremiumHostResponseCipher) UnmarshalJSON(b []byte) error {
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
