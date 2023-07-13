package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRequest Request Object
type ListInstanceRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
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
