package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceWebhookRequest Request Object
type DeleteInstanceWebhookRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 触发器策略ID
	PolicyId int32 `json:"policy_id"`
}

func (o DeleteInstanceWebhookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceWebhookRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceWebhookRequest", string(data)}, " ")
}
