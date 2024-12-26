package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfSubscriptionTargetRequest Request Object
type ShowDetailOfSubscriptionTargetRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 事件订阅目标ID
	TargetId string `json:"target_id"`

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDetailOfSubscriptionTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfSubscriptionTargetRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfSubscriptionTargetRequest", string(data)}, " ")
}
