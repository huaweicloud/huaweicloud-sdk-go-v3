package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostRequest Request Object
type ListHostRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *int32 `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。如果需要一次查全部域名，该参数值填-1。
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 要查询的防护域名，用于查询指定防护域名信息；可不传，查询用户云模式下所有防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 防护策略名，用于查询指定防护策略下的域名，可不传
	Policyname *string `json:"policyname,omitempty"`
}

func (o ListHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRequest struct{}"
	}

	return strings.Join([]string{"ListHostRequest", string(data)}, " ")
}
