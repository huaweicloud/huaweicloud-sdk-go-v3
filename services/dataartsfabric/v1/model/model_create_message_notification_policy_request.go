package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMessageNotificationPolicyRequest Request Object
type CreateMessageNotificationPolicyRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateMessageNotificationPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateMessageNotificationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageNotificationPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateMessageNotificationPolicyRequest", string(data)}, " ")
}
