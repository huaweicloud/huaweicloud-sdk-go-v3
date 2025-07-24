package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListIDcsRequest Request Object
type ListIDcsRequest struct {

	// 区域
	Region string `json:"region"`

	// 分页大小，取值范围为[1,2000]，默认2000
	Limit *string `json:"limit,omitempty"`

	// 排序升、降序
	Order *ListIDcsRequestOrder `json:"order,omitempty"`

	// 上一页数据的最后一条记录id
	Marker *string `json:"marker,omitempty"`
}

func (o ListIDcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIDcsRequest struct{}"
	}

	return strings.Join([]string{"ListIDcsRequest", string(data)}, " ")
}

type ListIDcsRequestOrder struct {
	value string
}

type ListIDcsRequestOrderEnum struct {
	ASC  ListIDcsRequestOrder
	DESC ListIDcsRequestOrder
}

func GetListIDcsRequestOrderEnum() ListIDcsRequestOrderEnum {
	return ListIDcsRequestOrderEnum{
		ASC: ListIDcsRequestOrder{
			value: "asc",
		},
		DESC: ListIDcsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListIDcsRequestOrder) Value() string {
	return c.value
}

func (c ListIDcsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListIDcsRequestOrder) UnmarshalJSON(b []byte) error {
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
