package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPremiumHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 页码

	Page *string `json:"page,omitempty"`
	// 每页条数

	Pagesize *string `json:"pagesize,omitempty"`
	// 域名

	Hostname *string `json:"hostname,omitempty"`
	// 策略名称

	Policyname *string `json:"policyname,omitempty"`
	// 域名防护状态

	ProtectStatus *int32 `json:"protect_status,omitempty"`
}

func (o ListPremiumHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPremiumHostRequest struct{}"
	}

	return strings.Join([]string{"ListPremiumHostRequest", string(data)}, " ")
}
