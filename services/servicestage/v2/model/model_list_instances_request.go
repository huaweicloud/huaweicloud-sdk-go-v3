package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 指定个数，明确指定的时候用于分页，取值[0, 100]。不指定的时候表示不分页，最多查询1000条记录。

	Limit *int32 `json:"limit,omitempty"`
	// 指定查询偏移量，默认偏移量为0.

	Offset *int32 `json:"offset,omitempty"`
	// 排序字段，默认按创建时间排序。  排序字段支持枚举值：create_time、name、update_time。

	OrderBy *string `json:"order_by,omitempty"`
	// desc/asc，默认desc。

	Order *ListInstancesRequestOrder `json:"order,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestOrder struct {
	value string
}

type ListInstancesRequestOrderEnum struct {
	DESC ListInstancesRequestOrder
	ASC  ListInstancesRequestOrder
}

func GetListInstancesRequestOrderEnum() ListInstancesRequestOrderEnum {
	return ListInstancesRequestOrderEnum{
		DESC: ListInstancesRequestOrder{
			value: "desc",
		},
		ASC: ListInstancesRequestOrder{
			value: "asc",
		},
	}
}

func (c ListInstancesRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestOrder) UnmarshalJSON(b []byte) error {
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
