package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupSubgroupsAndRepositoriesRequest Request Object
type ListGroupSubgroupsAndRepositoriesRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 检索条件，名称。
	Filter *int32 `json:"filter,omitempty"`

	// **参数解释：** 排序字段 id 唯一标识 name 名称 created_at 创建时间 updated_at 更新时间
	OrderBy *ListGroupSubgroupsAndRepositoriesRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序顺序 asc顺序 desc逆序
	Sort *ListGroupSubgroupsAndRepositoriesRequestSort `json:"sort,omitempty"`

	// **参数解释：** 是否归档
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGroupSubgroupsAndRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupSubgroupsAndRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListGroupSubgroupsAndRepositoriesRequest", string(data)}, " ")
}

type ListGroupSubgroupsAndRepositoriesRequestOrderBy struct {
	value string
}

type ListGroupSubgroupsAndRepositoriesRequestOrderByEnum struct {
	ID         ListGroupSubgroupsAndRepositoriesRequestOrderBy
	NAME       ListGroupSubgroupsAndRepositoriesRequestOrderBy
	CREATED_AT ListGroupSubgroupsAndRepositoriesRequestOrderBy
	UPDATED_AT ListGroupSubgroupsAndRepositoriesRequestOrderBy
}

func GetListGroupSubgroupsAndRepositoriesRequestOrderByEnum() ListGroupSubgroupsAndRepositoriesRequestOrderByEnum {
	return ListGroupSubgroupsAndRepositoriesRequestOrderByEnum{
		ID: ListGroupSubgroupsAndRepositoriesRequestOrderBy{
			value: "id",
		},
		NAME: ListGroupSubgroupsAndRepositoriesRequestOrderBy{
			value: "name",
		},
		CREATED_AT: ListGroupSubgroupsAndRepositoriesRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListGroupSubgroupsAndRepositoriesRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListGroupSubgroupsAndRepositoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListGroupSubgroupsAndRepositoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupSubgroupsAndRepositoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListGroupSubgroupsAndRepositoriesRequestSort struct {
	value string
}

type ListGroupSubgroupsAndRepositoriesRequestSortEnum struct {
	ASC  ListGroupSubgroupsAndRepositoriesRequestSort
	DESC ListGroupSubgroupsAndRepositoriesRequestSort
}

func GetListGroupSubgroupsAndRepositoriesRequestSortEnum() ListGroupSubgroupsAndRepositoriesRequestSortEnum {
	return ListGroupSubgroupsAndRepositoriesRequestSortEnum{
		ASC: ListGroupSubgroupsAndRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListGroupSubgroupsAndRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListGroupSubgroupsAndRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListGroupSubgroupsAndRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupSubgroupsAndRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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
