package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateSubscriptionRequest Request Object
type OperateSubscriptionRequest struct {

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SubscriptionOperateReq `json:"body,omitempty"`
}

func (o OperateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"OperateSubscriptionRequest", string(data)}, " ")
}
