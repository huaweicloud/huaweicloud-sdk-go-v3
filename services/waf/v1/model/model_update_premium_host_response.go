package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdatePremiumHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty" xml:"id"`

	// 创建的独享模式防护域名
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 对外协议，客户端（例如浏览器）请求访问网站的协议类型
	Protocol *string `json:"protocol,omitempty" xml:"protocol"`

	// 防护域名的源站服务器配置信息
	Server *[]PremiumWafServer `json:"server,omitempty" xml:"server"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty" xml:"proxy"`

	// 预留参数，用于后期设计冻结域名，解锁域名功能，目前暂不支持
	Locked *int32 `json:"locked,omitempty" xml:"locked"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对于低于最低TLS版本的请求，将无法正常访问网站
	Tls *UpdatePremiumHostResponseTls `json:"tls,omitempty" xml:"tls"`

	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *UpdatePremiumHostResponseCipher `json:"cipher,omitempty" xml:"cipher"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty" xml:"extend"`

	Flag *Flag `json:"flag,omitempty" xml:"flag"`

	// 域名描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 防护域名初始绑定的策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 帐号ID,对应华为云控制台用户名->我的凭证->帐号ID
	Domainid *string `json:"domainid,omitempty" xml:"domainid"`

	// 项目ID，对应华为云控制台用户名->我的凭证->项目列表->项目ID
	Projectid *string `json:"projectid,omitempty" xml:"projectid"`

	// 企业项目ID，对应华为云控制台用户名->企业->项目管理->点击项目名称->ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// https证书id
	Certificateid *string `json:"certificateid,omitempty" xml:"certificateid"`

	// 证书名称
	Certificatename *string `json:"certificatename,omitempty" xml:"certificatename"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty" xml:"protect_status"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty" xml:"access_status"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty" xml:"web_tag"`

	// LB负载均衡，默认轮询，不支持修改
	LbAlgorithm *string `json:"lb_algorithm,omitempty" xml:"lb_algorithm"`

	BlockPage *BlockPage `json:"block_page,omitempty" xml:"block_page"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty" xml:"traffic_mark"`

	TimeoutConfig *TimeoutConfig `json:"timeout_config,omitempty" xml:"timeout_config"`

	// 接入进度，仅用于新版console(前端)使用
	AccessProgress *[]AccessProgress `json:"access_progress,omitempty" xml:"access_progress"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdatePremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostResponse", string(data)}, " ")
}

type UpdatePremiumHostResponseTls struct {
	value string
}

type UpdatePremiumHostResponseTlsEnum struct {
	TLS_V1_0 UpdatePremiumHostResponseTls
	TLS_V1_1 UpdatePremiumHostResponseTls
	TLS_V1_2 UpdatePremiumHostResponseTls
	TLS_V1_3 UpdatePremiumHostResponseTls
}

func GetUpdatePremiumHostResponseTlsEnum() UpdatePremiumHostResponseTlsEnum {
	return UpdatePremiumHostResponseTlsEnum{
		TLS_V1_0: UpdatePremiumHostResponseTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: UpdatePremiumHostResponseTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: UpdatePremiumHostResponseTls{
			value: "TLS v1.2",
		},
		TLS_V1_3: UpdatePremiumHostResponseTls{
			value: "TLS v1.3",
		},
	}
}

func (c UpdatePremiumHostResponseTls) Value() string {
	return c.value
}

func (c UpdatePremiumHostResponseTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumHostResponseTls) UnmarshalJSON(b []byte) error {
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

type UpdatePremiumHostResponseCipher struct {
	value string
}

type UpdatePremiumHostResponseCipherEnum struct {
	CIPHER_1       UpdatePremiumHostResponseCipher
	CIPHER_2       UpdatePremiumHostResponseCipher
	CIPHER_3       UpdatePremiumHostResponseCipher
	CIPHER_4       UpdatePremiumHostResponseCipher
	CIPHER_DEFAULT UpdatePremiumHostResponseCipher
}

func GetUpdatePremiumHostResponseCipherEnum() UpdatePremiumHostResponseCipherEnum {
	return UpdatePremiumHostResponseCipherEnum{
		CIPHER_1: UpdatePremiumHostResponseCipher{
			value: "cipher_1",
		},
		CIPHER_2: UpdatePremiumHostResponseCipher{
			value: "cipher_2",
		},
		CIPHER_3: UpdatePremiumHostResponseCipher{
			value: "cipher_3",
		},
		CIPHER_4: UpdatePremiumHostResponseCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: UpdatePremiumHostResponseCipher{
			value: "cipher_default",
		},
	}
}

func (c UpdatePremiumHostResponseCipher) Value() string {
	return c.value
}

func (c UpdatePremiumHostResponseCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumHostResponseCipher) UnmarshalJSON(b []byte) error {
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
