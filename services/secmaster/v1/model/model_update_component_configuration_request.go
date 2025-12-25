package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComponentConfigurationRequest Request Object
type UpdateComponentConfigurationRequest struct {

	// 组件id
	ComponentId string `json:"component_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *ConfigurationInfoDto `json:"body,omitempty"`
}

func (o UpdateComponentConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateComponentConfigurationRequest", string(data)}, " ")
}
