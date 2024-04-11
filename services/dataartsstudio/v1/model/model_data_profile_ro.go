package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataProfileRo 更新概要请求体
type DataProfileRo struct {

	// 连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据源类型
	DbType *DataProfileRoDbType `json:"db_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 表ID
	TableId *string `json:"table_id,omitempty"`

	// 待更新概要字段列表
	ColumnNames *[]TableColumnDto `json:"column_names,omitempty"`

	// 字段名称列表
	FieldsName *[]string `json:"fields_name,omitempty"`

	// 执行更新语句队列
	Queue *string `json:"queue,omitempty"`

	// 是否采集唯一值
	Unique *bool `json:"unique,omitempty"`

	// 时间档案
	TimeProfile *bool `json:"time_profile,omitempty"`

	// 数据采样方式
	Sample *DataProfileRoSample `json:"sample,omitempty"`

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 是否取消
	Cancel *bool `json:"cancel,omitempty"`

	// 是否自动停止
	AutoStop *bool `json:"auto_stop,omitempty"`

	Obsconfig *ObsCommonConfig `json:"obsconfig,omitempty"`
}

func (o DataProfileRo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataProfileRo struct{}"
	}

	return strings.Join([]string{"DataProfileRo", string(data)}, " ")
}

type DataProfileRoDbType struct {
	value string
}

type DataProfileRoDbTypeEnum struct {
	DWS DataProfileRoDbType
	DLI DataProfileRoDbType
}

func GetDataProfileRoDbTypeEnum() DataProfileRoDbTypeEnum {
	return DataProfileRoDbTypeEnum{
		DWS: DataProfileRoDbType{
			value: "DWS",
		},
		DLI: DataProfileRoDbType{
			value: "DLI",
		},
	}
}

func (c DataProfileRoDbType) Value() string {
	return c.value
}

func (c DataProfileRoDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataProfileRoDbType) UnmarshalJSON(b []byte) error {
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

type DataProfileRoSample struct {
	value string
}

type DataProfileRoSampleEnum struct {
	N0UNDEFINED DataProfileRoSample
}

func GetDataProfileRoSampleEnum() DataProfileRoSampleEnum {
	return DataProfileRoSampleEnum{
		N0UNDEFINED: DataProfileRoSample{
			value: "基于采样数据，采样数量为N条：非0整数，默认值undefined",
		},
	}
}

func (c DataProfileRoSample) Value() string {
	return c.value
}

func (c DataProfileRoSample) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataProfileRoSample) UnmarshalJSON(b []byte) error {
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
