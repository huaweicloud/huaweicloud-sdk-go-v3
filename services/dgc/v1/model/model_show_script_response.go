package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ShowScriptResponse struct {
	Name *string `json:"name,omitempty"`
	// 脚本类型

	Type *ShowScriptResponseType `json:"type,omitempty"`
	// 脚本关联的目录

	Directory *string `json:"directory,omitempty"`
	// 脚本内容

	Content *string `json:"content,omitempty"`
	// 脚本关联的连接名称

	ConnectionName *string `json:"connectionName,omitempty"`
	// 脚本执行所在的数据库

	Database *string `json:"database,omitempty"`
	// 脚本关联的DLI队列名称

	QueueName *string `json:"queueName,omitempty"`
	// 脚本的配置项参数

	Configuration  *string `json:"configuration,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptResponse struct{}"
	}

	return strings.Join([]string{"ShowScriptResponse", string(data)}, " ")
}

type ShowScriptResponseType struct {
	value string
}

type ShowScriptResponseTypeEnum struct {
	FLINK_SQL ShowScriptResponseType
	DLISQL    ShowScriptResponseType
	SPARK_SQL ShowScriptResponseType
	HIVE_SQL  ShowScriptResponseType
	DWSSQL    ShowScriptResponseType
	SHELL     ShowScriptResponseType
}

func GetShowScriptResponseTypeEnum() ShowScriptResponseTypeEnum {
	return ShowScriptResponseTypeEnum{
		FLINK_SQL: ShowScriptResponseType{
			value: "FlinkSQL",
		},
		DLISQL: ShowScriptResponseType{
			value: "DLISQL",
		},
		SPARK_SQL: ShowScriptResponseType{
			value: "SparkSQL",
		},
		HIVE_SQL: ShowScriptResponseType{
			value: "HiveSQL",
		},
		DWSSQL: ShowScriptResponseType{
			value: "DWSSQL",
		},
		SHELL: ShowScriptResponseType{
			value: "Shell",
		},
	}
}

func (c ShowScriptResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowScriptResponseType) UnmarshalJSON(b []byte) error {
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
