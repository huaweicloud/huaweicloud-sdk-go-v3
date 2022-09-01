package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListSlowLogsRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartDate string `json:"start_date" xml:"start_date"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。只能查询当前时间前一个月内的慢日志。注：结束时间不能晚于当前时间。
	EndDate string `json:"end_date" xml:"end_date"`

	// 节点ID，取空值，表示查询实例下所有允许查询的节点。 使用请参考《DDS API参考》的“查询实例列表和详情”响应消息表“nodes 数据结构说明”的“id”。允许查询的节点如下： - 集群实例下面的 shard节点 - 副本集、单节点实例下面的所有节点
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 语句类型，取空值，表示查询所有语句类型，也可指定如下日志类型： - INSERT - QUERY - UPDATE - REMOVE - GETMORE - COMMAND - KILLCURSORS
	Type *ListSlowLogsRequestType `json:"type,omitempty" xml:"type"`

	// 索引位置，偏移量。取值范围为 [0, 1999]。 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询记录数。取值范围[1, 100]，默认10 （表示默认返回10条数据）。 注意： limit 与 offset 的和需要满足 <= 2000的条件。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListSlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogsRequest", string(data)}, " ")
}

type ListSlowLogsRequestType struct {
	value string
}

type ListSlowLogsRequestTypeEnum struct {
	INSERT      ListSlowLogsRequestType
	QUERY       ListSlowLogsRequestType
	UPDATE      ListSlowLogsRequestType
	REMOVE      ListSlowLogsRequestType
	GETMORE     ListSlowLogsRequestType
	COMMAND     ListSlowLogsRequestType
	KILLCURSORS ListSlowLogsRequestType
}

func GetListSlowLogsRequestTypeEnum() ListSlowLogsRequestTypeEnum {
	return ListSlowLogsRequestTypeEnum{
		INSERT: ListSlowLogsRequestType{
			value: "INSERT",
		},
		QUERY: ListSlowLogsRequestType{
			value: "QUERY",
		},
		UPDATE: ListSlowLogsRequestType{
			value: "UPDATE",
		},
		REMOVE: ListSlowLogsRequestType{
			value: "REMOVE",
		},
		GETMORE: ListSlowLogsRequestType{
			value: "GETMORE",
		},
		COMMAND: ListSlowLogsRequestType{
			value: "COMMAND",
		},
		KILLCURSORS: ListSlowLogsRequestType{
			value: "KILLCURSORS",
		},
	}
}

func (c ListSlowLogsRequestType) Value() string {
	return c.value
}

func (c ListSlowLogsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowLogsRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
