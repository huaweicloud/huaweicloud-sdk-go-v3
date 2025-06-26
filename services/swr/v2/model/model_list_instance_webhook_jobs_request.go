package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceWebhookJobsRequest Request Object
type ListInstanceWebhookJobsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 触发器策略ID
	PolicyId int32 `json:"policy_id"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`

	// 状态，可选Initialized、Pending、InProgress、Succeed、Failed、Stopped
	Status *string `json:"status,omitempty"`
}

func (o ListInstanceWebhookJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceWebhookJobsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceWebhookJobsRequest", string(data)}, " ")
}
