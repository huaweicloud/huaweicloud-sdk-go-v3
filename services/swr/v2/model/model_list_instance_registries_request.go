package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceRegistriesRequest Request Object
type ListInstanceRegistriesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 排序字段，支持created_at、updated_at、name，默认为created_at
	OrderColumn *ListInstanceRegistriesRequestOrderColumn `json:"order_column,omitempty"`

	// 排序方式，支持desc、asc，默认desc; 注意：order_type需要与order_column配合使用
	OrderType *ListInstanceRegistriesRequestOrderType `json:"order_type,omitempty"`

	// 名称，模糊查询
	Name *string `json:"name,omitempty"`

	// 仓库类型
	Type *string `json:"type,omitempty"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceRegistriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRegistriesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceRegistriesRequest", string(data)}, " ")
}

type ListInstanceRegistriesRequestOrderColumn struct {
	value string
}

type ListInstanceRegistriesRequestOrderColumnEnum struct {
	CREATED_AT ListInstanceRegistriesRequestOrderColumn
	UPDATED_AT ListInstanceRegistriesRequestOrderColumn
	NAME       ListInstanceRegistriesRequestOrderColumn
}

func GetListInstanceRegistriesRequestOrderColumnEnum() ListInstanceRegistriesRequestOrderColumnEnum {
	return ListInstanceRegistriesRequestOrderColumnEnum{
		CREATED_AT: ListInstanceRegistriesRequestOrderColumn{
			value: "created_at",
		},
		UPDATED_AT: ListInstanceRegistriesRequestOrderColumn{
			value: "updated_at",
		},
		NAME: ListInstanceRegistriesRequestOrderColumn{
			value: "name",
		},
	}
}

func (c ListInstanceRegistriesRequestOrderColumn) Value() string {
	return c.value
}

func (c ListInstanceRegistriesRequestOrderColumn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceRegistriesRequestOrderColumn) UnmarshalJSON(b []byte) error {
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

type ListInstanceRegistriesRequestOrderType struct {
	value string
}

type ListInstanceRegistriesRequestOrderTypeEnum struct {
	DESC ListInstanceRegistriesRequestOrderType
	ASC  ListInstanceRegistriesRequestOrderType
}

func GetListInstanceRegistriesRequestOrderTypeEnum() ListInstanceRegistriesRequestOrderTypeEnum {
	return ListInstanceRegistriesRequestOrderTypeEnum{
		DESC: ListInstanceRegistriesRequestOrderType{
			value: "desc",
		},
		ASC: ListInstanceRegistriesRequestOrderType{
			value: "asc",
		},
	}
}

func (c ListInstanceRegistriesRequestOrderType) Value() string {
	return c.value
}

func (c ListInstanceRegistriesRequestOrderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceRegistriesRequestOrderType) UnmarshalJSON(b []byte) error {
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
