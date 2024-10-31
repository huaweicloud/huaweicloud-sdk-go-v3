package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainsRequest Request Object
type ShowDomainsRequest struct {

	// 查询列表的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 查询列表每一页的条数
	Limit *int32 `json:"limit,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainsRequest", string(data)}, " ")
}
