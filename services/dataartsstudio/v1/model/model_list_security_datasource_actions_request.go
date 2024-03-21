package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityDatasourceActionsRequest Request Object
type ListSecurityDatasourceActionsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 父权限集ID。获取方法请参见[查询权限集列表](ListSecurityPermissionSets.xml) 注意： * 当该值为0时，则所有库表均支持查询 * 当该值为父权限集ID时，则基于父权限集中的权限查询
	ParentPermissionSetId string `json:"parent_permission_set_id"`

	// 集群ID，获取方法请参见[查询单个数据连接信息](ShowDataconnection.xml) * 查询Hive和DWS数据源操作信息时该数值为必填项，当数据源为DLI时无需填写
	ClusterId string `json:"cluster_id"`

	// 数据源类型 * HIVE数据源 * DWS数据源 * DLI数据源
	DatasourceType ListSecurityDatasourceActionsRequestDatasourceType `json:"datasource_type"`

	// 数据库名 `注意：该值作为查询关键字时，不能与url同时存在，需要指定其一进行查询`
	DatabaseName *string `json:"database_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据表名称
	TableName *string `json:"table_name,omitempty"`

	// 数据字段名称
	ColumnName *string `json:"column_name,omitempty"`

	// url路径名称 `注意：该值作为查询关键字时，不能与database_name同时存在，需要指定其一进行查询`
	Url *string `json:"url,omitempty"`
}

func (o ListSecurityDatasourceActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDatasourceActionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDatasourceActionsRequest", string(data)}, " ")
}

type ListSecurityDatasourceActionsRequestDatasourceType struct {
	value string
}

type ListSecurityDatasourceActionsRequestDatasourceTypeEnum struct {
	HIVE ListSecurityDatasourceActionsRequestDatasourceType
	DWS  ListSecurityDatasourceActionsRequestDatasourceType
	DLI  ListSecurityDatasourceActionsRequestDatasourceType
}

func GetListSecurityDatasourceActionsRequestDatasourceTypeEnum() ListSecurityDatasourceActionsRequestDatasourceTypeEnum {
	return ListSecurityDatasourceActionsRequestDatasourceTypeEnum{
		HIVE: ListSecurityDatasourceActionsRequestDatasourceType{
			value: "HIVE",
		},
		DWS: ListSecurityDatasourceActionsRequestDatasourceType{
			value: "DWS",
		},
		DLI: ListSecurityDatasourceActionsRequestDatasourceType{
			value: "DLI",
		},
	}
}

func (c ListSecurityDatasourceActionsRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityDatasourceActionsRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityDatasourceActionsRequestDatasourceType) UnmarshalJSON(b []byte) error {
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
