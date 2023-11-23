package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListMongodbSlowLogsRequestBody struct {

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：开始时间不得早于当前时间30天。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：结束时间不能晚于当前时间。
	EndTime string `json:"end_time"`

	// 表示每次查询的日志条数，最大限制100条。
	Limit int32 `json:"limit"`

	// 日志单行序列号，第一次查询时不需要此参数，下一次查询时需要使用，可从上一次查询的返回信息中获取。 说明：当次查询从line_num的下一条日志开始查询，不包含当前line_num日志。
	LineNum *string `json:"line_num,omitempty"`

	// 语句类型，取空值，表示查询所有语句类型。
	OperateType *ListMongodbSlowLogsRequestBodyOperateType `json:"operate_type,omitempty"`

	// 节点ID，取空值，表示查询实例下所有允许查询的节点。 具体取值请参考查询实例列表和详情接口\"ListInstances\"中nodes字段数据结构说明的“id”。
	NodeId *string `json:"node_id,omitempty"`

	// 根据多个关键字搜索日志全文，表示同时匹配所有关键字。 - 最多支持10个关键字。 - 每个关键字最大长度不超过512个字符。
	Keywords *[]string `json:"keywords,omitempty"`

	// 根据多个数据库库名关键字模糊搜索日志，表示匹配至少一个关键字。 - 只支持关键字前缀模糊搜索，最多支持10个关键字。 - 每个关键字最大长度不超过64个字符。
	DatabaseKeywords *[]string `json:"database_keywords,omitempty"`

	// 根据多个数据库集合名称关键字模糊搜索日志，表示匹配至少一个关键字。 - 只支持关键字前缀模糊搜索，最多支持10个关键字。 - 每个关键字最大长度不超过128个字符。
	CollectionKeywords *[]string `json:"collection_keywords,omitempty"`

	// 支持根据最大执行时间范围查找日志。单位：ms
	MaxCostTime *int32 `json:"max_cost_time,omitempty"`

	// 支持根据最小执行时间范围查找日志。单位：ms
	MinCostTime *int32 `json:"min_cost_time,omitempty"`
}

func (o ListMongodbSlowLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMongodbSlowLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListMongodbSlowLogsRequestBody", string(data)}, " ")
}

type ListMongodbSlowLogsRequestBodyOperateType struct {
	value string
}

type ListMongodbSlowLogsRequestBodyOperateTypeEnum struct {
	INSERT      ListMongodbSlowLogsRequestBodyOperateType
	QUERY       ListMongodbSlowLogsRequestBodyOperateType
	UPDATE      ListMongodbSlowLogsRequestBodyOperateType
	REMOVE      ListMongodbSlowLogsRequestBodyOperateType
	GETMORE     ListMongodbSlowLogsRequestBodyOperateType
	COMMAND     ListMongodbSlowLogsRequestBodyOperateType
	KILLCURSORS ListMongodbSlowLogsRequestBodyOperateType
}

func GetListMongodbSlowLogsRequestBodyOperateTypeEnum() ListMongodbSlowLogsRequestBodyOperateTypeEnum {
	return ListMongodbSlowLogsRequestBodyOperateTypeEnum{
		INSERT: ListMongodbSlowLogsRequestBodyOperateType{
			value: "insert",
		},
		QUERY: ListMongodbSlowLogsRequestBodyOperateType{
			value: "query",
		},
		UPDATE: ListMongodbSlowLogsRequestBodyOperateType{
			value: "update",
		},
		REMOVE: ListMongodbSlowLogsRequestBodyOperateType{
			value: "remove",
		},
		GETMORE: ListMongodbSlowLogsRequestBodyOperateType{
			value: "getmore",
		},
		COMMAND: ListMongodbSlowLogsRequestBodyOperateType{
			value: "command",
		},
		KILLCURSORS: ListMongodbSlowLogsRequestBodyOperateType{
			value: "killcursors",
		},
	}
}

func (c ListMongodbSlowLogsRequestBodyOperateType) Value() string {
	return c.value
}

func (c ListMongodbSlowLogsRequestBodyOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMongodbSlowLogsRequestBodyOperateType) UnmarshalJSON(b []byte) error {
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
