package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainsRequest Request Object
type ListCdnDomainsRequest struct {

	// 分页查询参数，偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询参数，每页显示limit条记录
	Limit *int32 `json:"limit,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListCdnDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListCdnDomainsRequest", string(data)}, " ")
}
