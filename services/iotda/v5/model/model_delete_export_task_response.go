package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskResponse Response Object
type DeleteExportTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskResponse", string(data)}, " ")
}
