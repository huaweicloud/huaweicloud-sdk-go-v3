package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisableEnterpriseProjectRequest struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *DisableAction `json:"body,omitempty"`
}

func (o DisableEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"DisableEnterpriseProjectRequest", string(data)}, " ")
}
