package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRefreshTasksRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *RefreshTaskRequest `json:"body,omitempty"`
}

func (o CreateRefreshTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRefreshTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateRefreshTasksRequest", string(data)}, " ")
}
