package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceNamespacesRequest Request Object
type ListInstanceNamespacesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`

	// 按列排序，可设置为updated_at（按更新时间排序）
	OrderColumn *ListInstanceNamespacesRequestOrderColumn `json:"order_column,omitempty"`

	// 排序类型，可设置为desc（降序）、asc（升序），与order_column配合使用
	OrderType *ListInstanceNamespacesRequestOrderType `json:"order_type,omitempty"`

	// 命名空间名称，小写字母或数字开头，后面跟小写字母、数字、点、下划线或中划线（其中点、下划线、中划线不能直接连续），小写字母或数字结尾，1-64个字符。
	Name *string `json:"name,omitempty"`

	// 是否公开，非空且非true默认返回false
	Public *string `json:"public,omitempty"`
}

func (o ListInstanceNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceNamespacesRequest", string(data)}, " ")
}

type ListInstanceNamespacesRequestOrderColumn struct {
	value string
}

type ListInstanceNamespacesRequestOrderColumnEnum struct {
	UPDATED_AT ListInstanceNamespacesRequestOrderColumn
}

func GetListInstanceNamespacesRequestOrderColumnEnum() ListInstanceNamespacesRequestOrderColumnEnum {
	return ListInstanceNamespacesRequestOrderColumnEnum{
		UPDATED_AT: ListInstanceNamespacesRequestOrderColumn{
			value: "updated_at",
		},
	}
}

func (c ListInstanceNamespacesRequestOrderColumn) Value() string {
	return c.value
}

func (c ListInstanceNamespacesRequestOrderColumn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceNamespacesRequestOrderColumn) UnmarshalJSON(b []byte) error {
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

type ListInstanceNamespacesRequestOrderType struct {
	value string
}

type ListInstanceNamespacesRequestOrderTypeEnum struct {
	DESC ListInstanceNamespacesRequestOrderType
	ASC  ListInstanceNamespacesRequestOrderType
}

func GetListInstanceNamespacesRequestOrderTypeEnum() ListInstanceNamespacesRequestOrderTypeEnum {
	return ListInstanceNamespacesRequestOrderTypeEnum{
		DESC: ListInstanceNamespacesRequestOrderType{
			value: "desc",
		},
		ASC: ListInstanceNamespacesRequestOrderType{
			value: "asc",
		},
	}
}

func (c ListInstanceNamespacesRequestOrderType) Value() string {
	return c.value
}

func (c ListInstanceNamespacesRequestOrderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceNamespacesRequestOrderType) UnmarshalJSON(b []byte) error {
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
