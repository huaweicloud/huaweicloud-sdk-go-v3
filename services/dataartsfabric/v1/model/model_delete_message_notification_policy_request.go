package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMessageNotificationPolicyRequest Request Object
type DeleteMessageNotificationPolicyRequest struct {

	// 消息通知策略ID
	MessagePolicyId string `json:"message_policy_id"`

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`
}

func (o DeleteMessageNotificationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageNotificationPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteMessageNotificationPolicyRequest", string(data)}, " ")
}
