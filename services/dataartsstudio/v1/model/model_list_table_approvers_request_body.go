package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTableApproversRequestBody 获取表审批人列表参数
type ListTableApproversRequestBody struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 数据源类型,hive,dws,dli
	DatasourceType *ListTableApproversRequestBodyDatasourceType `json:"datasource_type,omitempty"`

	// 集群id,dli传DLI，dws和mrs-hive传对应的集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// schema名称，dws需要传这个字段
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 权限到期时间时间戳，毫秒。
	ExpireTime *int64 `json:"expire_time,omitempty"`
}

func (o ListTableApproversRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableApproversRequestBody struct{}"
	}

	return strings.Join([]string{"ListTableApproversRequestBody", string(data)}, " ")
}

type ListTableApproversRequestBodyDatasourceType struct {
	value string
}

type ListTableApproversRequestBodyDatasourceTypeEnum struct {
	HIVE ListTableApproversRequestBodyDatasourceType
	DWS  ListTableApproversRequestBodyDatasourceType
	DLI  ListTableApproversRequestBodyDatasourceType
}

func GetListTableApproversRequestBodyDatasourceTypeEnum() ListTableApproversRequestBodyDatasourceTypeEnum {
	return ListTableApproversRequestBodyDatasourceTypeEnum{
		HIVE: ListTableApproversRequestBodyDatasourceType{
			value: "hive",
		},
		DWS: ListTableApproversRequestBodyDatasourceType{
			value: "dws",
		},
		DLI: ListTableApproversRequestBodyDatasourceType{
			value: "dli",
		},
	}
}

func (c ListTableApproversRequestBodyDatasourceType) Value() string {
	return c.value
}

func (c ListTableApproversRequestBodyDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableApproversRequestBodyDatasourceType) UnmarshalJSON(b []byte) error {
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
