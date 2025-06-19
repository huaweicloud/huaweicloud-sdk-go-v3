package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlJobResponse Response Object
type CreateSqlJobResponse struct {

	// 参数解释: 当“job_type”为“DCL”时，为请求执行是否成功。“true”表示请求执行成功 示例: true 约束限制:  无 取值范围: true,false 默认取值: 无
	IsSuccess *bool `json:"is_success,omitempty"`

	// 参数解释: 系统提示信息，执行成功时，信息可能为空 示例: success 约束限制:  无 取值范围: 无 默认取值: 无
	Message *string `json:"message,omitempty"`

	// 参数解释: 此SQL语句将生成并提交一个新作业，返回此作业的ID，可用于获取作业状态和作业结果 示例: 8ecb0777-9c70-4529-9935-29ea0946039c 约束限制:  无 取值范围: 无 默认取值: 无
	JobId *string `json:"job_id,omitempty"`

	// 参数解释:  指定查询的作业类型，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE，若要查询所有类型的作业，则传入ALL。 示例: QUERY 约束限制:  无 取值范围: DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE、ALL 默认取值: 无
	JobType *CreateSqlJobResponseJobType `json:"job_type,omitempty"`

	// 参数解释:  当语句类型为DDL时，返回其结果的列名称及类型 示例: [{\"col_name\": \"string\"},{\"data_type\": \"string\"},{\"comment\": \"string\"}] 约束限制:  无 取值范围: 无 默认取值: 无
	Schema *[]interface{} `json:"schema,omitempty"`

	// 参数解释:  当语句类型为DDL时，直接返回其执行结果 示例: [[\"c1\",\"int\",null],[\"c2\",\"string\",null]] 约束限制:  无 取值范围: 无 默认取值: 无
	Rows *[][]interface{} `json:"rows,omitempty"`

	// 参数解释:  表示作业执行方式，是同步还是异步的 示例: async 约束限制:  无 取值范围: async（异步） sync（同步） 默认取值: 无
	JobMode        *CreateSqlJobResponseJobMode `json:"job_mode,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CreateSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlJobResponse", string(data)}, " ")
}

type CreateSqlJobResponseJobType struct {
	value string
}

type CreateSqlJobResponseJobTypeEnum struct {
	DDL            CreateSqlJobResponseJobType
	DCL            CreateSqlJobResponseJobType
	IMPORT         CreateSqlJobResponseJobType
	EXPORT         CreateSqlJobResponseJobType
	QUERY          CreateSqlJobResponseJobType
	INSERT         CreateSqlJobResponseJobType
	DATA_MIGRATION CreateSqlJobResponseJobType
	UPDATE         CreateSqlJobResponseJobType
	DELETE         CreateSqlJobResponseJobType
	RESTART_QUEUE  CreateSqlJobResponseJobType
	SCALE_QUEUE    CreateSqlJobResponseJobType
	ALL            CreateSqlJobResponseJobType
}

func GetCreateSqlJobResponseJobTypeEnum() CreateSqlJobResponseJobTypeEnum {
	return CreateSqlJobResponseJobTypeEnum{
		DDL: CreateSqlJobResponseJobType{
			value: "DDL",
		},
		DCL: CreateSqlJobResponseJobType{
			value: "DCL",
		},
		IMPORT: CreateSqlJobResponseJobType{
			value: "IMPORT",
		},
		EXPORT: CreateSqlJobResponseJobType{
			value: "EXPORT",
		},
		QUERY: CreateSqlJobResponseJobType{
			value: "QUERY",
		},
		INSERT: CreateSqlJobResponseJobType{
			value: "INSERT",
		},
		DATA_MIGRATION: CreateSqlJobResponseJobType{
			value: "DATA_MIGRATION",
		},
		UPDATE: CreateSqlJobResponseJobType{
			value: "UPDATE",
		},
		DELETE: CreateSqlJobResponseJobType{
			value: "DELETE",
		},
		RESTART_QUEUE: CreateSqlJobResponseJobType{
			value: "RESTART_QUEUE",
		},
		SCALE_QUEUE: CreateSqlJobResponseJobType{
			value: "SCALE_QUEUE",
		},
		ALL: CreateSqlJobResponseJobType{
			value: "ALL",
		},
	}
}

func (c CreateSqlJobResponseJobType) Value() string {
	return c.value
}

func (c CreateSqlJobResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlJobResponseJobType) UnmarshalJSON(b []byte) error {
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

type CreateSqlJobResponseJobMode struct {
	value string
}

type CreateSqlJobResponseJobModeEnum struct {
	ASYNC CreateSqlJobResponseJobMode
	SYNC  CreateSqlJobResponseJobMode
}

func GetCreateSqlJobResponseJobModeEnum() CreateSqlJobResponseJobModeEnum {
	return CreateSqlJobResponseJobModeEnum{
		ASYNC: CreateSqlJobResponseJobMode{
			value: "async",
		},
		SYNC: CreateSqlJobResponseJobMode{
			value: "sync",
		},
	}
}

func (c CreateSqlJobResponseJobMode) Value() string {
	return c.value
}

func (c CreateSqlJobResponseJobMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlJobResponseJobMode) UnmarshalJSON(b []byte) error {
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
