package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteImportTaskResponse Response Object
type ExecuteImportTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteImportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteImportTaskResponse struct{}"
	}

	return strings.Join([]string{"ExecuteImportTaskResponse", string(data)}, " ")
}
