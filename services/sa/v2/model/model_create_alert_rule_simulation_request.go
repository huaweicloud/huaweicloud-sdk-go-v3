package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRuleSimulationRequest Request Object
type CreateAlertRuleSimulationRequest struct {

	// workspace_id
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
