package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePlaybookActionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	// Create actions
	Body *[]interface{} `json:"body,omitempty"`
}

func (o CreatePlaybookActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookActionRequest struct{}"
	}

	return strings.Join([]string{"CreatePlaybookActionRequest", string(data)}, " ")
}
