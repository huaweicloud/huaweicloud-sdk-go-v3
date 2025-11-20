package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainNameConfigResponse Response Object
type ShowDomainNameConfigResponse struct {

	// 域名id
	DomainId *string `json:"domain_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 源站类型，0 - 源站IP， 1 - 源站域名
	RealServerType *int32 `json:"real_server_type,omitempty"`

	// HTTP端口，与port_https不能同时为空。DDoS高防支持的HTTP端口可在控制台查看。
	PortHttp *[]int32 `json:"port_http,omitempty"`

	// HTTPS端口，与port_http不能同时为空。DDoS高防支持的HTTPS端口可在控制台查看。
	PortHttps *[]int32 `json:"port_https,omitempty"`

	// 源站ip/源站域名
	RealServer *string `json:"real_server,omitempty"`

	// pp是否开启, 1-开启，0-关闭
	PpEnable *int32 `json:"pp_enable,omitempty"`

	// 防护区域,0-大陆,1-海外
	OverseasType *int32 `json:"overseas_type,omitempty"`

	// tls(请求参数type=waf时)
	Tls *string `json:"tls,omitempty"`

	// 加密套件(请求参数type=waf时)
	Cipher *string `json:"cipher,omitempty"`

	// 是否允许http2(请求参数type=waf时)
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 字段转发(请求参数type=waf时)
	HeaderMap      *interface{} `json:"header_map,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDomainNameConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainNameConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainNameConfigResponse", string(data)}, " ")
}
