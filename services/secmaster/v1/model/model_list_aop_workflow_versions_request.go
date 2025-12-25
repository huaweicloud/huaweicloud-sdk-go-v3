package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAopWorkflowVersionsRequest Request Object
type ListAopWorkflowVersionsRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 流程ID **取值范围**: 不涉及
	WorkflowId string `json:"workflow_id"`

	// **参数解释**: 流程的状态 - pending_submit 草稿 - pending_approval 待审核 - publishing 发布中 - publish_failed 发布失败 - not_activated 未激活 - activated 已激活 - rejected 审核拒绝  **取值范围**: - pending_submit - pending_approval - publishing - publish_failed - not_activated - activated - rejected
	Status *string `json:"status,omitempty"`
}

func (o ListAopWorkflowVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAopWorkflowVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListAopWorkflowVersionsRequest", string(data)}, " ")
}
