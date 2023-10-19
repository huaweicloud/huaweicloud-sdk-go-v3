package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRestoreDatabaseTableInfo struct {

	// 数据库名称。
	Database string `json:"database"`

	// 表信息。
	Tables []CreateRestoreTableInfo `json:"tables"`
}

func (o CreateRestoreDatabaseTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreDatabaseTableInfo struct{}"
	}

	return strings.Join([]string{"CreateRestoreDatabaseTableInfo", string(data)}, " ")
}
