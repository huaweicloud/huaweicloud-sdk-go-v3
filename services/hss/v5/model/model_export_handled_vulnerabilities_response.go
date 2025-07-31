package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportHandledVulnerabilitiesResponse Response Object
type ExportHandledVulnerabilitiesResponse struct {

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportHandledVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportHandledVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ExportHandledVulnerabilitiesResponse", string(data)}, " ")
}
