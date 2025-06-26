package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceRepositoriesRequest Request Object
type ListInstanceRepositoriesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段，支持created_at、updated_at，默认为created_at
	OrderColumn *ListInstanceRepositoriesRequestOrderColumn `json:"order_column,omitempty"`

	// 排序方式，支持desc、asc，默认desc
	OrderType *ListInstanceRepositoriesRequestOrderType `json:"order_type,omitempty"`

	// 所属命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`
}

func (o ListInstanceRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceRepositoriesRequest", string(data)}, " ")
}

type ListInstanceRepositoriesRequestOrderColumn struct {
	value string
}

type ListInstanceRepositoriesRequestOrderColumnEnum struct {
	CREATED_AT ListInstanceRepositoriesRequestOrderColumn
	UPDATED_AT ListInstanceRepositoriesRequestOrderColumn
}

func GetListInstanceRepositoriesRequestOrderColumnEnum() ListInstanceRepositoriesRequestOrderColumnEnum {
	return ListInstanceRepositoriesRequestOrderColumnEnum{
		CREATED_AT: ListInstanceRepositoriesRequestOrderColumn{
			value: "created_at",
		},
		UPDATED_AT: ListInstanceRepositoriesRequestOrderColumn{
			value: "updated_at",
		},
	}
}

func (c ListInstanceRepositoriesRequestOrderColumn) Value() string {
	return c.value
}

func (c ListInstanceRepositoriesRequestOrderColumn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceRepositoriesRequestOrderColumn) UnmarshalJSON(b []byte) error {
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

type ListInstanceRepositoriesRequestOrderType struct {
	value string
}

type ListInstanceRepositoriesRequestOrderTypeEnum struct {
	DESC ListInstanceRepositoriesRequestOrderType
	ASC  ListInstanceRepositoriesRequestOrderType
}

func GetListInstanceRepositoriesRequestOrderTypeEnum() ListInstanceRepositoriesRequestOrderTypeEnum {
	return ListInstanceRepositoriesRequestOrderTypeEnum{
		DESC: ListInstanceRepositoriesRequestOrderType{
			value: "desc",
		},
		ASC: ListInstanceRepositoriesRequestOrderType{
			value: "asc",
		},
	}
}

func (c ListInstanceRepositoriesRequestOrderType) Value() string {
	return c.value
}

func (c ListInstanceRepositoriesRequestOrderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceRepositoriesRequestOrderType) UnmarshalJSON(b []byte) error {
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
