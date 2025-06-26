package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceReplicationPolicyRequest Request Object
type DeleteInstanceReplicationPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 同步策略ID
	PolicyId int32 `json:"policy_id"`
}

func (o DeleteInstanceReplicationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceReplicationPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceReplicationPolicyRequest", string(data)}, " ")
}
