package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeFailoverModeResponse struct {

	// 实例Id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId"`

	// 同步模式
	ReplicationMode *string `json:"replicationMode,omitempty" xml:"replicationMode"`

	// 任务id
	WorkflowId     *string `json:"workflowId,omitempty" xml:"workflowId"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeFailoverModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFailoverModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeFailoverModeResponse", string(data)}, " ")
}
