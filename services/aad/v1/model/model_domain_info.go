package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainInfo struct {

	// 域名ID
	DomainId *string `json:"domain_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 域名cname
	Cname *string `json:"cname,omitempty"`

	// 域名协议
	Protocol *[]string `json:"protocol,omitempty"`

	// 源站类型
	RealServerType *int32 `json:"real_server_type,omitempty"`

	// 源站
	RealServers *string `json:"real_servers,omitempty"`

	// waf防护状态
	WafStatus *int32 `json:"waf_status,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainInfo struct{}"
	}

	return strings.Join([]string{"DomainInfo", string(data)}, " ")
}
