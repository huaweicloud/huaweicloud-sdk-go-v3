package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteExportTaskResponse Response Object
type ExecuteExportTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteExportTaskResponse struct{}"
	}

	return strings.Join([]string{"ExecuteExportTaskResponse", string(data)}, " ")
}
