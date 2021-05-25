package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEnterpriseProjectRequest struct {
	// 企业项目ID。 可以通过查询企业项目列表接口获取。

	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o ShowEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectRequest", string(data)}, " ")
}
