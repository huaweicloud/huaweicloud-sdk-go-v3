package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListLtsSlowLogsRequestBody struct {

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：开始时间不得早于当前时间30天。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：结束时间不能晚于当前时间。
	EndTime string `json:"end_time"`

	// 表示每次查询的日志条数，最大限制100条。
	Limit int32 `json:"limit"`

	// 日志单行序列号，第一次查询时不需要此参数，下一次查询时需要使用，可从上一次查询的返回信息中获取。 说明：当次查询从line_num的下一条日志开始查询，不包含当前line_num日志。
	LineNum *string `json:"line_num,omitempty"`

	// 语句类型，取空值，表示查询所有语句类型。
	OperateType *ListLtsSlowLogsRequestBodyOperateType `json:"operate_type,omitempty"`

	// 节点ID，取空值，表示查询实例下所有允许查询的节点。 使用请参考《DDS API参考》的“查询实例列表和详情”响应消息表“nodes 数据结构说明”的“id”。允许查询的节点如下： - 集群实例下面的 shard节点 - 副本集、单节点实例下面的所有节点
	NodeId *string `json:"node_id,omitempty"`

	// 根据多个关键字搜索日志全文，表示同时匹配所有关键字。 - 最多支持10个关键字。 - 每个关键字最大长度不超过512个字符。
	Keywords *[]string `json:"keywords,omitempty"`

	// 根据多个数据库表名关键字模糊搜索日志，表示匹配至少一个关键字。 - 最多支持10个关键字。 - 每个关键字最大长度不超过64个字符。
	DatabaseKeywords *[]string `json:"database_keywords,omitempty"`

	// 根据多个数据库表名关键字模糊搜索日志，表示匹配至少一个关键字。 - 最多支持10个关键字。 - 每个关键字最大长度不超过128个字符。
	CollectionKeywords *[]string `json:"collection_keywords,omitempty"`

	// 支持根据最大执行时间范围查找日志。单位：ms
	MaxCostTime *int32 `json:"max_cost_time,omitempty"`

	// 支持根据最小执行时间范围查找日志。单位：ms
	MinCostTime *int32 `json:"min_cost_time,omitempty"`
}

func (o ListLtsSlowLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLtsSlowLogsRequestBody", string(data)}, " ")
}

type ListLtsSlowLogsRequestBodyOperateType struct {
	value string
}

type ListLtsSlowLogsRequestBodyOperateTypeEnum struct {
	INSERT      ListLtsSlowLogsRequestBodyOperateType
	QUERY       ListLtsSlowLogsRequestBodyOperateType
	UPDATE      ListLtsSlowLogsRequestBodyOperateType
	REMOVE      ListLtsSlowLogsRequestBodyOperateType
	GETMORE     ListLtsSlowLogsRequestBodyOperateType
	COMMAND     ListLtsSlowLogsRequestBodyOperateType
	KILLCURSORS ListLtsSlowLogsRequestBodyOperateType
}

func GetListLtsSlowLogsRequestBodyOperateTypeEnum() ListLtsSlowLogsRequestBodyOperateTypeEnum {
	return ListLtsSlowLogsRequestBodyOperateTypeEnum{
		INSERT: ListLtsSlowLogsRequestBodyOperateType{
			value: "insert",
		},
		QUERY: ListLtsSlowLogsRequestBodyOperateType{
			value: "query",
		},
		UPDATE: ListLtsSlowLogsRequestBodyOperateType{
			value: "update",
		},
		REMOVE: ListLtsSlowLogsRequestBodyOperateType{
			value: "remove",
		},
		GETMORE: ListLtsSlowLogsRequestBodyOperateType{
			value: "getmore",
		},
		COMMAND: ListLtsSlowLogsRequestBodyOperateType{
			value: "command",
		},
		KILLCURSORS: ListLtsSlowLogsRequestBodyOperateType{
			value: "killcursors",
		},
	}
}

func (c ListLtsSlowLogsRequestBodyOperateType) Value() string {
	return c.value
}

func (c ListLtsSlowLogsRequestBodyOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsSlowLogsRequestBodyOperateType) UnmarshalJSON(b []byte) error {
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
