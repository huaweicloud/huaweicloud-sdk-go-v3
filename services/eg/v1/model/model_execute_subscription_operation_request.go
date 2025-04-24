package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSubscriptionOperationRequest Request Object
type ExecuteSubscriptionOperationRequest struct {

	// 创建订阅时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SubscriptionOperateReq `json:"body,omitempty"`
}

func (o ExecuteSubscriptionOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSubscriptionOperationRequest struct{}"
	}

	return strings.Join([]string{"ExecuteSubscriptionOperationRequest", string(data)}, " ")
}
