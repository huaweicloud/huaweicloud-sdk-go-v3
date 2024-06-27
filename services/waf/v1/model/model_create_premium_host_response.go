package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePremiumHostResponse Response Object
type CreatePremiumHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 对外协议，客户端（例如浏览器）请求访问网站的协议类型
	Protocol *CreatePremiumHostResponseProtocol `json:"protocol,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]PremiumWafServer `json:"server,omitempty"`

	// 是否使用代理   - true：代表使用代理   - false：代表未使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 域名冻结状态，0表示未冻结，1表示为冻结，冗余参数
	Locked *int32 `json:"locked,omitempty"`

	// 创建域名的时间，13位毫秒时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对于低于最低TLS版本的请求，将无法正常访问网站
	Tls *CreatePremiumHostResponseTls `json:"tls,omitempty"`

	// 对外协议为https时才有cipher参数，加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_5：加密算法为AES128-SHA:AES256-SHA:AES128-SHA256:AES256-SHA256:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!EXPORT:!DES:!MD5:!PSK:!RC4:!DHE:@STRENGTH    cipher_6：加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *CreatePremiumHostResponseCipher `json:"cipher,omitempty"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 防护域名初始绑定的防护策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// 账号ID,对应华为云控制台用户名->我的凭证->账号ID
	Domainid *string `json:"domainid,omitempty"`

	// 项目ID，对应华为云控制台用户名->我的凭证->项目列表->项目ID
	Projectid *string `json:"projectid,omitempty"`

	// 企业项目ID，对应华为云控制台用户名->企业->项目管理->点击项目名称->ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名防护状态：  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreatePremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostResponse", string(data)}, " ")
}

type CreatePremiumHostResponseProtocol struct {
	value string
}

type CreatePremiumHostResponseProtocolEnum struct {
	HTTPS     CreatePremiumHostResponseProtocol
	HTTP      CreatePremiumHostResponseProtocol
	HTTPHTTPS CreatePremiumHostResponseProtocol
}

func GetCreatePremiumHostResponseProtocolEnum() CreatePremiumHostResponseProtocolEnum {
	return CreatePremiumHostResponseProtocolEnum{
		HTTPS: CreatePremiumHostResponseProtocol{
			value: "HTTPS",
		},
		HTTP: CreatePremiumHostResponseProtocol{
			value: "HTTP",
		},
		HTTPHTTPS: CreatePremiumHostResponseProtocol{
			value: "HTTP&HTTPS",
		},
	}
}

func (c CreatePremiumHostResponseProtocol) Value() string {
	return c.value
}

func (c CreatePremiumHostResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePremiumHostResponseProtocol) UnmarshalJSON(b []byte) error {
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

type CreatePremiumHostResponseTls struct {
	value string
}

type CreatePremiumHostResponseTlsEnum struct {
	TLS_V1_0 CreatePremiumHostResponseTls
	TLS_V1_1 CreatePremiumHostResponseTls
	TLS_V1_2 CreatePremiumHostResponseTls
}

func GetCreatePremiumHostResponseTlsEnum() CreatePremiumHostResponseTlsEnum {
	return CreatePremiumHostResponseTlsEnum{
		TLS_V1_0: CreatePremiumHostResponseTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: CreatePremiumHostResponseTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: CreatePremiumHostResponseTls{
			value: "TLS v1.2",
		},
	}
}

func (c CreatePremiumHostResponseTls) Value() string {
	return c.value
}

func (c CreatePremiumHostResponseTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePremiumHostResponseTls) UnmarshalJSON(b []byte) error {
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

type CreatePremiumHostResponseCipher struct {
	value string
}

type CreatePremiumHostResponseCipherEnum struct {
	CIPHER_1       CreatePremiumHostResponseCipher
	CIPHER_2       CreatePremiumHostResponseCipher
	CIPHER_3       CreatePremiumHostResponseCipher
	CIPHER_4       CreatePremiumHostResponseCipher
	CIPHER_5       CreatePremiumHostResponseCipher
	CIPHER_6       CreatePremiumHostResponseCipher
	CIPHER_DEFAULT CreatePremiumHostResponseCipher
}

func GetCreatePremiumHostResponseCipherEnum() CreatePremiumHostResponseCipherEnum {
	return CreatePremiumHostResponseCipherEnum{
		CIPHER_1: CreatePremiumHostResponseCipher{
			value: "cipher_1",
		},
		CIPHER_2: CreatePremiumHostResponseCipher{
			value: "cipher_2",
		},
		CIPHER_3: CreatePremiumHostResponseCipher{
			value: "cipher_3",
		},
		CIPHER_4: CreatePremiumHostResponseCipher{
			value: "cipher_4",
		},
		CIPHER_5: CreatePremiumHostResponseCipher{
			value: "cipher_5",
		},
		CIPHER_6: CreatePremiumHostResponseCipher{
			value: "cipher_6",
		},
		CIPHER_DEFAULT: CreatePremiumHostResponseCipher{
			value: "cipher_default",
		},
	}
}

func (c CreatePremiumHostResponseCipher) Value() string {
	return c.value
}

func (c CreatePremiumHostResponseCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePremiumHostResponseCipher) UnmarshalJSON(b []byte) error {
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
