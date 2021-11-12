package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableEnterpriseProjectRequest struct {
	// 企业项目ID，不能为0。

	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *EnableAction `json:"body,omitempty"`
}

func (o EnableEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"EnableEnterpriseProjectRequest", string(data)}, " ")
}
