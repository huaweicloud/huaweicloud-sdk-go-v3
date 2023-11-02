package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookRequest Request Object
type DeletePlaybookRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// ID of playbook
	PlaybookId string `json:"playbook_id"`
}

func (o DeletePlaybookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookRequest struct{}"
	}

	return strings.Join([]string{"DeletePlaybookRequest", string(data)}, " ")
}
