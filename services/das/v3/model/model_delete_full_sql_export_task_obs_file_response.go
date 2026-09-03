package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFullSqlExportTaskObsFileResponse Response Object
type DeleteFullSqlExportTaskObsFileResponse struct {

	// 是否删除成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteFullSqlExportTaskObsFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFullSqlExportTaskObsFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteFullSqlExportTaskObsFileResponse", string(data)}, " ")
}
