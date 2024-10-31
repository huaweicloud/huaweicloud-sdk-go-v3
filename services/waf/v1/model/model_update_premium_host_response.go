package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePremiumHostResponse Response Object
type UpdatePremiumHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 创建的独享模式防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 对外协议，客户端（例如浏览器）请求访问网站的协议类型
	Protocol *string `json:"protocol,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]PremiumWafServer `json:"server,omitempty"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 预留参数，用于后期设计冻结域名，解锁域名功能，目前暂不支持
	Locked *int32 `json:"locked,omitempty"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对于低于最低TLS版本的请求，将无法正常访问网站
	Tls *UpdatePremiumHostResponseTls `json:"tls,omitempty"`

	// 加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH    cipher_5：加密算法为AES128-SHA:AES256-SHA:AES128-SHA256:AES256-SHA256:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!EXPORT:!DES:!MD5:!PSK:!RC4:!DHE:@STRENGTH    cipher_6：加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *UpdatePremiumHostResponseCipher `json:"cipher,omitempty"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 云模式elb接入域名返回此字段：elb-shared
	Mode *string `json:"mode,omitempty"`

	// 云模式elb接入域名返回此字段，表示负载均衡器（ELB）id
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 云模式elb接入域名返回此字段，表示监听器id
	ListenerId *string `json:"listener_id,omitempty"`

	// 云模式elb接入域名返回此字段， 表示业务端口
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 防护域名初始绑定的策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// 账号ID,对应华为云控制台用户名->我的凭证->账号ID
	Domainid *string `json:"domainid,omitempty"`

	// 项目ID，对应华为云控制台用户名->我的凭证->项目列表->项目ID
	Projectid *string `json:"projectid,omitempty"`

	// 企业项目ID，对应华为云控制台用户名->企业->项目管理->点击项目名称->ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// https证书id
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名称
	Certificatename *string `json:"certificatename,omitempty"`

	// 域名防护状态：  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// LB负载均衡，默认轮询，不支持修改
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`

	TimeoutConfig *TimeoutConfig `json:"timeout_config,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`

	// 接入进度，仅用于新版console(前端)使用
	AccessProgress *[]AccessProgress `json:"access_progress,omitempty"`
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

type UpdatePremiumHostResponseCipher struct {
	value string
}

type UpdatePremiumHostResponseCipherEnum struct {
	CIPHER_1       UpdatePremiumHostResponseCipher
	CIPHER_2       UpdatePremiumHostResponseCipher
	CIPHER_3       UpdatePremiumHostResponseCipher
	CIPHER_4       UpdatePremiumHostResponseCipher
	CIPHER_5       UpdatePremiumHostResponseCipher
	CIPHER_6       UpdatePremiumHostResponseCipher
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
		CIPHER_5: UpdatePremiumHostResponseCipher{
			value: "cipher_5",
		},
		CIPHER_6: UpdatePremiumHostResponseCipher{
			value: "cipher_6",
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
