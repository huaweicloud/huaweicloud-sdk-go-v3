package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceWebhooksRequest Request Object
type ListInstanceWebhooksRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 排序字段，支持created_at、updated_at、name，默认为created_at
	OrderColumn *string `json:"order_column,omitempty"`

	// 排序方式，支持desc、asc，默认desc;order_type需要与order_column配合使用
	OrderType *ListInstanceWebhooksRequestOrderType `json:"order_type,omitempty"`

	// 策略名称，模糊查询
	Name *string `json:"name,omitempty"`

	// 所属命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceWebhooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceWebhooksRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceWebhooksRequest", string(data)}, " ")
}

type ListInstanceWebhooksRequestOrderType struct {
	value string
}

type ListInstanceWebhooksRequestOrderTypeEnum struct {
	DESC ListInstanceWebhooksRequestOrderType
	ASC  ListInstanceWebhooksRequestOrderType
}

func GetListInstanceWebhooksRequestOrderTypeEnum() ListInstanceWebhooksRequestOrderTypeEnum {
	return ListInstanceWebhooksRequestOrderTypeEnum{
		DESC: ListInstanceWebhooksRequestOrderType{
			value: "desc",
		},
		ASC: ListInstanceWebhooksRequestOrderType{
			value: "asc",
		},
	}
}

func (c ListInstanceWebhooksRequestOrderType) Value() string {
	return c.value
}

func (c ListInstanceWebhooksRequestOrderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceWebhooksRequestOrderType) UnmarshalJSON(b []byte) error {
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
