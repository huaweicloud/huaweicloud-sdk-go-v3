package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookRuleRequest Request Object
type CreatePlaybookRuleRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	Body *CreateRuleInfo `json:"body,omitempty"`
}

func (o CreatePlaybookRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookRuleRequest struct{}"
	}

	return strings.Join([]string{"CreatePlaybookRuleRequest", string(data)}, " ")
}
