package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceResponse Response Object
type UpdateWorkspaceResponse struct {

	// 修改云办公服务属性的任务ID
	JobId *string `json:"job_id,omitempty"`

	// 企业ID
	EnterpriseId   *string `json:"enterprise_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceResponse", string(data)}, " ")
}
