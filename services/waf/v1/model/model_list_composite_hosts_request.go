package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCompositeHostsRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// page

	Page *int32 `json:"page,omitempty"`
	// 每页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 域名名称

	Hostname *string `json:"hostname,omitempty"`
	// 防护策略名称

	Policyname *string `json:"policyname,omitempty"`
	// 域名防护状态

	ProtectStatus *int32 `json:"protect_status,omitempty"`
	// 域名所属WAF模式

	WafType *string `json:"waf_type,omitempty"`
	// 域名是否使用HTTPS

	IsHttps *bool `json:"is_https,omitempty"`
}

func (o ListCompositeHostsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCompositeHostsRequest struct{}"
	}

	return strings.Join([]string{"ListCompositeHostsRequest", string(data)}, " ")
}
