package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteInstanceReplicationPolicyRequest Request Object
type ExecuteInstanceReplicationPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateReplicationExecutionRequestBody `json:"body,omitempty"`
}

func (o ExecuteInstanceReplicationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInstanceReplicationPolicyRequest struct{}"
	}

	return strings.Join([]string{"ExecuteInstanceReplicationPolicyRequest", string(data)}, " ")
}
