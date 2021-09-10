package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBlackWhiteListRequest struct {
	// 需要查询IP黑白名单的域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowBlackWhiteListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"ShowBlackWhiteListRequest", string(data)}, " ")
}
