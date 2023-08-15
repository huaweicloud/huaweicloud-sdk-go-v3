package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceMonitorResponse Response Object
type UpdateInstanceMonitorResponse struct {

	// 修改秒级监控的任务流ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceMonitorResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceMonitorResponse", string(data)}, " ")
}
