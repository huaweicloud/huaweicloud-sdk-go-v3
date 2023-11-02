package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRuleRequest Request Object
type DeleteAlertRuleRequest struct {

	// 工作空间 ID。Workspace ID.
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
