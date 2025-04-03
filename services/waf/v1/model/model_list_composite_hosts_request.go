package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCompositeHostsRequest Request Object
type ListCompositeHostsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。默认值为0，表示默认企业项目。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。  缺省值：0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *int32 `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。如果需要一次查全部域名，该参数值填-1。
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 域名名称
	Hostname *string `json:"hostname,omitempty"`

	// 防护策略名称
	Policyname *string `json:"policyname,omitempty"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名所属WAF模式
	WafType *string `json:"waf_type,omitempty"`

	// 域名是否使用HTTPS
	IsHttps *bool `json:"is_https,omitempty"`
}

func (o ListCompositeHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompositeHostsRequest struct{}"
	}

	return strings.Join([]string{"ListCompositeHostsRequest", string(data)}, " ")
}
