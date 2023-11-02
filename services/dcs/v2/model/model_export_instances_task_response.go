package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstancesTaskResponse Response Object
type ExportInstancesTaskResponse struct {

	// 导出实例的任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportInstancesTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstancesTaskResponse struct{}"
	}

	return strings.Join([]string{"ExportInstancesTaskResponse", string(data)}, " ")
}
