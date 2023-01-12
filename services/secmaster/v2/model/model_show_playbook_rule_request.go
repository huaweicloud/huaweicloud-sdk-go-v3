package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPlaybookRuleRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	// version Id value
	RuleId string `json:"rule_id"`
}

func (o ShowPlaybookRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookRuleRequest", string(data)}, " ")
}
