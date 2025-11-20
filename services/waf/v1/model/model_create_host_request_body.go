package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHostRequestBody 域名请求体
type CreateHostRequestBody struct {

	// 域名（域名只能由字母、数字、-、_和.组成，长度不能超过64个字符，如www.domain.com）
	Hostname string `json:"hostname"`

	// 防护域名初始绑定的策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// 防护域名的源站服务器配置信息
	Server []CloudWafServer `json:"server"`

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificatename *string `json:"certificatename,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// 是否使用独享ip   - true：使用独享ip   - false：不使用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 套餐付费模式，默认值为prePaid。prePaid：包周期模式；postPaid：按需模式。
	PaidType *string `json:"paid_type,omitempty"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy bool `json:"proxy"`

	// LB负载均衡，仅专业版和企业版支持配置负载均衡算法   - 源IP Hash：将某个IP的请求定向到同一个服务器   - 加权轮询：所有请求将按权重轮流分配给源站服务器   - Session Hash：将某个Session标识的请求定向到同一个源站服务器，请确保在域名添加完毕后配置攻击惩罚的流量标识，否则Session Hash配置不生效
	LbAlgorithm *CreateHostRequestBodyLbAlgorithm `json:"lb_algorithm,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`
}

func (o CreateHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostRequestBody", string(data)}, " ")
}

type CreateHostRequestBodyLbAlgorithm struct {
	value string
}

type CreateHostRequestBodyLbAlgorithmEnum struct {
	IP_HASH      CreateHostRequestBodyLbAlgorithm
	ROUND_ROBIN  CreateHostRequestBodyLbAlgorithm
	SESSION_HASH CreateHostRequestBodyLbAlgorithm
}

func GetCreateHostRequestBodyLbAlgorithmEnum() CreateHostRequestBodyLbAlgorithmEnum {
	return CreateHostRequestBodyLbAlgorithmEnum{
		IP_HASH: CreateHostRequestBodyLbAlgorithm{
			value: "ip_hash",
		},
		ROUND_ROBIN: CreateHostRequestBodyLbAlgorithm{
			value: "round_robin",
		},
		SESSION_HASH: CreateHostRequestBodyLbAlgorithm{
			value: "session_hash",
		},
	}
}

func (c CreateHostRequestBodyLbAlgorithm) Value() string {
	return c.value
}

func (c CreateHostRequestBodyLbAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostRequestBodyLbAlgorithm) UnmarshalJSON(b []byte) error {
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
