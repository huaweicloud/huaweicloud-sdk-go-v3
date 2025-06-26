package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopInstanceReplicationPolicyExecutionResponse Response Object
type StopInstanceReplicationPolicyExecutionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopInstanceReplicationPolicyExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceReplicationPolicyExecutionResponse struct{}"
	}

	return strings.Join([]string{"StopInstanceReplicationPolicyExecutionResponse", string(data)}, " ")
}
