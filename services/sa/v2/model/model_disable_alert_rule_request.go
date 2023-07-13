package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableAlertRuleRequest Request Object
type DisableAlertRuleRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o DisableAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"DisableAlertRuleRequest", string(data)}, " ")
}
