package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePremiumHostRequestBody 创建独享模式域名的请求体
type CreatePremiumHostRequestBody struct {

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificatename *string `json:"certificatename,omitempty"`

	// 防护域名或IP（可带端口）
	Hostname string `json:"hostname"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy bool `json:"proxy"`

	// 防护域名初始绑定的防护策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// 防护域名的源站服务器配置信息
	Server []PremiumWafServer `json:"server"`

	BlockPage *BlockPage `json:"block_page,omitempty"`

	// 字段转发配置，WAF会将添加的字段插到header中，转给源站；Key不能跟nginx原生字段重复。Value支持的值包括:   - $time_local   - $request_id   - $connection_requests   - $tenant_id   - $project_id   - $remote_addr   - $remote_port   - $scheme   - $request_method   - $http_host   -$origin_uri   - $request_length   - $ssl_server_name   - $ssl_protocol   - $ssl_curves   - $ssl_session_reused
	ForwardHeaderMap map[string]string `json:"forward_header_map,omitempty"`

	// 防护域名备注
	Description *string `json:"description,omitempty"`
}

func (o CreatePremiumHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostRequestBody", string(data)}, " ")
}
