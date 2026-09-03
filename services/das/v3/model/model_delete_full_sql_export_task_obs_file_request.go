package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFullSqlExportTaskObsFileRequest Request Object
type DeleteFullSqlExportTaskObsFileRequest struct {
	Body *DeleteFullSqlExportTaskObsFileRequestBody `json:"body,omitempty"`
}

func (o DeleteFullSqlExportTaskObsFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFullSqlExportTaskObsFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteFullSqlExportTaskObsFileRequest", string(data)}, " ")
}
