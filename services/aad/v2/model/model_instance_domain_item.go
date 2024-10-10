package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceDomainItem struct {

	// 域名ID
	DomainId *string `json:"domain_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 域名cname
	Cname *string `json:"cname,omitempty"`

	// 域名状态 NORMAL = '0', FREEZE = '1'
	DomainStatus *string `json:"domain_status,omitempty"`

	// cc防护状态
	CcStatus *int32 `json:"cc_status,omitempty"`

	// 证书状态：1---已上传  2---未上传
	HttpsCertStatus *int32 `json:"https_cert_status,omitempty"`

	// 证书名称
	CertName *string `json:"cert_name,omitempty"`

	// 域名协议
	ProtocolType *[]string `json:"protocol_type,omitempty"`

	// 源站类型
	RealServerType *int32 `json:"real_server_type,omitempty"`

	// 源站
	RealServers *string `json:"real_servers,omitempty"`

	// waf防护状态
	WafStatus *int32 `json:"waf_status,omitempty"`
}

func (o InstanceDomainItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDomainItem struct{}"
	}

	return strings.Join([]string{"InstanceDomainItem", string(data)}, " ")
}
