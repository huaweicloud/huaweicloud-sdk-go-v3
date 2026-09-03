package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFullSqlExportTaskObsFileRequestBody 删除全量SQL导出任务OBS文件请求体
type DeleteFullSqlExportTaskObsFileRequestBody struct {

	// 导出任务ID
	Id int64 `json:"id"`
}

func (o DeleteFullSqlExportTaskObsFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFullSqlExportTaskObsFileRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteFullSqlExportTaskObsFileRequestBody", string(data)}, " ")
}
