package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteKillPhysicalProcessesRequest Request Object
type ExecuteKillPhysicalProcessesRequest struct {

	// 关联RDS的ID。
	InstanceId string `json:"instance_id"`

	Body *KillProcessesOpenRequest `json:"body,omitempty"`
}

func (o ExecuteKillPhysicalProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteKillPhysicalProcessesRequest struct{}"
	}

	return strings.Join([]string{"ExecuteKillPhysicalProcessesRequest", string(data)}, " ")
}
