package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListIRackRequest Request Object
type ListIRackRequest struct {

	// 区域
	Region string `json:"region"`

	// 取值为上一页数据的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 分页大小，取值范围为[1,2000]，默认2000
	Limit *string `json:"limit,omitempty"`

	// 排序升、降序
	Order *ListIRackRequestOrder `json:"order,omitempty"`
}

func (o ListIRackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIRackRequest struct{}"
	}

	return strings.Join([]string{"ListIRackRequest", string(data)}, " ")
}

type ListIRackRequestOrder struct {
	value string
}

type ListIRackRequestOrderEnum struct {
	ASC  ListIRackRequestOrder
	DESC ListIRackRequestOrder
}

func GetListIRackRequestOrderEnum() ListIRackRequestOrderEnum {
	return ListIRackRequestOrderEnum{
		ASC: ListIRackRequestOrder{
			value: "asc",
		},
		DESC: ListIRackRequestOrder{
			value: "desc",
		},
	}
}

func (c ListIRackRequestOrder) Value() string {
	return c.value
}

func (c ListIRackRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListIRackRequestOrder) UnmarshalJSON(b []byte) error {
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
