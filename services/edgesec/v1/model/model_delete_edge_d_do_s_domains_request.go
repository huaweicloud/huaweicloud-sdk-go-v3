package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeDDoSDomainsRequest Request Object
type DeleteEdgeDDoSDomainsRequest struct {

	// 域名ID
	Domainid string `json:"domainid"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeleteEdgeDDoSDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeDDoSDomainsRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeDDoSDomainsRequest", string(data)}, " ")
}
