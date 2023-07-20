package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFlinkTemplateRequestBody 新建模板
type CreateFlinkTemplateRequestBody struct {

	// 模块名称，长度限制：0-57个字符 。
	Name string `json:"name"`

	// 模板描述，长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

	// Stream SQL语句，至少包含source，query，sink三个部分, 长度限制：0-2048个字符。
	SqlBody *string `json:"sql_body,omitempty"`

	// 标签
	Tags *[]TmsTagEntity `json:"tags,omitempty"`

	// 作业模板的类型，默认为flink_sql_job，仅支持flink_sql_job和flink_opensource_sql_job
	JobType *CreateFlinkTemplateRequestBodyJobType `json:"job_type,omitempty"`
}

func (o CreateFlinkTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkTemplateRequestBody", string(data)}, " ")
}

type CreateFlinkTemplateRequestBodyJobType struct {
	value string
}

type CreateFlinkTemplateRequestBodyJobTypeEnum struct {
	FLINK_SQL_JOB            CreateFlinkTemplateRequestBodyJobType
	FLINK_OPENSOURCE_SQL_JOB CreateFlinkTemplateRequestBodyJobType
}

func GetCreateFlinkTemplateRequestBodyJobTypeEnum() CreateFlinkTemplateRequestBodyJobTypeEnum {
	return CreateFlinkTemplateRequestBodyJobTypeEnum{
		FLINK_SQL_JOB: CreateFlinkTemplateRequestBodyJobType{
			value: "flink_sql_job",
		},
		FLINK_OPENSOURCE_SQL_JOB: CreateFlinkTemplateRequestBodyJobType{
			value: "flink_opensource_sql_job",
		},
	}
}

func (c CreateFlinkTemplateRequestBodyJobType) Value() string {
	return c.value
}

func (c CreateFlinkTemplateRequestBodyJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlinkTemplateRequestBodyJobType) UnmarshalJSON(b []byte) error {
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
