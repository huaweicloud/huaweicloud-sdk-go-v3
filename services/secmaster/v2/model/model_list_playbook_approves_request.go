package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPlaybookApprovesRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// Resource Id, PlaybookId or AopworkflowId
	ResourceId *string `json:"resource_id,omitempty"`

	// PLAYBOOK, AOP_WORKFLOW
	ApproveType *string `json:"approve_type,omitempty"`
}

func (o ListPlaybookApprovesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookApprovesRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookApprovesRequest", string(data)}, " ")
}
