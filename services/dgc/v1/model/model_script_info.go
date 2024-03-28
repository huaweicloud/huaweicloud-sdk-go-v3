package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScriptInfo struct {
	Name string `json:"name"`

	// 脚本类型： - FlinkSQL - DLISQL - SparkSQL - HiveSQL - DWSSQL - RDSSQL - Shell - PRESTO - ClickHouseSQL - HetuEngineSQL - PYTHON - ImpalaSQL - SparkPython
	Type ScriptInfoType `json:"type"`

	// 脚本关联的目录。通过DataArts Studio管理控制台 > 数据开发，左侧列表选择“数据开发 > 脚本开发”。在脚本的目录树上，可以查看到当前已经创建的目录，默认在根目录/。
	Directory *string `json:"directory,omitempty"`

	// 脚本内容。最大支持4M。
	Content string `json:"content"`

	// 脚本关联的连接名称。当type参数值为DLISQL、SparkSQL、HiveSQL、DWSSQL、Shell、PRESTO、ClickHouseSQL、HetuEngineSQL、RDSSQL、ImpalaSQL、PYTHON、SparkPython其中之一时，这个参数是必选的。用户可以通过查询连接列表（待下线）接口获取当前系统中已经存在的连接。默认值为空。
	ConnectionName *string `json:"connectionName,omitempty"`

	// 脚本执行所在的数据库。当type参数值为DLISQL、SparkSQL、HiveSQL、DWSSQL、PRESTO、ClickHouseSQL、ImpalaSQL、HetuEngineSQL、RDSSQL其中之一时，才支持此参数。type为DLI SQL时，可以通过查看所有数据库接口获取数据库信息。type为其他类型的时候，必选。
	Database *string `json:"database,omitempty"`

	// 脚本关联的DLI队列名称。当type参数值为DLI SQL时，才支持此参数。可以通过查询队列列表接口获取队列信息。默认值为空。
	QueueName *string `json:"queueName,omitempty"`

	// 脚本的配置项参数。当type参数值为DLISQL时，才支持此参数。当前支持的配置项列表请参考 conf参数说明。默认值为空。
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	// 脚本描述，长度不能超过255个字符
	Description *string `json:"description,omitempty"`

	// 责任人名称
	Owner *string `json:"owner,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建脚本的目标状态，有三种状态：SAVED、SUBMITTED和PRODUCTION，分别表示脚本创建后是保存态，提交态，生产态: - 保存态表示脚本仅保存，无法调度运行，需要提交并审核通过后才能运行。 - 提交态表示脚本保存后会自动提交，需要审核通过才能运行。 - 生产态表示脚本跳过审批环节，创建后可以直接运行。注意：只有工作空间的管理员用户才能创建生产态的脚本。
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
