package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPolicyRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询时，返回第几页数据。范围0-100000，默认值为1，表示返回第1页数据。
	Page *int32 `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`
}

func (o ListPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyRequest", string(data)}, " ")
}
