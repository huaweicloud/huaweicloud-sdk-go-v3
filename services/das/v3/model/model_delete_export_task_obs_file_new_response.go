package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskObsFileNewResponse Response Object
type DeleteExportTaskObsFileNewResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteExportTaskObsFileNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskObsFileNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskObsFileNewResponse", string(data)}, " ")
}
