package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpReferenceTablesRequest Request Object
type ShowHttpReferenceTablesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 模糊查询引用表名称
	Name *string `json:"name,omitempty"`
}

func (o ShowHttpReferenceTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpReferenceTablesRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpReferenceTablesRequest", string(data)}, " ")
}
