package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type ListVolumesByTagsRequestBody struct {
	// 操作标识。  根据标签查询云硬盘实例详情时使用“filter”。

	Action ListVolumesByTagsRequestBodyAction `json:"action"`
	// 查询记录数。最小值1，最大值1000，默认为1000。返回的结果中记录数不超过limit值

	Limit *int32 `json:"limit,omitempty"`
	// 资源本身支持的查询条件。标签列表中的标签key值不允许重复。

	Matches *[]Match `json:"matches,omitempty"`
	// 索引位置。最小值0，默认为0。返回的结果中第一条记录为符合查询条件的第“offset值+1”条记录

	Offset *int32 `json:"offset,omitempty"`
	// 标签的键值对。标签列表中最多包含10个key 。标签列表中的标签key值不允许重复。标签列表中多个key之间是“与”的关系，云硬盘必须满足请求中所有key才会匹配出来。

	Tags []TagsForListVolumes `json:"tags"`
}

func (o ListVolumesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListVolumesByTagsRequestBody", string(data)}, " ")
}

type ListVolumesByTagsRequestBodyAction struct {
	value string
}

type ListVolumesByTagsRequestBodyActionEnum struct {
	FILTER ListVolumesByTagsRequestBodyAction
}

func GetListVolumesByTagsRequestBodyActionEnum() ListVolumesByTagsRequestBodyActionEnum {
	return ListVolumesByTagsRequestBodyActionEnum{
		FILTER: ListVolumesByTagsRequestBodyAction{
			value: "filter",
		},
	}
}

func (c ListVolumesByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVolumesByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
