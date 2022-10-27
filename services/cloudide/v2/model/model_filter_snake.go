package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FilterSnake struct {

	// 过滤集合
	Criteria *[]CriteriaSnake `json:"criteria,omitempty"`

	// 页码
	PageNumber int64 `json:"page_number"`

	// 分页大小
	PageSize int64 `json:"page_size"`

	// 排序字段. - 1 修改日期 - 2 插件名称 - 3 插件作者名称
	SortBy *FilterSnakeSortBy `json:"sort_by,omitempty"`

	// 排序顺序. - 1 升序 - 2 降序
	SortOrder *FilterSnakeSortOrder `json:"sort_order,omitempty"`
}

func (o FilterSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterSnake struct{}"
	}

	return strings.Join([]string{"FilterSnake", string(data)}, " ")
}

type FilterSnakeSortBy struct {
	value int64
}

type FilterSnakeSortByEnum struct {
	E_1 FilterSnakeSortBy
	E_2 FilterSnakeSortBy
	E_3 FilterSnakeSortBy
}

func GetFilterSnakeSortByEnum() FilterSnakeSortByEnum {
	return FilterSnakeSortByEnum{
		E_1: FilterSnakeSortBy{
			value: 1,
		}, E_2: FilterSnakeSortBy{
			value: 2,
		}, E_3: FilterSnakeSortBy{
			value: 3,
		},
	}
}

func (c FilterSnakeSortBy) Value() int64 {
	return c.value
}

func (c FilterSnakeSortBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FilterSnakeSortBy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}

type FilterSnakeSortOrder struct {
	value int64
}

type FilterSnakeSortOrderEnum struct {
	E_1 FilterSnakeSortOrder
	E_2 FilterSnakeSortOrder
}

func GetFilterSnakeSortOrderEnum() FilterSnakeSortOrderEnum {
	return FilterSnakeSortOrderEnum{
		E_1: FilterSnakeSortOrder{
			value: 1,
		}, E_2: FilterSnakeSortOrder{
			value: 2,
		},
	}
}

func (c FilterSnakeSortOrder) Value() int64 {
	return c.value
}

func (c FilterSnakeSortOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FilterSnakeSortOrder) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}
