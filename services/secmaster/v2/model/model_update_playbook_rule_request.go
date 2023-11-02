package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookRuleRequest Request Object
type UpdatePlaybookRuleRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	// 剧本规则ID
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
