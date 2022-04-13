package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEnterpriseProjectRequest struct {
	// 企业项目ID，不能为0。 可以通过查询企业项目列表接口获取。

	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *EnterpriseProject `json:"body,omitempty"`
}

func (o UpdateEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectRequest", string(data)}, " ")
}
