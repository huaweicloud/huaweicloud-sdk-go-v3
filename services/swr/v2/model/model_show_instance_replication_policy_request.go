package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceReplicationPolicyRequest Request Object
type ShowInstanceReplicationPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 策略ID
	PolicyId int32 `json:"policy_id"`
}

func (o ShowInstanceReplicationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceReplicationPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceReplicationPolicyRequest", string(data)}, " ")
}
