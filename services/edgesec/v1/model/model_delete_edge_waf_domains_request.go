package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeWafDomainsRequest Request Object
type DeleteEdgeWafDomainsRequest struct {

	// 防护域名id
	Domainid string `json:"domainid"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeleteEdgeWafDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeWafDomainsRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeWafDomainsRequest", string(data)}, " ")
}
