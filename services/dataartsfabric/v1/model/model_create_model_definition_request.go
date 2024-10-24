package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelDefinitionRequest Request Object
type CreateModelDefinitionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateModelInput `json:"body,omitempty"`
}

func (o CreateModelDefinitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelDefinitionRequest struct{}"
	}

	return strings.Join([]string{"CreateModelDefinitionRequest", string(data)}, " ")
}
