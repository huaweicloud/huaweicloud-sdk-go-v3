package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListEnvironmentsRequest struct {
	// 指定个数，明确指定的时候用于分页，取值[0, 100]。不指定的时候表示不分页，最多查询1000条记录。

	Limit *int32 `json:"limit,omitempty"`
	// 指定查询偏移量，默认偏移量为0.

	Offset *int32 `json:"offset,omitempty"`
	// 排序字段，默认按创建时间排序。  排序字段支持枚举值：create_time、name、update_time。

	OrderBy *string `json:"order_by,omitempty"`
	// desc/asc，默认desc。

	Order *ListEnvironmentsRequestOrder `json:"order,omitempty"`
}

func (o ListEnvironmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsRequest", string(data)}, " ")
}

type ListEnvironmentsRequestOrder struct {
	value string
}

type ListEnvironmentsRequestOrderEnum struct {
	DESC ListEnvironmentsRequestOrder
	ASC  ListEnvironmentsRequestOrder
}

func GetListEnvironmentsRequestOrderEnum() ListEnvironmentsRequestOrderEnum {
	return ListEnvironmentsRequestOrderEnum{
		DESC: ListEnvironmentsRequestOrder{
			value: "desc",
		},
		ASC: ListEnvironmentsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListEnvironmentsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnvironmentsRequestOrder) UnmarshalJSON(b []byte) error {
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
