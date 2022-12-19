package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePlaybookActionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	// ID of action
	ActionId string `json:"action_id"`
}

func (o DeletePlaybookActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookActionRequest struct{}"
	}

	return strings.Join([]string{"DeletePlaybookActionRequest", string(data)}, " ")
}
