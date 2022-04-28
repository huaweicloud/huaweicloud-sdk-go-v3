package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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

	// 接入状态
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 返回的客户端协议类型
	Protocol *string `json:"protocol,omitempty"`

	// 返回的证书id
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名称
	Certificatename *string `json:"certificatename,omitempty"`

	// 源站信息
	Server *[]CloudWafServer `json:"server,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 是否开启了代理
	Proxy *bool `json:"proxy,omitempty"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 是否使用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 是否支持http2
	Http2Enable    *bool `json:"http2_enable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostResponse struct{}"
	}

	return strings.Join([]string{"CreateHostResponse", string(data)}, " ")
}
