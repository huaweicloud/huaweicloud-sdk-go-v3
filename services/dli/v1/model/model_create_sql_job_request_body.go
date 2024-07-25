package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlJobRequestBody 提交SQL作业
type CreateSqlJobRequestBody struct {

	// 待执行的SQL语句。
	Sql string `json:"sql"`

	// 待提交作业的队列引擎名称，名称只能包含英文字母。
	EngineType *CreateSqlJobRequestBodyEngineType `json:"engine_type,omitempty"`

	// 待提交作业的表默认catalog。
	CurrentCatalog *string `json:"current_catalog,omitempty"`

	// SQL语句执行所在的数据库。当创建新数据库时，不需要提供此参数。
	Currentdb *string `json:"currentdb,omitempty"`

	// 待提交作业的队列名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头。
	QueueName *string `json:"queue_name,omitempty"`

	// 用户以“key/value”的形式设置用于此作业的配置参数。目前支持的配置项请参考表3。
	Conf *[]string `json:"conf,omitempty"`

	// 作业标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateSqlJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlJobRequestBody", string(data)}, " ")
}

type CreateSqlJobRequestBodyEngineType struct {
	value string
}

type CreateSqlJobRequestBodyEngineTypeEnum struct {
	HETU_ENGINE CreateSqlJobRequestBodyEngineType
	SPARK       CreateSqlJobRequestBodyEngineType
}

func GetCreateSqlJobRequestBodyEngineTypeEnum() CreateSqlJobRequestBodyEngineTypeEnum {
	return CreateSqlJobRequestBodyEngineTypeEnum{
		HETU_ENGINE: CreateSqlJobRequestBodyEngineType{
			value: "hetuEngine",
		},
		SPARK: CreateSqlJobRequestBodyEngineType{
			value: "spark",
		},
	}
}

func (c CreateSqlJobRequestBodyEngineType) Value() string {
	return c.value
}

func (c CreateSqlJobRequestBodyEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlJobRequestBodyEngineType) UnmarshalJSON(b []byte) error {
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
