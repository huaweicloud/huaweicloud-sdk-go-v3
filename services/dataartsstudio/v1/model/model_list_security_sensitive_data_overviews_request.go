package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecuritySensitiveDataOverviewsRequest Request Object
type ListSecuritySensitiveDataOverviewsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 数据源类型,HIVE数据源,DWS数据源,DLI数据源
	Datasource *ListSecuritySensitiveDataOverviewsRequestDatasource `json:"datasource,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`
}

func (o ListSecuritySensitiveDataOverviewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecuritySensitiveDataOverviewsRequest struct{}"
	}

	return strings.Join([]string{"ListSecuritySensitiveDataOverviewsRequest", string(data)}, " ")
}

type ListSecuritySensitiveDataOverviewsRequestDatasource struct {
	value string
}

type ListSecuritySensitiveDataOverviewsRequestDatasourceEnum struct {
	HIVE ListSecuritySensitiveDataOverviewsRequestDatasource
	DWS  ListSecuritySensitiveDataOverviewsRequestDatasource
	DLI  ListSecuritySensitiveDataOverviewsRequestDatasource
}

func GetListSecuritySensitiveDataOverviewsRequestDatasourceEnum() ListSecuritySensitiveDataOverviewsRequestDatasourceEnum {
	return ListSecuritySensitiveDataOverviewsRequestDatasourceEnum{
		HIVE: ListSecuritySensitiveDataOverviewsRequestDatasource{
			value: "HIVE",
		},
		DWS: ListSecuritySensitiveDataOverviewsRequestDatasource{
			value: "DWS",
		},
		DLI: ListSecuritySensitiveDataOverviewsRequestDatasource{
			value: "DLI",
		},
	}
}

func (c ListSecuritySensitiveDataOverviewsRequestDatasource) Value() string {
	return c.value
}

func (c ListSecuritySensitiveDataOverviewsRequestDatasource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecuritySensitiveDataOverviewsRequestDatasource) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
