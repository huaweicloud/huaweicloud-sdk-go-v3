package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVariableGroupRequest Request Object
type CreateVariableGroupRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *CreateVariableGroupReq `json:"body,omitempty"`
}

func (o CreateVariableGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateVariableGroupRequest", string(data)}, " ")
}
