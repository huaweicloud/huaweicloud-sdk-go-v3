package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteClusterSwitchoverResponse Response Object
type ExecuteClusterSwitchoverResponse struct {

	// 集群分片倒换的任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteClusterSwitchoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterSwitchoverResponse struct{}"
	}

	return strings.Join([]string{"ExecuteClusterSwitchoverResponse", string(data)}, " ")
}
