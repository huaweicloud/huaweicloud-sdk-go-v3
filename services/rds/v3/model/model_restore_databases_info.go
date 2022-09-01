package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreDatabasesInfo struct {

	// 库名
	Database string `json:"database" xml:"database"`

	// 表信息
	Tables []RestoreTableInfo `json:"tables" xml:"tables"`
}

func (o RestoreDatabasesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDatabasesInfo struct{}"
	}

	return strings.Join([]string{"RestoreDatabasesInfo", string(data)}, " ")
}
