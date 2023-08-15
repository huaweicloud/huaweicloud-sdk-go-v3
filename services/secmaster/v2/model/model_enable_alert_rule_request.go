package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAlertRuleRequest Request Object
type EnableAlertRuleRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o EnableAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"EnableAlertRuleRequest", string(data)}, " ")
}
