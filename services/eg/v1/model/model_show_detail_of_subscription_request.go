package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfSubscriptionRequest Request Object
type ShowDetailOfSubscriptionRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDetailOfSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfSubscriptionRequest", string(data)}, " ")
}
