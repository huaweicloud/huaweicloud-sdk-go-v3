package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageNotificationPolicyRequest Request Object
type ListMessageNotificationPolicyRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 消息类型。job:任务执行结果消息。
	MessageType *string `json:"message_type,omitempty"`

	// 名称样式。支持模糊匹配，区分大小写
	NamePattern *string `json:"name_pattern,omitempty"`

	// 消息类型。SUCCESS:成功通知；FAILED：失败通知
	NotifyType *string `json:"notify_type,omitempty"`
}

func (o ListMessageNotificationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageNotificationPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListMessageNotificationPolicyRequest", string(data)}, " ")
}
