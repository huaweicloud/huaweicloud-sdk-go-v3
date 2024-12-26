package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionTargetRequest Request Object
type CreateSubscriptionTargetRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SubscriptionTarget `json:"body,omitempty"`
}

func (o CreateSubscriptionTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionTargetRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionTargetRequest", string(data)}, " ")
}
