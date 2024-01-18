package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowScriptResponse Response Object
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
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	// 脚本描述，长度不能超过255个字符
	Description *string `json:"description,omitempty"`

	// 责任人名称
	Owner *string `json:"owner,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建脚本的目标状态，有三种状态：SAVED、SUBMITTED和PRODUCTION，分别表示脚本创建后是保存态，提交态，生产态。
	TargetStatus *ShowScriptResponseTargetStatus `json:"targetStatus,omitempty"`

	// 在开启审批开关后，需要填写该字段，表示脚本审批人
	Approvers      *[]JobApprover `json:"approvers,omitempty"`
	HttpStatusCode int            `json:"-"`
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
	FLINK_SQL       ShowScriptResponseType
	DLISQL          ShowScriptResponseType
	SPARK_SQL       ShowScriptResponseType
	HIVE_SQL        ShowScriptResponseType
	DWSSQL          ShowScriptResponseType
	RDSSQL          ShowScriptResponseType
	SHELL           ShowScriptResponseType
	PRESTO          ShowScriptResponseType
	CLICK_HOUSE_SQL ShowScriptResponseType
	HETU_ENGINE_SQL ShowScriptResponseType
	PYTHON          ShowScriptResponseType
	IMPALA_SQL      ShowScriptResponseType
	SPARK_PYTHON    ShowScriptResponseType
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
		RDSSQL: ShowScriptResponseType{
			value: "RDSSQL",
		},
		SHELL: ShowScriptResponseType{
			value: "Shell",
		},
		PRESTO: ShowScriptResponseType{
			value: "PRESTO",
		},
		CLICK_HOUSE_SQL: ShowScriptResponseType{
			value: "ClickHouseSQL",
		},
		HETU_ENGINE_SQL: ShowScriptResponseType{
			value: "HetuEngineSQL",
		},
		PYTHON: ShowScriptResponseType{
			value: "PYTHON",
		},
		IMPALA_SQL: ShowScriptResponseType{
			value: "ImpalaSQL",
		},
		SPARK_PYTHON: ShowScriptResponseType{
			value: "SparkPython",
		},
	}
}

func (c ShowScriptResponseType) Value() string {
	return c.value
}

func (c ShowScriptResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowScriptResponseType) UnmarshalJSON(b []byte) error {
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

type ShowScriptResponseTargetStatus struct {
	value string
}

type ShowScriptResponseTargetStatusEnum struct {
	SAVED      ShowScriptResponseTargetStatus
	SUBMITTED  ShowScriptResponseTargetStatus
	PRODUCTION ShowScriptResponseTargetStatus
}

func GetShowScriptResponseTargetStatusEnum() ShowScriptResponseTargetStatusEnum {
	return ShowScriptResponseTargetStatusEnum{
		SAVED: ShowScriptResponseTargetStatus{
			value: "SAVED",
		},
		SUBMITTED: ShowScriptResponseTargetStatus{
			value: "SUBMITTED",
		},
		PRODUCTION: ShowScriptResponseTargetStatus{
			value: "PRODUCTION",
		},
	}
}

func (c ShowScriptResponseTargetStatus) Value() string {
	return c.value
}

func (c ShowScriptResponseTargetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowScriptResponseTargetStatus) UnmarshalJSON(b []byte) error {
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
