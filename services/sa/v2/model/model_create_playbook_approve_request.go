package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookApproveRequest Request Object
type CreatePlaybookApproveRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	Body *ApprovePlaybookInfo `json:"body,omitempty"`
}

func (o CreatePlaybookApproveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookApproveRequest struct{}"
	}

	return strings.Join([]string{"CreatePlaybookApproveRequest", string(data)}, " ")
}
