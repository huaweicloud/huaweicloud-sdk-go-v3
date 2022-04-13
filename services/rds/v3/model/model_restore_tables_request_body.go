package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTablesRequestBody struct {
	// 恢复时间戳

	RestoreTime int64 `json:"restoreTime"`
	// 表信息

	RestoreTables []RestoreDatabasesInfo `json:"restoreTables"`
}

func (o RestoreTablesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreTablesRequestBody", string(data)}, " ")
}
