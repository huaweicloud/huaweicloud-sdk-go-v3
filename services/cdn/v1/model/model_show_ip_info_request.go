package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowIpInfoRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// IP地址列表，以“，”分割，最多20个。

	Ips string `json:"ips"`
}

func (o ShowIpInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowIpInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowIpInfoRequest", string(data)}, " ")
}
