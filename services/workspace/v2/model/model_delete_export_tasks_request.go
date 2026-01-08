package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTasksRequest Request Object
type DeleteExportTasksRequest struct {
	Body *DeleteExportTaskRequestBody `json:"body,omitempty"`
}

func (o DeleteExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTasksRequest struct{}"
	}

	return strings.Join([]string{"DeleteExportTasksRequest", string(data)}, " ")
}
