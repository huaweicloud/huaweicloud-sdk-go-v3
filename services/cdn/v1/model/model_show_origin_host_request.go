package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowOriginHostRequest struct {
	// 加速域名ID。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowOriginHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOriginHostRequest struct{}"
	}

	return strings.Join([]string{"ShowOriginHostRequest", string(data)}, " ")
}
