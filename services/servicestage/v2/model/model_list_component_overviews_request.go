package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListComponentOverviewsRequest Request Object
type ListComponentOverviewsRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`

	// 指定个数，明确指定的时候用于分页，取值[0, 100]。不指定的时候表示不分页，最多查询1000条记录。
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询偏移量，默认偏移量为0.
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段，默认按创建时间排序。  排序字段支持枚举值：create_time、name、update_time。
	OrderBy *string `json:"order_by,omitempty"`

	// desc/asc，默认desc。
	Order *ListComponentOverviewsRequestOrder `json:"order,omitempty"`
}

func (o ListComponentOverviewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentOverviewsRequest struct{}"
	}

	return strings.Join([]string{"ListComponentOverviewsRequest", string(data)}, " ")
}

type ListComponentOverviewsRequestOrder struct {
	value string
}

type ListComponentOverviewsRequestOrderEnum struct {
	DESC ListComponentOverviewsRequestOrder
	ASC  ListComponentOverviewsRequestOrder
}

func GetListComponentOverviewsRequestOrderEnum() ListComponentOverviewsRequestOrderEnum {
	return ListComponentOverviewsRequestOrderEnum{
		DESC: ListComponentOverviewsRequestOrder{
			value: "desc",
		},
		ASC: ListComponentOverviewsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListComponentOverviewsRequestOrder) Value() string {
	return c.value
}

func (c ListComponentOverviewsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentOverviewsRequestOrder) UnmarshalJSON(b []byte) error {
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
