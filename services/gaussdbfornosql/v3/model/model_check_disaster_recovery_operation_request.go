package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckDisasterRecoveryOperationRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *PrecheckDisasterRecoveryOperationBody `json:"body,omitempty"`
}

func (o CheckDisasterRecoveryOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDisasterRecoveryOperationRequest struct{}"
	}

	return strings.Join([]string{"CheckDisasterRecoveryOperationRequest", string(data)}, " ")
}
