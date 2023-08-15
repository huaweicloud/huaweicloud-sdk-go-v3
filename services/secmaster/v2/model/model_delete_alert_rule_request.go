package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRuleRequest Request Object
type DeleteAlertRuleRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o DeleteAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlertRuleRequest", string(data)}, " ")
}
