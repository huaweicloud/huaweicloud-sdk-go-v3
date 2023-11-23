package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListMongodbErrorLogsRequestBody struct {

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：开始时间不得早于当前时间30天。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：结束时间不能晚于当前时间。
	EndTime string `json:"end_time"`

	// 表示每次查询的日志条数，最大限制100条。
	Limit int32 `json:"limit"`

	// 日志单行序列号，第一次查询时不需要此参数，下一次查询时需要使用，可从上一次查询的返回信息中获取。 说明：当次查询从line_num的下一条日志开始查询，不包含当前line_num日志。
	LineNum *string `json:"line_num,omitempty"`

	// 日志级别，取空值，表示查询所有日志级别的日志。
	Severity *ListMongodbErrorLogsRequestBodySeverity `json:"severity,omitempty"`

	// 节点ID，取空值，表示查询实例下所有允许查询的节点。 具体取值请参考查询实例列表和详情接口\"ListInstances\"中nodes字段数据结构说明的“id”。
	NodeId *string `json:"node_id,omitempty"`

	// 根据多个关键字搜索日志全文，表示同时匹配所有关键字。 - 最多支持10个关键字。 - 每个关键字最大长度不超过512个字符。
	Keywords *[]string `json:"keywords,omitempty"`
}

func (o ListMongodbErrorLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMongodbErrorLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListMongodbErrorLogsRequestBody", string(data)}, " ")
}

type ListMongodbErrorLogsRequestBodySeverity struct {
	value string
}

type ListMongodbErrorLogsRequestBodySeverityEnum struct {
	WARNING ListMongodbErrorLogsRequestBodySeverity
	ERROR   ListMongodbErrorLogsRequestBodySeverity
}

func GetListMongodbErrorLogsRequestBodySeverityEnum() ListMongodbErrorLogsRequestBodySeverityEnum {
	return ListMongodbErrorLogsRequestBodySeverityEnum{
		WARNING: ListMongodbErrorLogsRequestBodySeverity{
			value: "Warning",
		},
		ERROR: ListMongodbErrorLogsRequestBodySeverity{
			value: "Error",
		},
	}
}

func (c ListMongodbErrorLogsRequestBodySeverity) Value() string {
	return c.value
}

func (c ListMongodbErrorLogsRequestBodySeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMongodbErrorLogsRequestBodySeverity) UnmarshalJSON(b []byte) error {
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
