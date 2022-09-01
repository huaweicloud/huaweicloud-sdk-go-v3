package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScriptInfo struct {
	Name *string `json:"name,omitempty" xml:"name"`

	// 脚本类型
	Type *ScriptInfoType `json:"type,omitempty" xml:"type"`

	// 脚本关联的目录
	Directory *string `json:"directory,omitempty" xml:"directory"`

	// 脚本内容
	Content *string `json:"content,omitempty" xml:"content"`

	// 脚本关联的连接名称
	ConnectionName *string `json:"connectionName,omitempty" xml:"connectionName"`

	// 脚本执行所在的数据库
	Database *string `json:"database,omitempty" xml:"database"`

	// 脚本关联的DLI队列名称
	QueueName *string `json:"queueName,omitempty" xml:"queueName"`

	// 脚本的配置项参数
	Configuration *string `json:"configuration,omitempty" xml:"configuration"`
}

func (o ScriptInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptInfo struct{}"
	}

	return strings.Join([]string{"ScriptInfo", string(data)}, " ")
}

type ScriptInfoType struct {
	value string
}

type ScriptInfoTypeEnum struct {
	FLINK_SQL ScriptInfoType
	DLISQL    ScriptInfoType
	SPARK_SQL ScriptInfoType
	HIVE_SQL  ScriptInfoType
	DWSSQL    ScriptInfoType
	SHELL     ScriptInfoType
}

func GetScriptInfoTypeEnum() ScriptInfoTypeEnum {
	return ScriptInfoTypeEnum{
		FLINK_SQL: ScriptInfoType{
			value: "FlinkSQL",
		},
		DLISQL: ScriptInfoType{
			value: "DLISQL",
		},
		SPARK_SQL: ScriptInfoType{
			value: "SparkSQL",
		},
		HIVE_SQL: ScriptInfoType{
			value: "HiveSQL",
		},
		DWSSQL: ScriptInfoType{
			value: "DWSSQL",
		},
		SHELL: ScriptInfoType{
			value: "Shell",
		},
	}
}

func (c ScriptInfoType) Value() string {
	return c.value
}

func (c ScriptInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptInfoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
