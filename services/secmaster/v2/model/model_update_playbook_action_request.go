package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookActionRequest Request Object
type UpdatePlaybookActionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	// 剧本动作ID
	ActionId string `json:"action_id"`

	Body *ModifyActionInfo `json:"body,omitempty"`
}

func (o UpdatePlaybookActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookActionRequest struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookActionRequest", string(data)}, " ")
}
