package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHttpsInfoRequest struct {
	// 加速域名ID。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。注意：当使用子账号调用接口时，该参数必传。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *HttpInfoRequest `json:"body,omitempty"`
}

func (o UpdateHttpsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpsInfoRequest", string(data)}, " ")
}
