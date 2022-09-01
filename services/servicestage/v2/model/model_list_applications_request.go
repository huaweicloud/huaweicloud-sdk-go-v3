package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListApplicationsRequest struct {

	// 指定个数，明确指定的时候用于分页，取值[0, 100]。不指定的时候表示不分页，最多查询1000条记录。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 指定查询偏移量，默认偏移量为0.
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 排序字段，默认按创建时间排序。  排序字段支持枚举值：create_time、name、update_time。
	OrderBy *string `json:"order_by,omitempty" xml:"order_by"`

	// desc/asc，默认desc。
	Order *ListApplicationsRequestOrder `json:"order,omitempty" xml:"order"`
}

func (o ListApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationsRequest", string(data)}, " ")
}

type ListApplicationsRequestOrder struct {
	value string
}

type ListApplicationsRequestOrderEnum struct {
	DESC ListApplicationsRequestOrder
	ASC  ListApplicationsRequestOrder
}

func GetListApplicationsRequestOrderEnum() ListApplicationsRequestOrderEnum {
	return ListApplicationsRequestOrderEnum{
		DESC: ListApplicationsRequestOrder{
			value: "desc",
		},
		ASC: ListApplicationsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListApplicationsRequestOrder) Value() string {
	return c.value
}

func (c ListApplicationsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicationsRequestOrder) UnmarshalJSON(b []byte) error {
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
