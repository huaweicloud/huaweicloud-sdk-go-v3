package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 按标签查询专属主机列表请求参数。
type ReqListDehByTags struct {
	// 查询包含所有指定标签的专属主机结构体不能缺失。 1.最多包含10个key，每个key下面的value最多10个。 2.结构体不能缺失。 3.key不能为空或者空字符串。 4.key不能重复。 5.同一个key中value不能重复。

	Tags *[]Tag `json:"tags,omitempty"`
	// 查询不包含所有指定标签的专属主机。 1.最多包含10个key，每个key下面的value最多10个。 2.结构体不能缺失。 3.key不能为空或者空字符串。 4.key不能重复。 5.同一个key中value不能重复。

	NotTags *[]Tag `json:"not_tags,omitempty"`
	// 查询返回的专属主机数量限制，最多为1000，不能为负数。 1.如果action的值为count，此参数无效。 2.如果action的值为filter，limit默认为1000。

	Limit *int32 `json:"limit,omitempty"`
	// 索引位置，从offset指定的下一条数据开始查询。必须为数字，不能为负数。 查询第一页数据时，不需要传入此参数。查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数。 1.如果action的值为count，此参数无效。 2.如果action的值为filter，offset默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 操作标识，包括filter和count两种。 1.filter：表示按标签过滤专属主机，返回符合条件的专属主机列表。此时，为分页查询。 2.count：表示按标签搜索专属主机，返回符合条件的专属主机个数。

	Action ReqListDehByTagsAction `json:"action"`
	// 包含任意标签。 1.最多包含10个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。 2.key不能重复，同一个key中value不能重复。 3.结果返回包含标签的资源列表，key之间是“或”的关系，key-value结构中value是“或”的关系。 4.无过滤条件时返回全量数据。

	TagsAny *[]Tag `json:"tags_any,omitempty"`
	// 不包含任意标签。 1.最多包含10个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。 2.key不能重复，同一个key中value不能重复。 3.结果返回不包含标签的资源列表，key之间是“或”的关系，key-value结构中value是或的关系。 4.无过滤条件时返回全量数据。

	NotTagsAny *[]Tag `json:"not_tags_any,omitempty"`
	// 搜索字段，用于按条件搜索专属主机。  当前仅支持按resource_name进行搜索。

	Matches *[]Match `json:"matches,omitempty"`
}

func (o ReqListDehByTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqListDehByTags struct{}"
	}

	return strings.Join([]string{"ReqListDehByTags", string(data)}, " ")
}

type ReqListDehByTagsAction struct {
	value string
}

type ReqListDehByTagsActionEnum struct {
	FILTER ReqListDehByTagsAction
	COUNT  ReqListDehByTagsAction
}

func GetReqListDehByTagsActionEnum() ReqListDehByTagsActionEnum {
	return ReqListDehByTagsActionEnum{
		FILTER: ReqListDehByTagsAction{
			value: "filter",
		},
		COUNT: ReqListDehByTagsAction{
			value: "count",
		},
	}
}

func (c ReqListDehByTagsAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqListDehByTagsAction) UnmarshalJSON(b []byte) error {
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
