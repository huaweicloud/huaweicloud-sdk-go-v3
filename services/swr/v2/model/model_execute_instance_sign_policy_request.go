package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteInstanceSignPolicyRequest Request Object
type ExecuteInstanceSignPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 策略ID
	PolicyId int32 `json:"policy_id"`
}

func (o ExecuteInstanceSignPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInstanceSignPolicyRequest struct{}"
	}

	return strings.Join([]string{"ExecuteInstanceSignPolicyRequest", string(data)}, " ")
}
