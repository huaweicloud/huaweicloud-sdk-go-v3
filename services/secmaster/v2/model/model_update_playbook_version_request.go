package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookVersionRequest Request Object
type UpdatePlaybookVersionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	Body *ModifyPlaybookVersionInfo `json:"body,omitempty"`
}

func (o UpdatePlaybookVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookVersionRequest", string(data)}, " ")
}
