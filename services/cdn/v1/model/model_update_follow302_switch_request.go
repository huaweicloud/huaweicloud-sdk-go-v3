package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateFollow302SwitchRequest struct {
	// 加速域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示资源所属项目，不传表示查询默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *Follow302StatusRequest `json:"body,omitempty"`
}

func (o UpdateFollow302SwitchRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFollow302SwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateFollow302SwitchRequest", string(data)}, " ")
}
