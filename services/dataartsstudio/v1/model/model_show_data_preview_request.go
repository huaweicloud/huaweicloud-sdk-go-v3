package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataPreviewRequest Request Object
type ShowDataPreviewRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 资产guid，获取方法请参见[数据资产guid](dataartsstudio_02_0351.xml)。
	Guid string `json:"guid"`

	// 连接id
	DataConnectionId *string `json:"data_connection_id,omitempty"`

	// 数据源表类型，取值范围：hive_table、dws_table、dli_table、dli_table_managed、dli_table_external、dli_table_view、mysql_table、gbase_table、postgre_table、hbase_table、dm_table、doris_table、sqlserver_table。
	DataType *string `json:"data_type,omitempty"`

	// database名称
	Database *string `json:"database,omitempty"`

	// schema名称
	Schema *string `json:"schema,omitempty"`

	// table名称
	Table *string `json:"table,omitempty"`

	// 数据源空间id
	DatasourceWorkspaceId *string `json:"datasource_workspace_id,omitempty"`

	// 分区名称，hive类型数据源可使用预览分区中数据
	PartitionsCondition *string `json:"partitions_condition,omitempty"`
}

func (o ShowDataPreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataPreviewRequest struct{}"
	}

	return strings.Join([]string{"ShowDataPreviewRequest", string(data)}, " ")
}
