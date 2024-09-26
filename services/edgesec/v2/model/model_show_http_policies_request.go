package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpPoliciesRequest Request Object
type ShowHttpPoliciesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页显示pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 模糊查询策略名称
	Name *string `json:"name,omitempty"`

	// 根据域名模糊查询策略
	Hostname *string `json:"hostname,omitempty"`
}

func (o ShowHttpPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpPoliciesRequest", string(data)}, " ")
}
