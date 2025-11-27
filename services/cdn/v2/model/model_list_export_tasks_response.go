package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExportTasksResponse Response Object
type ListExportTasksResponse struct {

	// 导出任务总数
	Total *int32 `json:"total,omitempty"`

	Data           *[]ExportTask `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListExportTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExportTasksResponse struct{}"
	}

	return strings.Join([]string{"ListExportTasksResponse", string(data)}, " ")
}
