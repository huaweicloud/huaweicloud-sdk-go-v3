package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportScheduledTasksRecordsResponse Response Object
type ExportScheduledTasksRecordsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportScheduledTasksRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportScheduledTasksRecordsResponse struct{}"
	}

	return strings.Join([]string{"ExportScheduledTasksRecordsResponse", string(data)}, " ")
}
