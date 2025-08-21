package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupsRequest Request Object
type ListGroupsRequest struct {

	// **参数解释：** 检索内容
	Search *string `json:"search,omitempty"`

	// **参数解释：** 所有可用的代码组。
	AllAvailable *bool `json:"all_available,omitempty"`

	// **参数解释：** 排序字段，name 名称 path 路径 id 唯一标示 created_at 创建时间 updated_at 更新时间
	OrderBy *ListGroupsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序顺序 asc顺序 desc逆序
	Sort *ListGroupsRequestSort `json:"sort,omitempty"`

	// **参数解释：** 是否关注。
	Starred *bool `json:"starred,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：**
	Owned *bool `json:"owned,omitempty"`
}

func (o ListGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupsRequest", string(data)}, " ")
}

type ListGroupsRequestOrderBy struct {
	value string
}

type ListGroupsRequestOrderByEnum struct {
	NAME       ListGroupsRequestOrderBy
	PATH       ListGroupsRequestOrderBy
	ID         ListGroupsRequestOrderBy
	CREATED_AT ListGroupsRequestOrderBy
	UPDATED_AT ListGroupsRequestOrderBy
}

func GetListGroupsRequestOrderByEnum() ListGroupsRequestOrderByEnum {
	return ListGroupsRequestOrderByEnum{
		NAME: ListGroupsRequestOrderBy{
			value: "name",
		},
		PATH: ListGroupsRequestOrderBy{
			value: "path",
		},
		ID: ListGroupsRequestOrderBy{
			value: "id",
		},
		CREATED_AT: ListGroupsRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListGroupsRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListGroupsRequestOrderBy) Value() string {
	return c.value
}

func (c ListGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListGroupsRequestSort struct {
	value string
}

type ListGroupsRequestSortEnum struct {
	ASC  ListGroupsRequestSort
	DESC ListGroupsRequestSort
}

func GetListGroupsRequestSortEnum() ListGroupsRequestSortEnum {
	return ListGroupsRequestSortEnum{
		ASC: ListGroupsRequestSort{
			value: "asc",
		},
		DESC: ListGroupsRequestSort{
			value: "desc",
		},
	}
}

func (c ListGroupsRequestSort) Value() string {
	return c.value
}

func (c ListGroupsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupsRequestSort) UnmarshalJSON(b []byte) error {
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
