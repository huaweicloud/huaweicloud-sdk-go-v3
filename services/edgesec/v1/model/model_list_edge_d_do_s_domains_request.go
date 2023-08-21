package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeDDoSDomainsRequest Request Object
type ListEdgeDDoSDomainsRequest struct {

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEdgeDDoSDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeDDoSDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeDDoSDomainsRequest", string(data)}, " ")
}
