package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataCompareDatabaseObject 数据加工对象信息
type CreateDataCompareDatabaseObject struct {

	// 两层数据库场景： 数据库名称和数据库表名称，例如格式为t_auto_db-*-*-users，其中t_auto_db为数据库名称，users为表名称。  三层数据库场景： 数据库名称、数据库schema名称、数据库表名称，例如格式为t_auto_db-*-*-schema1-*-*-users，其中t_auto_db为数据库名称，schema1为数据库schema名称，users为表名称。
	Id *string `json:"id,omitempty"`
}

func (o CreateDataCompareDatabaseObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataCompareDatabaseObject struct{}"
	}

	return strings.Join([]string{"CreateDataCompareDatabaseObject", string(data)}, " ")
}
