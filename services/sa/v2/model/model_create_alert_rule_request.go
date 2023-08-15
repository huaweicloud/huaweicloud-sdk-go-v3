package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRuleRequest Request Object
type CreateAlertRuleRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAlertRuleRequestBody `json:"body,omitempty"`
}

func (o CreateAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleRequest", string(data)}, " ")
}
