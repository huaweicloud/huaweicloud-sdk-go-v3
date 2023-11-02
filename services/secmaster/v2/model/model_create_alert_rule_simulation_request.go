package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRuleSimulationRequest Request Object
type CreateAlertRuleSimulationRequest struct {

	// 工作空间 ID。Workspace ID.
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAlertRuleSimulationRequestBody `json:"body,omitempty"`
}

func (o CreateAlertRuleSimulationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleSimulationRequest struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleSimulationRequest", string(data)}, " ")
}
