package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHostResponse Response Object
type CreateHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// LB负载均衡，仅专业版（原企业版）和铂金版（原旗舰版）支持配置负载均衡算法   - 源IP Hash：将某个IP的请求定向到同一个服务器   - 加权轮询：所有请求将按权重轮流分配给源站服务器   - Session Hash：将某个Session标识的请求定向到同一个源站服务器，请确保在域名添加完毕后配置攻击惩罚的流量标识，否则Session Hash配置不生效
	LbAlgorithm *CreateHostResponseLbAlgorithm `json:"lb_algorithm,omitempty"`

	// 返回的客户端协议类型
	Protocol *string `json:"protocol,omitempty"`

	// 返回的证书id
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名称
	Certificatename *string `json:"certificatename,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]CloudWafServer `json:"server,omitempty"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 是否使用独享ip   - true：使用独享ip   - false：不实用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// 是否支持http2   - true：表示支持http2   - false：表示不支持http2
	Http2Enable *bool `json:"http2_enable,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostResponse struct{}"
	}

	return strings.Join([]string{"CreateHostResponse", string(data)}, " ")
}

type CreateHostResponseLbAlgorithm struct {
	value string
}

type CreateHostResponseLbAlgorithmEnum struct {
	IP_HASH      CreateHostResponseLbAlgorithm
	ROUND_ROBIN  CreateHostResponseLbAlgorithm
	SESSION_HASH CreateHostResponseLbAlgorithm
}

func GetCreateHostResponseLbAlgorithmEnum() CreateHostResponseLbAlgorithmEnum {
	return CreateHostResponseLbAlgorithmEnum{
		IP_HASH: CreateHostResponseLbAlgorithm{
			value: "ip_hash",
		},
		ROUND_ROBIN: CreateHostResponseLbAlgorithm{
			value: "round_robin",
		},
		SESSION_HASH: CreateHostResponseLbAlgorithm{
			value: "session_hash",
		},
	}
}

func (c CreateHostResponseLbAlgorithm) Value() string {
	return c.value
}

func (c CreateHostResponseLbAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostResponseLbAlgorithm) UnmarshalJSON(b []byte) error {
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
