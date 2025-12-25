package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookRuleRequest Request Object
type DeletePlaybookRuleRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	// 规则ID
	RuleId string `json:"rule_id"`
}

func (o DeletePlaybookRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookRuleRequest struct{}"
	}

	return strings.Join([]string{"DeletePlaybookRuleRequest", string(data)}, " ")
}
