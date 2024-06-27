package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHostResponse Response Object
type ShowHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 防护域名的防护策略id
	Policyid *string `json:"policyid,omitempty"`

	// 账号ID,对应华为云控制台用户名->我的凭证->账号ID
	Domainid *string `json:"domainid,omitempty"`

	// 项目ID，对应华为云控制台用户名->我的凭证->项目列表->项目ID
	Projectid *string `json:"projectid,omitempty"`

	// 企业项目ID，对应华为云控制台用户名->企业->项目管理->点击项目名称->ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 后端包含的协议类型：HTTPS、HTTP、HTTP&HTTPS
	Protocol *string `json:"protocol,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]CloudWafServer `json:"server,omitempty"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty"`

	// 预留参数，用于后期设计冻结域名，解锁域名功能，目前暂不支持
	Locked *int32 `json:"locked,omitempty"`

	// 创建防护域名的时间戳（毫秒）
	Timestamp *int64 `json:"timestamp,omitempty"`

	// https证书id
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名称
	Certificatename *string `json:"certificatename,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对于低于最低TLS版本的请求，将无法正常访问网站
	Tls *ShowHostResponseTls `json:"tls,omitempty"`

	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_5：加密算法为AES128-SHA:AES256-SHA:AES128-SHA256:AES256-SHA256:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!EXPORT:!DES:!MD5:!PSK:!RC4:!DHE:@STRENGTH    cipher_6：加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *ShowHostResponseCipher `json:"cipher,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	// 扩展字段，用于存放Web基础防护中一些开关配置等信息
	Extend map[string]string `json:"extend,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`

	CircuitBreaker *CircuitBreaker `json:"circuit_breaker,omitempty"`

	// LB负载均衡，仅专业版（原企业版）和铂金版（原旗舰版）支持配置负载均衡算法   - 源IP Hash：将某个IP的请求定向到同一个服务器   - 加权轮询：所有请求将按权重轮流分配给源站服务器   - Session Hash：将某个Session标识的请求定向到同一个源站服务器，请确保在域名添加完毕后配置攻击惩罚的流量标识，否则Session Hash配置不生效
	LbAlgorithm *ShowHostResponseLbAlgorithm `json:"lb_algorithm,omitempty"`

	TimeoutConfig *TimeoutConfig `json:"timeout_config,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 网站备注
	Description *string `json:"description,omitempty"`

	// 是否支持http2   - true：表示支持http2   - false：表示不支持http2
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 是否使用独享ip   - true：使用独享ip   - false：不实用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 接入进度，仅用于新版console(前端)使用
	AccessProgress *[]AccessProgress `json:"access_progress,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostResponse struct{}"
	}

	return strings.Join([]string{"ShowHostResponse", string(data)}, " ")
}

type ShowHostResponseTls struct {
	value string
}

type ShowHostResponseTlsEnum struct {
	TLS_V1_0 ShowHostResponseTls
	TLS_V1_1 ShowHostResponseTls
	TLS_V1_2 ShowHostResponseTls
}

func GetShowHostResponseTlsEnum() ShowHostResponseTlsEnum {
	return ShowHostResponseTlsEnum{
		TLS_V1_0: ShowHostResponseTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: ShowHostResponseTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: ShowHostResponseTls{
			value: "TLS v1.2",
		},
	}
}

func (c ShowHostResponseTls) Value() string {
	return c.value
}

func (c ShowHostResponseTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHostResponseTls) UnmarshalJSON(b []byte) error {
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

type ShowHostResponseCipher struct {
	value string
}

type ShowHostResponseCipherEnum struct {
	CIPHER_1       ShowHostResponseCipher
	CIPHER_2       ShowHostResponseCipher
	CIPHER_3       ShowHostResponseCipher
	CIPHER_4       ShowHostResponseCipher
	CIPHER_5       ShowHostResponseCipher
	CIPHER_6       ShowHostResponseCipher
	CIPHER_DEFAULT ShowHostResponseCipher
}

func GetShowHostResponseCipherEnum() ShowHostResponseCipherEnum {
	return ShowHostResponseCipherEnum{
		CIPHER_1: ShowHostResponseCipher{
			value: "cipher_1",
		},
		CIPHER_2: ShowHostResponseCipher{
			value: "cipher_2",
		},
		CIPHER_3: ShowHostResponseCipher{
			value: "cipher_3",
		},
		CIPHER_4: ShowHostResponseCipher{
			value: "cipher_4",
		},
		CIPHER_5: ShowHostResponseCipher{
			value: "cipher_5",
		},
		CIPHER_6: ShowHostResponseCipher{
			value: "cipher_6",
		},
		CIPHER_DEFAULT: ShowHostResponseCipher{
			value: "cipher_default",
		},
	}
}

func (c ShowHostResponseCipher) Value() string {
	return c.value
}

func (c ShowHostResponseCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHostResponseCipher) UnmarshalJSON(b []byte) error {
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

type ShowHostResponseLbAlgorithm struct {
	value string
}

type ShowHostResponseLbAlgorithmEnum struct {
	IP_HASH      ShowHostResponseLbAlgorithm
	ROUND_ROBIN  ShowHostResponseLbAlgorithm
	SESSION_HASH ShowHostResponseLbAlgorithm
}

func GetShowHostResponseLbAlgorithmEnum() ShowHostResponseLbAlgorithmEnum {
	return ShowHostResponseLbAlgorithmEnum{
		IP_HASH: ShowHostResponseLbAlgorithm{
			value: "ip_hash",
		},
		ROUND_ROBIN: ShowHostResponseLbAlgorithm{
			value: "round_robin",
		},
		SESSION_HASH: ShowHostResponseLbAlgorithm{
			value: "session_hash",
		},
	}
}

func (c ShowHostResponseLbAlgorithm) Value() string {
	return c.value
}

func (c ShowHostResponseLbAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHostResponseLbAlgorithm) UnmarshalJSON(b []byte) error {
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
