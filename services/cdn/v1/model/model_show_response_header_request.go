package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResponseHeaderRequest struct {
	// 加速域名ID。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。注意：当使用子账号调用接口时，该参数必传。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowResponseHeaderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResponseHeaderRequest struct{}"
	}

	return strings.Join([]string{"ShowResponseHeaderRequest", string(data)}, " ")
}
