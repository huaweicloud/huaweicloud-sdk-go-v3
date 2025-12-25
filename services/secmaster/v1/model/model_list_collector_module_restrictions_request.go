package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorModuleRestrictionsRequest Request Object
type ListCollectorModuleRestrictionsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ModuleTemplateFieldDto `json:"body,omitempty"`
}

func (o ListCollectorModuleRestrictionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorModuleRestrictionsRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorModuleRestrictionsRequest", string(data)}, " ")
}
