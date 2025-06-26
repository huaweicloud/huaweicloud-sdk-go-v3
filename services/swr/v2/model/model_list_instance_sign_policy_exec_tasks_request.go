package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSignPolicyExecTasksRequest Request Object
type ListInstanceSignPolicyExecTasksRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 策略ID
	PolicyId int32 `json:"policy_id"`

	// 签名策略执行记录ID
	ExecutionId int32 `json:"execution_id"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceSignPolicyExecTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSignPolicyExecTasksRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceSignPolicyExecTasksRequest", string(data)}, " ")
}
