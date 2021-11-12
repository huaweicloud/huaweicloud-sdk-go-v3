package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowHttpInfoRequest struct {
	// 加速域名ID。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowHttpInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpInfoRequest", string(data)}, " ")
}
