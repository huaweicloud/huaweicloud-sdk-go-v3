package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstanceRequest struct {

	// 通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询参数，第page页，默认值为1
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页pagesize条记录，默认值为10
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 模糊查询，独享引擎名称
	Instancename *string `json:"instancename,omitempty"`
}

func (o ListInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceRequest", string(data)}, " ")
}
