package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHostRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 分页查询时，返回第几页数据。范围0-100000，默认值为1，表示返回第1页数据。

	Page *int32 `json:"page,omitempty"`
	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 域名

	Hostname *string `json:"hostname,omitempty"`
	// 策略名

	Policyname *string `json:"policyname,omitempty"`
}

func (o ListHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRequest struct{}"
	}

	return strings.Join([]string{"ListHostRequest", string(data)}, " ")
}
