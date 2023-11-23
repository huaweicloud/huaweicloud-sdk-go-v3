package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreDatabaseInfos struct {

	// 数据库名称。
	Name *string `json:"name,omitempty"`

	// 总表数。
	TotalTables *int32 `json:"total_tables,omitempty"`

	// 表信息。
	Tables *[]RestoreDatabaseTableInfo `json:"tables,omitempty"`
}

func (o RestoreDatabaseInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDatabaseInfos struct{}"
	}

	return strings.Join([]string{"RestoreDatabaseInfos", string(data)}, " ")
}
