package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PreviewJobResultResponse Response Object
type PreviewJobResultResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	JobId *string `json:"job_id,omitempty"`

	// 作业类型，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT。  目前仅支持查看“QUERY”类型作业的执行结果。
	JobType *PreviewJobResultResponseJobType `json:"job_type,omitempty"`

	// 作业结果总条数。
	RowCount *int32 `json:"row_count,omitempty"`

	// 作业执行过程中扫描的数据量。
	InputSize *int64 `json:"input_size,omitempty"`

	// 作业结果列名称和类型。
	Schema *[]interface{} `json:"schema,omitempty"`

	// 作业结果集。
	Rows           *[][]string `json:"rows,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o PreviewJobResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewJobResultResponse struct{}"
	}

	return strings.Join([]string{"PreviewJobResultResponse", string(data)}, " ")
}

type PreviewJobResultResponseJobType struct {
	value string
}

type PreviewJobResultResponseJobTypeEnum struct {
	DDL    PreviewJobResultResponseJobType
	DCL    PreviewJobResultResponseJobType
	IMPORT PreviewJobResultResponseJobType
	EXPORT PreviewJobResultResponseJobType
	QUERY  PreviewJobResultResponseJobType
	INSERT PreviewJobResultResponseJobType
}

func GetPreviewJobResultResponseJobTypeEnum() PreviewJobResultResponseJobTypeEnum {
	return PreviewJobResultResponseJobTypeEnum{
		DDL: PreviewJobResultResponseJobType{
			value: "DDL",
		},
		DCL: PreviewJobResultResponseJobType{
			value: "DCL",
		},
		IMPORT: PreviewJobResultResponseJobType{
			value: "IMPORT",
		},
		EXPORT: PreviewJobResultResponseJobType{
			value: "EXPORT",
		},
		QUERY: PreviewJobResultResponseJobType{
			value: "QUERY",
		},
		INSERT: PreviewJobResultResponseJobType{
			value: "INSERT",
		},
	}
}

func (c PreviewJobResultResponseJobType) Value() string {
	return c.value
}

func (c PreviewJobResultResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewJobResultResponseJobType) UnmarshalJSON(b []byte) error {
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
