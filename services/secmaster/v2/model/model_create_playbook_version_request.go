package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookVersionRequest Request Object
type CreatePlaybookVersionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本ID
	PlaybookId string `json:"playbook_id"`

	Body *CreatePlaybookVersionInfo `json:"body,omitempty"`
}

func (o CreatePlaybookVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookVersionRequest struct{}"
	}

	return strings.Join([]string{"CreatePlaybookVersionRequest", string(data)}, " ")
}
