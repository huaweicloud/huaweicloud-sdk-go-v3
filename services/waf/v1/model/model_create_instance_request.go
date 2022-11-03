package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {

	// 通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
