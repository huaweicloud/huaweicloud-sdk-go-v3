package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListClientsRequest Request Object
type ListClientsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点ID。
	NodeId string `json:"node_id"`

	// 偏移量，表示从此偏移量开始查询， 偏移量大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示条数，最小值为1，最大值为1000，若不设置该参数，则为10.
	Limit *int32 `json:"limit,omitempty"`

	// 按客户端连接地址过滤。
	Addr *string `json:"addr,omitempty"`

	// 排序字段。
	Sort *string `json:"sort,omitempty"`

	// 排序方式
	Order *ListClientsRequestOrder `json:"order,omitempty"`
}

func (o ListClientsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClientsRequest struct{}"
	}

	return strings.Join([]string{"ListClientsRequest", string(data)}, " ")
}

type ListClientsRequestOrder struct {
	value string
}

type ListClientsRequestOrderEnum struct {
	ASC  ListClientsRequestOrder
	DESC ListClientsRequestOrder
}

func GetListClientsRequestOrderEnum() ListClientsRequestOrderEnum {
	return ListClientsRequestOrderEnum{
		ASC: ListClientsRequestOrder{
			value: "asc",
		},
		DESC: ListClientsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListClientsRequestOrder) Value() string {
	return c.value
}

func (c ListClientsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListClientsRequestOrder) UnmarshalJSON(b []byte) error {
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
