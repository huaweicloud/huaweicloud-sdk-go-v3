package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceReplicationPolicyRequest Request Object
type UpdateInstanceReplicationPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 同步策略ID
	PolicyId int32 `json:"policy_id"`

	Body *UpdateReplicationPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceReplicationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReplicationPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReplicationPolicyRequest", string(data)}, " ")
}
