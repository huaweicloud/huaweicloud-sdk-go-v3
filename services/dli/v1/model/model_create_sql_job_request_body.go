package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlJobRequestBody 提交SQL作业
type CreateSqlJobRequestBody struct {

	// 参数解释: 待执行的SQL语句 示例: desc table1 约束限制:  无 取值范围: 无 默认取值: 无
	Sql string `json:"sql"`

	// 参数解释: 待提交作业的队列引擎名称，名称只能包含英文字母 示例: hetuEngine 约束限制:  无 取值范围: hetuEngine,spark 默认取值: 无
	EngineType *CreateSqlJobRequestBodyEngineType `json:"engine_type,omitempty"`

	// 参数解释: 待提交作业的表默认catalog 示例: dli 约束限制:  匹配正则表达式'^(?!_)(?![0-9]+$)[A-Za-z0-9_]*$'且长度在[0,128]范围内的字符串 取值范围: 无 默认取值: 无
	CurrentCatalog *string `json:"current_catalog,omitempty"`

	// 参数解释: SQL语句执行所在的数据库。当创建新数据库时，不需要提供此参数 示例: db1 约束限制:  匹配正则表达式'^(?!_)(?![0-9]+$)[A-Za-z0-9_]*$'且长度在[0,128]范围内的字符串 取值范围: 无 默认取值: 无
	Currentdb *string `json:"currentdb,omitempty"`

	// 参数解释: 待提交作业的队列名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头 示例: q1 约束限制:  匹配正则表达式'^(?!_)(?![0-9]+$)[A-Za-z0-9_]*$'且长度在[1,128]范围内的字符串 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释: 用户以“key/value”的形式设置用于此作业的配置参数 示例: spark.sql.files.maxRecordsPerFile = 0, spark.sql.shuffle.partitions = 200 约束限制:  元素满足key=value格式的字符串数组 取值范围: 无 默认取值: dli.sql.sqlasync.enabled = true
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
