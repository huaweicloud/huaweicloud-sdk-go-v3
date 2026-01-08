package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTasksResponse Response Object
type DeleteExportTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExportTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteExportTasksResponse", string(data)}, " ")
}
