package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookRequest Request Object
type UpdatePlaybookRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// ID of playbook
	PlaybookId string `json:"playbook_id"`

	Body *ModifyPlaybookInfo `json:"body,omitempty"`
}

func (o UpdatePlaybookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookRequest struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookRequest", string(data)}, " ")
}
