package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionTargetRequest Request Object
type UpdateSubscriptionTargetRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 事件订阅目标ID
	TargetId string `json:"target_id"`

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SubscriptionTarget `json:"body,omitempty"`
}

func (o UpdateSubscriptionTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionTargetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionTargetRequest", string(data)}, " ")
}
