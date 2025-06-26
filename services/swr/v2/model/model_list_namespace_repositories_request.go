package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNamespaceRepositoriesRequest Request Object
type ListNamespaceRepositoriesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段，支持created_at、updated_at，默认为created_at
	OrderColumn *ListNamespaceRepositoriesRequestOrderColumn `json:"order_column,omitempty"`

	// 排序方式，支持desc、asc，默认desc;order_column和order_type参数需要配套使用
	OrderType *ListNamespaceRepositoriesRequestOrderType `json:"order_type,omitempty"`

	// 所属命名空间名称
	NamespaceName string `json:"namespace_name"`
}

func (o ListNamespaceRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespaceRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListNamespaceRepositoriesRequest", string(data)}, " ")
}

type ListNamespaceRepositoriesRequestOrderColumn struct {
	value string
}

type ListNamespaceRepositoriesRequestOrderColumnEnum struct {
	CREATED_AT ListNamespaceRepositoriesRequestOrderColumn
	UPDATED_AT ListNamespaceRepositoriesRequestOrderColumn
}

func GetListNamespaceRepositoriesRequestOrderColumnEnum() ListNamespaceRepositoriesRequestOrderColumnEnum {
	return ListNamespaceRepositoriesRequestOrderColumnEnum{
		CREATED_AT: ListNamespaceRepositoriesRequestOrderColumn{
			value: "created_at",
		},
		UPDATED_AT: ListNamespaceRepositoriesRequestOrderColumn{
			value: "updated_at",
		},
	}
}

func (c ListNamespaceRepositoriesRequestOrderColumn) Value() string {
	return c.value
}

func (c ListNamespaceRepositoriesRequestOrderColumn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNamespaceRepositoriesRequestOrderColumn) UnmarshalJSON(b []byte) error {
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

type ListNamespaceRepositoriesRequestOrderType struct {
	value string
}

type ListNamespaceRepositoriesRequestOrderTypeEnum struct {
	DESC ListNamespaceRepositoriesRequestOrderType
	ASC  ListNamespaceRepositoriesRequestOrderType
}

func GetListNamespaceRepositoriesRequestOrderTypeEnum() ListNamespaceRepositoriesRequestOrderTypeEnum {
	return ListNamespaceRepositoriesRequestOrderTypeEnum{
		DESC: ListNamespaceRepositoriesRequestOrderType{
			value: "desc",
		},
		ASC: ListNamespaceRepositoriesRequestOrderType{
			value: "asc",
		},
	}
}

func (c ListNamespaceRepositoriesRequestOrderType) Value() string {
	return c.value
}

func (c ListNamespaceRepositoriesRequestOrderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNamespaceRepositoriesRequestOrderType) UnmarshalJSON(b []byte) error {
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
