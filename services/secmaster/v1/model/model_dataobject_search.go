package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataobjectSearch 搜索列表条件
type DataobjectSearch struct {

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段：create_time | update_time
	SortBy *string `json:"sort_by,omitempty"`

	// 排序方式：DESC | ASC
	Order *DataobjectSearchOrder `json:"order,omitempty"`

	// 搜索开始时间，例如：2023-02-20T00:00:00.000Z
	FromDate *string `json:"from_date,omitempty"`

	// 搜索结束时间，例如：2023-02-27T23:59:59.999Z
	ToDate *string `json:"to_date,omitempty"`

	Condition *DataobjectSearchCondition `json:"condition,omitempty"`
}

func (o DataobjectSearch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectSearch struct{}"
	}

	return strings.Join([]string{"DataobjectSearch", string(data)}, " ")
}

type DataobjectSearchOrder struct {
	value string
}

type DataobjectSearchOrderEnum struct {
	DESC DataobjectSearchOrder
	ASC  DataobjectSearchOrder
}

func GetDataobjectSearchOrderEnum() DataobjectSearchOrderEnum {
	return DataobjectSearchOrderEnum{
		DESC: DataobjectSearchOrder{
			value: "DESC",
		},
		ASC: DataobjectSearchOrder{
			value: "ASC",
		},
	}
}

func (c DataobjectSearchOrder) Value() string {
	return c.value
}

func (c DataobjectSearchOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataobjectSearchOrder) UnmarshalJSON(b []byte) error {
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
