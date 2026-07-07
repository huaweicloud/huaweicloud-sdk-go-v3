package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseDeadLockRequest Request Object
type ParseDeadLockRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ParseDeadLockRequestBody `json:"body,omitempty"`
}

func (o ParseDeadLockRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseDeadLockRequest struct{}"
	}

	return strings.Join([]string{"ParseDeadLockRequest", string(data)}, " ")
}
