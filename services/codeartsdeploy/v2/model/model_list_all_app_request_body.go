package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAllAppRequestBody struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 分页页码， 表示从此页开始查询， page大于等于1
	Page *int32 `json:"page,omitempty"`

	// 每页显示的条目数量，size小于等于100
	Size *int32 `json:"size,omitempty"`

	// 排序字段的名称，当前仅支持name和startTime
	SortName *ListAllAppRequestBodySortName `json:"sort_name,omitempty"`

	// 排序顺序，正序（ASC）或者逆序（DESC)
	SortBy *string `json:"sort_by,omitempty"`
}

func (o ListAllAppRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllAppRequestBody struct{}"
	}

	return strings.Join([]string{"ListAllAppRequestBody", string(data)}, " ")
}

type ListAllAppRequestBodySortName struct {
	value string
}

type ListAllAppRequestBodySortNameEnum struct {
	NAME       ListAllAppRequestBodySortName
	START_TIME ListAllAppRequestBodySortName
}

func GetListAllAppRequestBodySortNameEnum() ListAllAppRequestBodySortNameEnum {
	return ListAllAppRequestBodySortNameEnum{
		NAME: ListAllAppRequestBodySortName{
			value: "name",
		},
		START_TIME: ListAllAppRequestBodySortName{
			value: "startTime",
		},
	}
}

func (c ListAllAppRequestBodySortName) Value() string {
	return c.value
}

func (c ListAllAppRequestBodySortName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllAppRequestBodySortName) UnmarshalJSON(b []byte) error {
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
