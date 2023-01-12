package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePlaybookRuleRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	// rule_id
	RuleId string `json:"rule_id"`

	Body *ModifyRuleInfo `json:"body,omitempty"`
}

func (o UpdatePlaybookRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookRuleRequest", string(data)}, " ")
}
