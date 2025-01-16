package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotifyPolicyRequest Request Object
type DeleteNotifyPolicyRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 通知策略ID。
	NotifyPolicyId string `json:"notify_policy_id"`
}

func (o DeleteNotifyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotifyPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotifyPolicyRequest", string(data)}, " ")
}
