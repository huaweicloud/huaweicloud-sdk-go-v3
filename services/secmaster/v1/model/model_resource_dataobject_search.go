package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceDataobjectSearch 搜索列表条件
type ResourceDataobjectSearch struct {

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段：create_time | update_time
	SortBy *string `json:"sort_by,omitempty"`

	// 排序方式：DESC | ASC
	Order *ResourceDataobjectSearchOrder `json:"order,omitempty"`

	// 搜索开始时间，例如：2023-02-20T00:00:00.000Z
	FromDate *string `json:"from_date,omitempty"`

	// 搜索结束时间，例如：2023-02-27T23:59:59.999Z
	ToDate *string `json:"to_date,omitempty"`

	Condition *DataobjectSearchCondition `json:"condition,omitempty"`
}

func (o ResourceDataobjectSearch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDataobjectSearch struct{}"
	}

	return strings.Join([]string{"ResourceDataobjectSearch", string(data)}, " ")
}

type ResourceDataobjectSearchOrder struct {
	value string
}

type ResourceDataobjectSearchOrderEnum struct {
	DESC ResourceDataobjectSearchOrder
	ASC  ResourceDataobjectSearchOrder
}

func GetResourceDataobjectSearchOrderEnum() ResourceDataobjectSearchOrderEnum {
	return ResourceDataobjectSearchOrderEnum{
		DESC: ResourceDataobjectSearchOrder{
			value: "DESC",
		},
		ASC: ResourceDataobjectSearchOrder{
			value: "ASC",
		},
	}
}

func (c ResourceDataobjectSearchOrder) Value() string {
	return c.value
}

func (c ResourceDataobjectSearchOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceDataobjectSearchOrder) UnmarshalJSON(b []byte) error {
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
