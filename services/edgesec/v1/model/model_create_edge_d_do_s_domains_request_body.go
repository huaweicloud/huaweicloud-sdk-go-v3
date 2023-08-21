package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEdgeDDoSDomainsRequestBody struct {

	// waf防护域名ID
	DomainId string `json:"domain_id"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateEdgeDDoSDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeDDoSDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEdgeDDoSDomainsRequestBody", string(data)}, " ")
}
