package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScriptInfo struct {
	Name *string `json:"name,omitempty"`

	// 脚本类型
	Type *ScriptInfoType `json:"type,omitempty"`

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
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	// 脚本描述，长度不能超过255个字符
	Description *string `json:"description,omitempty"`

	// 责任人名称
	Owner *string `json:"owner,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建脚本的目标状态，有三种状态：SAVED、SUBMITTED和PRODUCTION，分别表示脚本创建后是保存态，提交态，生产态。
	TargetStatus *ScriptInfoTargetStatus `json:"targetStatus,omitempty"`

	// 在开启审批开关后，需要填写该字段，表示脚本审批人
	Approvers *[]JobApprover `json:"approvers,omitempty"`
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
	FLINK_SQL       ScriptInfoType
	DLISQL          ScriptInfoType
	SPARK_SQL       ScriptInfoType
	HIVE_SQL        ScriptInfoType
	DWSSQL          ScriptInfoType
	RDSSQL          ScriptInfoType
	SHELL           ScriptInfoType
	PRESTO          ScriptInfoType
	CLICK_HOUSE_SQL ScriptInfoType
	HETU_ENGINE_SQL ScriptInfoType
	PYTHON          ScriptInfoType
	IMPALA_SQL      ScriptInfoType
	SPARK_PYTHON    ScriptInfoType
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
		RDSSQL: ScriptInfoType{
			value: "RDSSQL",
		},
		SHELL: ScriptInfoType{
			value: "Shell",
		},
		PRESTO: ScriptInfoType{
			value: "PRESTO",
		},
		CLICK_HOUSE_SQL: ScriptInfoType{
			value: "ClickHouseSQL",
		},
		HETU_ENGINE_SQL: ScriptInfoType{
			value: "HetuEngineSQL",
		},
		PYTHON: ScriptInfoType{
			value: "PYTHON",
		},
		IMPALA_SQL: ScriptInfoType{
			value: "ImpalaSQL",
		},
		SPARK_PYTHON: ScriptInfoType{
			value: "SparkPython",
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

type ScriptInfoTargetStatus struct {
	value string
}

type ScriptInfoTargetStatusEnum struct {
	SAVED      ScriptInfoTargetStatus
	SUBMITTED  ScriptInfoTargetStatus
	PRODUCTION ScriptInfoTargetStatus
}

func GetScriptInfoTargetStatusEnum() ScriptInfoTargetStatusEnum {
	return ScriptInfoTargetStatusEnum{
		SAVED: ScriptInfoTargetStatus{
			value: "SAVED",
		},
		SUBMITTED: ScriptInfoTargetStatus{
			value: "SUBMITTED",
		},
		PRODUCTION: ScriptInfoTargetStatus{
			value: "PRODUCTION",
		},
	}
}

func (c ScriptInfoTargetStatus) Value() string {
	return c.value
}

func (c ScriptInfoTargetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptInfoTargetStatus) UnmarshalJSON(b []byte) error {
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
