package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListLtsErrorLogsRequestBody struct {

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：开始时间不得早于当前时间30天。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：结束时间不能晚于当前时间。
	EndTime string `json:"end_time"`

	// 表示每次查询的日志条数，最大限制100条。
	Limit int32 `json:"limit"`

	// 日志单行序列号，第一次查询时不需要此参数，下一页查询时需要使用，可从上一次查询的返回信息中获取。 说明：当次查询从line_num的下一条日志开始查询，不包含当前line_num日志。
	LineNum *string `json:"line_num,omitempty"`

	// 日志级别，取空值表示查询所有日志级别。
	Severity *ListLtsErrorLogsRequestBodySeverity `json:"severity,omitempty"`

	// 日志查询方式，需结合line_num使用, 以line_num的日志为起始点: - 取值\"backwards\"时，表示查询上一页limit大小的日志 - 取值\"forwards\"时，表示查询下一页limit大小的日志 - 不传默认\"forwards\"
	SearchType *ListLtsErrorLogsRequestBodySearchType `json:"search_type,omitempty"`

	// 节点ID，取空值，表示查询实例下所有允许查询的节点。 使用请参考《DDS API参考》的“查询实例列表和详情”响应消息表“nodes 数据结构说明”的“id”。允许查询的节点如下： - 集群实例下面的 shard节点 - 副本集、单节点实例下面的所有节点
	NodeId *string `json:"node_id,omitempty"`

	// 根据多个关键字搜索日志全文，表示同时匹配所有关键字。 - 只支持关键字前缀模糊搜索，最多支持10个关键字。 - 每个关键字最大长度不超过512个字符。
	Keywords *[]string `json:"keywords,omitempty"`
}

func (o ListLtsErrorLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsErrorLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLtsErrorLogsRequestBody", string(data)}, " ")
}

type ListLtsErrorLogsRequestBodySeverity struct {
	value string
}

type ListLtsErrorLogsRequestBodySeverityEnum struct {
	WARNING ListLtsErrorLogsRequestBodySeverity
	ERROR   ListLtsErrorLogsRequestBodySeverity
}

func GetListLtsErrorLogsRequestBodySeverityEnum() ListLtsErrorLogsRequestBodySeverityEnum {
	return ListLtsErrorLogsRequestBodySeverityEnum{
		WARNING: ListLtsErrorLogsRequestBodySeverity{
			value: "Warning",
		},
		ERROR: ListLtsErrorLogsRequestBodySeverity{
			value: "Error",
		},
	}
}

func (c ListLtsErrorLogsRequestBodySeverity) Value() string {
	return c.value
}

func (c ListLtsErrorLogsRequestBodySeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsErrorLogsRequestBodySeverity) UnmarshalJSON(b []byte) error {
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

type ListLtsErrorLogsRequestBodySearchType struct {
	value string
}

type ListLtsErrorLogsRequestBodySearchTypeEnum struct {
	BACKWARDS ListLtsErrorLogsRequestBodySearchType
	FORWARDS  ListLtsErrorLogsRequestBodySearchType
}

func GetListLtsErrorLogsRequestBodySearchTypeEnum() ListLtsErrorLogsRequestBodySearchTypeEnum {
	return ListLtsErrorLogsRequestBodySearchTypeEnum{
		BACKWARDS: ListLtsErrorLogsRequestBodySearchType{
			value: "backwards",
		},
		FORWARDS: ListLtsErrorLogsRequestBodySearchType{
			value: "forwards",
		},
	}
}

func (c ListLtsErrorLogsRequestBodySearchType) Value() string {
	return c.value
}

func (c ListLtsErrorLogsRequestBodySearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsErrorLogsRequestBodySearchType) UnmarshalJSON(b []byte) error {
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
