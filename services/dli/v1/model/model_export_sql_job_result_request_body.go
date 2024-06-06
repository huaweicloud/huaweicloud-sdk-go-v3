package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportSqlJobResultRequestBody 导出查询结果
type ExportSqlJobResultRequestBody struct {

	// ExportResult
	DataPath string `json:"data_path"`

	// 导出数据的压缩格式，目前支持gzip，bzip2和deflate压缩方式； 默认值为none，表示不压缩。
	Compress *string `json:"compress,omitempty"`

	// 导出数据的存储格式，暂时只支持csv格式。
	DataType string `json:"data_type"`

	// 指定执行该任务的队列。若不指定队列，将采用default队列执行操作。
	QueueName *string `json:"queue_name,omitempty"`

	// 导出模式，目前支持“ErrorIfExists”，“Overwrite”，不指定“export_mode”则默认为“ErrorIfExists”。  “ErrorIfExists”：存在即报错。指定的导出目录必须不存在，如果指定目录已经存在，系统将返回错误信息，无法执行导出操作。 “Overwrite”：覆盖。在指定目录下新建文件，会删除已有文件。
	ExportMode *ExportSqlJobResultRequestBodyExportMode `json:"export_mode,omitempty"`

	// 导出csv格式数据时，是否导出列名。  设置为“true”，表示导出列名。 设置为“false”，表示不导出列名。 若为空，默认为“false”。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// 导出数据条数，默认为0表示全部
	LimitNum *int32 `json:"limit_num,omitempty"`

	// 导出数据的编码格式。支持\"utf-8\"，\"gb2312\"，\"gbk\"三种，如果不填写默认为\"utf-8\"。
	EncodingType *string `json:"encoding_type,omitempty"`

	// 用户自定义引用字符
	QuoteChar *string `json:"quote_char,omitempty"`

	// 用户自定义转义字符
	EscapeChar *string `json:"escape_char,omitempty"`
}

func (o ExportSqlJobResultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSqlJobResultRequestBody struct{}"
	}

	return strings.Join([]string{"ExportSqlJobResultRequestBody", string(data)}, " ")
}

type ExportSqlJobResultRequestBodyExportMode struct {
	value string
}

type ExportSqlJobResultRequestBodyExportModeEnum struct {
	ERROR_IF_EXISTS ExportSqlJobResultRequestBodyExportMode
	OVERWRITE       ExportSqlJobResultRequestBodyExportMode
}

func GetExportSqlJobResultRequestBodyExportModeEnum() ExportSqlJobResultRequestBodyExportModeEnum {
	return ExportSqlJobResultRequestBodyExportModeEnum{
		ERROR_IF_EXISTS: ExportSqlJobResultRequestBodyExportMode{
			value: "ErrorIfExists",
		},
		OVERWRITE: ExportSqlJobResultRequestBodyExportMode{
			value: "Overwrite",
		},
	}
}

func (c ExportSqlJobResultRequestBodyExportMode) Value() string {
	return c.value
}

func (c ExportSqlJobResultRequestBodyExportMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSqlJobResultRequestBodyExportMode) UnmarshalJSON(b []byte) error {
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
