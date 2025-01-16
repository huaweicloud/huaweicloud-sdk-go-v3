package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotifyPolicyRequest Request Object
type UpdateNotifyPolicyRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 通知策略ID。
	NotifyPolicyId string `json:"notify_policy_id"`

	Body *NotifyPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateNotifyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotifyPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotifyPolicyRequest", string(data)}, " ")
}
