package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDomainDetailRequest struct {
	// 加速域名ID。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 企业项目ID。该参数仅对开启了企业项目功能的用户生效，不传表示查询default项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailRequest", string(data)}, " ")
}
