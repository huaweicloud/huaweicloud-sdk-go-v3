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

	// 根据仓库名称或命名空间名称进行模糊查询
	Name *string `json:"name,omitempty"`

	// 是否返回仓库的标签
	WithTag *bool `json:"with_tag,omitempty"`

	// 根据标签过滤，多个标签以英文逗号隔开，最多支持20个标签同时过滤查询。格式为key1=value1,key2=value2。如果标签的key/value值中有特殊字符=，请先把 = 替换成 %3D 再进行请求。注意：使用curl命令请求，需要把 = 替换成 %253D(即 %3D 的 %25 被编码为 %25)再次请求。
	Tags *string `json:"tags,omitempty"`

	// 支持按多个所属命名空间ID过滤，多个ID之间用英文逗号分隔。注意：namespace_id与namespace_ids不支持同时传参。
	NamespaceIds *string `json:"namespace_ids,omitempty"`
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
