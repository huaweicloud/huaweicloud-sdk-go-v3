package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionRequest Request Object
type CreateSubscriptionRequest struct {

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SubscriptionCreateReq `json:"body,omitempty"`
}

func (o CreateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionRequest", string(data)}, " ")
}
