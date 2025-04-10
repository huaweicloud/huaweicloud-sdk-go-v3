package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityListUserTableListTableList struct {

	// 必填，数据源类型，枚举：HIVE[、DLI](tag:nohcs)、DWS
	DatasourceType string `json:"datasource_type"`

	// 数据源集群id ，Mrs和dws数据源必填[，dli数据源可不传](tag:nohcs)
	ClusterId *string `json:"cluster_id,omitempty"`

	// 必填，数据库名称
	DatabaseName string `json:"database_name"`

	// Mrs[、dli数据源](tag:nohcs)非必填，dws数据源必填，dws数据源Schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 必填，表名称
	TableName string `json:"table_name"`
}

func (o SecurityListUserTableListTableList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityListUserTableListTableList struct{}"
	}

	return strings.Join([]string{"SecurityListUserTableListTableList", string(data)}, " ")
}
