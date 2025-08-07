package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListIRacksRequest Request Object
type ListIRacksRequest struct {

	// 区域
	Region string `json:"region"`

	// 取值为上一页数据的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 分页大小，取值范围为[1,2000]，默认2000
	Limit *string `json:"limit,omitempty"`

	// 排序升、降序
	Order *ListIRacksRequestOrder `json:"order,omitempty"`
}

func (o ListIRacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIRacksRequest struct{}"
	}

	return strings.Join([]string{"ListIRacksRequest", string(data)}, " ")
}

type ListIRacksRequestOrder struct {
	value string
}

type ListIRacksRequestOrderEnum struct {
	ASC  ListIRacksRequestOrder
	DESC ListIRacksRequestOrder
}

func GetListIRacksRequestOrderEnum() ListIRacksRequestOrderEnum {
	return ListIRacksRequestOrderEnum{
		ASC: ListIRacksRequestOrder{
			value: "asc",
		},
		DESC: ListIRacksRequestOrder{
			value: "desc",
		},
	}
}

func (c ListIRacksRequestOrder) Value() string {
	return c.value
}

func (c ListIRacksRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListIRacksRequestOrder) UnmarshalJSON(b []byte) error {
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
