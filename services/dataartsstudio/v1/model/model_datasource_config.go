package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DatasourceConfig struct {

	// 数据源的类型
	Type *DatasourceConfigType `json:"type,omitempty"`

	// 数据连接名称
	ConnectionName *string `json:"connection_name,omitempty"`

	// 数据连接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 数据库名
	Database *string `json:"database,omitempty"`

	// 数据表名称
	Datatable *string `json:"datatable,omitempty"`

	// 数据表ID
	TableId *string `json:"table_id,omitempty"`

	// DLI的队列名称
	Queue *string `json:"queue,omitempty"`

	// 取数方式
	AccessType *DatasourceConfigAccessType `json:"access_type,omitempty"`

	// 获取数据的模式
	AccessMode *DatasourceConfigAccessMode `json:"access_mode,omitempty"`

	Pagination *DatasourceConfigPagination `json:"pagination,omitempty"`

	// 脚本模式下的sql语句
	Sql *string `json:"sql,omitempty"`

	// API后端参数
	BackendParas *[]ApiRequestPara `json:"backend_paras,omitempty"`

	// 配置类API返回参数
	ResponseParas *[]ApiResponsePara `json:"response_paras,omitempty"`

	// 排序参数
	OrderParas *[]DatasourceOrderPara `json:"order_paras,omitempty"`
}

func (o DatasourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasourceConfig struct{}"
	}

	return strings.Join([]string{"DatasourceConfig", string(data)}, " ")
}

type DatasourceConfigType struct {
	value string
}

type DatasourceConfigTypeEnum struct {
	MYSQL DatasourceConfigType
	DLI   DatasourceConfigType
	DWS   DatasourceConfigType
	HIVE  DatasourceConfigType
	HBASE DatasourceConfigType
}

func GetDatasourceConfigTypeEnum() DatasourceConfigTypeEnum {
	return DatasourceConfigTypeEnum{
		MYSQL: DatasourceConfigType{
			value: "MYSQL",
		},
		DLI: DatasourceConfigType{
			value: "DLI",
		},
		DWS: DatasourceConfigType{
			value: "DWS",
		},
		HIVE: DatasourceConfigType{
			value: "HIVE",
		},
		HBASE: DatasourceConfigType{
			value: "HBASE",
		},
	}
}

func (c DatasourceConfigType) Value() string {
	return c.value
}

func (c DatasourceConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasourceConfigType) UnmarshalJSON(b []byte) error {
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

type DatasourceConfigAccessType struct {
	value string
}

type DatasourceConfigAccessTypeEnum struct {
	SCRIPT        DatasourceConfigAccessType
	CONFIGURAITON DatasourceConfigAccessType
}

func GetDatasourceConfigAccessTypeEnum() DatasourceConfigAccessTypeEnum {
	return DatasourceConfigAccessTypeEnum{
		SCRIPT: DatasourceConfigAccessType{
			value: "SCRIPT",
		},
		CONFIGURAITON: DatasourceConfigAccessType{
			value: "CONFIGURAITON",
		},
	}
}

func (c DatasourceConfigAccessType) Value() string {
	return c.value
}

func (c DatasourceConfigAccessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasourceConfigAccessType) UnmarshalJSON(b []byte) error {
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

type DatasourceConfigAccessMode struct {
	value string
}

type DatasourceConfigAccessModeEnum struct {
	SQL           DatasourceConfigAccessMode
	ROW_KEY       DatasourceConfigAccessMode
	PREFIX_FILTER DatasourceConfigAccessMode
}

func GetDatasourceConfigAccessModeEnum() DatasourceConfigAccessModeEnum {
	return DatasourceConfigAccessModeEnum{
		SQL: DatasourceConfigAccessMode{
			value: "SQL",
		},
		ROW_KEY: DatasourceConfigAccessMode{
			value: "ROW_KEY",
		},
		PREFIX_FILTER: DatasourceConfigAccessMode{
			value: "PREFIX_FILTER",
		},
	}
}

func (c DatasourceConfigAccessMode) Value() string {
	return c.value
}

func (c DatasourceConfigAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasourceConfigAccessMode) UnmarshalJSON(b []byte) error {
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

type DatasourceConfigPagination struct {
	value string
}

type DatasourceConfigPaginationEnum struct {
	DEFAULT DatasourceConfigPagination
	CUSTOM  DatasourceConfigPagination
}

func GetDatasourceConfigPaginationEnum() DatasourceConfigPaginationEnum {
	return DatasourceConfigPaginationEnum{
		DEFAULT: DatasourceConfigPagination{
			value: "DEFAULT",
		},
		CUSTOM: DatasourceConfigPagination{
			value: "CUSTOM",
		},
	}
}

func (c DatasourceConfigPagination) Value() string {
	return c.value
}

func (c DatasourceConfigPagination) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasourceConfigPagination) UnmarshalJSON(b []byte) error {
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
