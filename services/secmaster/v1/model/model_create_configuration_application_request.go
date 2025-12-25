package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationApplicationRequest Request Object
type CreateConfigurationApplicationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 组件ID
	ComponentId string `json:"component_id"`

	Body *CreateConfigurationApplicationRequestBody `json:"body,omitempty"`
}

func (o CreateConfigurationApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationApplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationApplicationRequest", string(data)}, " ")
}
