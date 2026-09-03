package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskNewResponse Response Object
type DeleteExportTaskNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExportTaskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskNewResponse", string(data)}, " ")
}
