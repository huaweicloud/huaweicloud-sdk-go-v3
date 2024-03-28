package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFlinkSqlJobTemplateRequestBody 新建模板
type CreateFlinkSqlJobTemplateRequestBody struct {

	// 模块名称，长度限制：0-57个字符 。
	Name string `json:"name"`

	// 模板描述，长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

	// Stream SQL语句，至少包含source，query，sink三个部分, 长度限制：0-2048个字符。
	SqlBody *string `json:"sql_body,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 作业模板的类型，默认为flink_sql_job，仅支持flink_sql_job和flink_opensource_sql_job
	JobType *CreateFlinkSqlJobTemplateRequestBodyJobType `json:"job_type,omitempty"`
}

func (o CreateFlinkSqlJobTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobTemplateRequestBody", string(data)}, " ")
}

type CreateFlinkSqlJobTemplateRequestBodyJobType struct {
	value string
}

type CreateFlinkSqlJobTemplateRequestBodyJobTypeEnum struct {
	FLINK_SQL_JOB            CreateFlinkSqlJobTemplateRequestBodyJobType
	FLINK_OPENSOURCE_SQL_JOB CreateFlinkSqlJobTemplateRequestBodyJobType
}

func GetCreateFlinkSqlJobTemplateRequestBodyJobTypeEnum() CreateFlinkSqlJobTemplateRequestBodyJobTypeEnum {
	return CreateFlinkSqlJobTemplateRequestBodyJobTypeEnum{
		FLINK_SQL_JOB: CreateFlinkSqlJobTemplateRequestBodyJobType{
			value: "flink_sql_job",
		},
		FLINK_OPENSOURCE_SQL_JOB: CreateFlinkSqlJobTemplateRequestBodyJobType{
			value: "flink_opensource_sql_job",
		},
	}
}

func (c CreateFlinkSqlJobTemplateRequestBodyJobType) Value() string {
	return c.value
}

func (c CreateFlinkSqlJobTemplateRequestBodyJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlinkSqlJobTemplateRequestBodyJobType) UnmarshalJSON(b []byte) error {
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
