package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceWebhookRequest Request Object
type CreateInstanceWebhookRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	Body *CreateWebhookPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceWebhookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceWebhookRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceWebhookRequest", string(data)}, " ")
}
