package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRestoreTablesRequestBody struct {

	// 备份时间点，时间戳格式。
	RestoreTime string `json:"restore_time"`

	// 是否是最新库表。默认为false。 - true：是最新库表。 - false：是恢复时间点库表。
	LastTableInfo *string `json:"last_table_info,omitempty"`

	// 数据库信息。
	RestoreTables []CreateRestoreDatabaseTableInfo `json:"restore_tables"`
}

func (o CreateRestoreTablesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreTablesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRestoreTablesRequestBody", string(data)}, " ")
}
