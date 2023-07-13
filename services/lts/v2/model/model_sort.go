package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Sort struct {

	// 排序字段
	OrderBy []string `json:"order_by"`

	// 排序顺序
	Order SortOrder `json:"order"`
}

func (o Sort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sort struct{}"
	}

	return strings.Join([]string{"Sort", string(data)}, " ")
}

type SortOrder struct {
	value string
}

type SortOrderEnum struct {
	DESC SortOrder
	ASC  SortOrder
}

func GetSortOrderEnum() SortOrderEnum {
	return SortOrderEnum{
		DESC: SortOrder{
			value: "desc",
		},
		ASC: SortOrder{
			value: "asc",
		},
	}
}

func (c SortOrder) Value() string {
	return c.value
}

func (c SortOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SortOrder) UnmarshalJSON(b []byte) error {
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
