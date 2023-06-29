package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupRequest Request Object
type CreateIpGroupRequest struct {

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateIpGroupRequestBody `json:"body,omitempty"`
}

func (o CreateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequest", string(data)}, " ")
}
