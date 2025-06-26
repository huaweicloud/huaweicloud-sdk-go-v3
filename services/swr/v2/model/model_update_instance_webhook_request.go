package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceWebhookRequest Request Object
type UpdateInstanceWebhookRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 触发器ID
	PolicyId int32 `json:"policy_id"`

	Body *UpdateWebhookPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceWebhookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceWebhookRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceWebhookRequest", string(data)}, " ")
}
