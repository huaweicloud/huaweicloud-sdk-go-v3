package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSharesByTagRequestBody 通过标签查询文件系统列表的请求体
type ListSharesByTagRequestBody struct {

	// 通过标签查询文件系统列表的操作类型。仅支持取值为\"filter\" 或 \"count\"。
	Action ListSharesByTagRequestBodyAction `json:"action"`

	// 设置返回的文件系统个数的最大值。
	Limit *string `json:"limit,omitempty"`

	// 设置返回的文件系统的偏移量
	Offset *string `json:"offset,omitempty"`

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”字段。该字段为false或者未提供该参数时，该条件不生效，即返回所有资源或按\"tags\"，\"matches\"等条件过滤。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含20个key，每个key下面的value最多20个，每个key对应的value可以为空数组但结构体不能缺失。key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]Tag `json:"tags,omitempty"`

	// 搜索字段,key为要匹配的字段，仅支持取值“resource_name”。value为匹配的值，当value以“\\*”结尾时，为前缀搜索。例如：value值为“sfsturbo\\*”时，返回名称为“sfsturbo”开头的所有资源列表。
	Matches *[]ResourceTag `json:"matches,omitempty"`
}

func (o ListSharesByTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesByTagRequestBody struct{}"
	}

	return strings.Join([]string{"ListSharesByTagRequestBody", string(data)}, " ")
}

type ListSharesByTagRequestBodyAction struct {
	value string
}

type ListSharesByTagRequestBodyActionEnum struct {
	FILTER ListSharesByTagRequestBodyAction
	COUNT  ListSharesByTagRequestBodyAction
}

func GetListSharesByTagRequestBodyActionEnum() ListSharesByTagRequestBodyActionEnum {
	return ListSharesByTagRequestBodyActionEnum{
		FILTER: ListSharesByTagRequestBodyAction{
			value: "filter",
		},
		COUNT: ListSharesByTagRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListSharesByTagRequestBodyAction) Value() string {
	return c.value
}

func (c ListSharesByTagRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSharesByTagRequestBodyAction) UnmarshalJSON(b []byte) error {
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
