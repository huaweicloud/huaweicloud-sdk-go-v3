package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAnalysisRequestBody struct {

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 查询开始时间点
	From int64 `json:"from"`

	// 查询返回的原始日志条数，最大值为500
	Limit int32 `json:"limit"`

	// 查询偏移值
	Offset int32 `json:"offset"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 查询语句
	Query string `json:"query"`

	// 查询类型：SQL, CBSL.
	QueryType CreateAnalysisRequestBodyQueryType `json:"query_type"`

	// 查询结束时间点
	To int64 `json:"to"`
}

func (o CreateAnalysisRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnalysisRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAnalysisRequestBody", string(data)}, " ")
}

type CreateAnalysisRequestBodyQueryType struct {
	value string
}

type CreateAnalysisRequestBodyQueryTypeEnum struct {
	SQL  CreateAnalysisRequestBodyQueryType
	CBSL CreateAnalysisRequestBodyQueryType
}

func GetCreateAnalysisRequestBodyQueryTypeEnum() CreateAnalysisRequestBodyQueryTypeEnum {
	return CreateAnalysisRequestBodyQueryTypeEnum{
		SQL: CreateAnalysisRequestBodyQueryType{
			value: "SQL",
		},
		CBSL: CreateAnalysisRequestBodyQueryType{
			value: "CBSL",
		},
	}
}

func (c CreateAnalysisRequestBodyQueryType) Value() string {
	return c.value
}

func (c CreateAnalysisRequestBodyQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAnalysisRequestBodyQueryType) UnmarshalJSON(b []byte) error {
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
