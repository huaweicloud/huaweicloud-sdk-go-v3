package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelDefinitionRequest Request Object
type UpdateModelDefinitionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// Service ID
	ModelId string `json:"model_id"`

	Body *UpdateModelInput `json:"body,omitempty"`
}

func (o UpdateModelDefinitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelDefinitionRequest struct{}"
	}

	return strings.Join([]string{"UpdateModelDefinitionRequest", string(data)}, " ")
}
