package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExportTaskResponse Response Object
type CreateExportTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExportTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateExportTaskResponse", string(data)}, " ")
}
