package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlJobResponse Response Object
type CreateSqlJobResponse struct {

	// 当“job_type”为“DCL”时，为请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 此SQL语句将生成并提交一个新作业，返回此作业的ID，可用于获取作业状态和作业结果。
	JobId *string `json:"job_id,omitempty"`

	// 作业类型。  DDL DCL IMPORT EXPORT QUERY INSERT
	JobType *CreateSqlJobResponseJobType `json:"job_type,omitempty"`

	// 当语句类型为DDL时，返回其结果的列名称及类型。
	Schema *[]interface{} `json:"schema,omitempty"`

	// 当语句类型为DDL时，直接返回其执行结果。
	Rows *[][]interface{} `json:"rows,omitempty"`

	// 表示作业执行方式，是同步还是异步的
	JobMode        *string `json:"job_mode,omitempty"`
	HttpStatusCode int     `json:"-"`
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
	DDL    CreateSqlJobResponseJobType
	DCL    CreateSqlJobResponseJobType
	IMPORT CreateSqlJobResponseJobType
	EXPORT CreateSqlJobResponseJobType
	QUERY  CreateSqlJobResponseJobType
	INSERT CreateSqlJobResponseJobType
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
