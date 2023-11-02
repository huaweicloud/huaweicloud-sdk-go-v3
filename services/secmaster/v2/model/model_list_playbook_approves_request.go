package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookApprovesRequest Request Object
type ListPlaybookApprovesRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 审核类型。（PLAYBOOK-剧本, AOP_WORKFLOW--流程)
	ApproveType *string `json:"approve_type,omitempty"`
}

func (o ListPlaybookApprovesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookApprovesRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookApprovesRequest", string(data)}, " ")
}
