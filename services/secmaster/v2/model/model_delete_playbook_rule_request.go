package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookRuleRequest Request Object
type DeletePlaybookRuleRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	// rule_id
	RuleId string `json:"rule_id"`
}

func (o DeletePlaybookRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookRuleRequest struct{}"
	}

	return strings.Join([]string{"DeletePlaybookRuleRequest", string(data)}, " ")
}
