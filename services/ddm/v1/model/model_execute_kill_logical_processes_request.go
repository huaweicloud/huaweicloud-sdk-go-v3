package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteKillLogicalProcessesRequest Request Object
type ExecuteKillLogicalProcessesRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	Body *KillProcessesOpenRequest `json:"body,omitempty"`
}

func (o ExecuteKillLogicalProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteKillLogicalProcessesRequest struct{}"
	}

	return strings.Join([]string{"ExecuteKillLogicalProcessesRequest", string(data)}, " ")
}
