package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPolicyRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 每页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 策略名称

	Name *string `json:"name,omitempty"`
}

func (o ListPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyRequest", string(data)}, " ")
}
