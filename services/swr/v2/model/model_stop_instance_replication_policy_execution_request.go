package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopInstanceReplicationPolicyExecutionRequest Request Object
type StopInstanceReplicationPolicyExecutionRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 同步策略执行记录ID
	ExecutionId int32 `json:"execution_id"`
}

func (o StopInstanceReplicationPolicyExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceReplicationPolicyExecutionRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceReplicationPolicyExecutionRequest", string(data)}, " ")
}
