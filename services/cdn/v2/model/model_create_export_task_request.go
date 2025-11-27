package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExportTaskRequest Request Object
type CreateExportTaskRequest struct {
	Body *ExportTaskVo `json:"body,omitempty"`
}

func (o CreateExportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExportTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateExportTaskRequest", string(data)}, " ")
}
