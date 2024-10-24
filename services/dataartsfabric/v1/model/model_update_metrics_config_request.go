package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetricsConfigRequest Request Object
type UpdateMetricsConfigRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateMetricsConfigInput `json:"body,omitempty"`
}

func (o UpdateMetricsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetricsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateMetricsConfigRequest", string(data)}, " ")
}
