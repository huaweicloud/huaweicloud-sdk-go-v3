package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceReplicationPolicyRequest Request Object
type CreateInstanceReplicationPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateReplicationPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceReplicationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceReplicationPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceReplicationPolicyRequest", string(data)}, " ")
}
