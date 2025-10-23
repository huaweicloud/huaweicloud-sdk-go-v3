package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupRepositoriesRequest Request Object
type ListGroupRepositoriesRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 仓库名称搜索关键字。 **取值范围：** 不涉及。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段。 **取值范围：** - id，仓库ID。 - name，仓库名称。 - created_at，创建时间。 - updated_at，更新时间。 **约束限制：** 不涉及。 **默认取值：** updated_at。
	OrderBy *ListGroupRepositoriesRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序字段。 **取值范围：** - asc，升序。 - desc，逆序。 **约束限制：** 不涉及。 **默认取值：** desc。
	Sort *ListGroupRepositoriesRequestSort `json:"sort,omitempty"`
}

func (o ListGroupRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListGroupRepositoriesRequest", string(data)}, " ")
}

type ListGroupRepositoriesRequestOrderBy struct {
	value string
}

type ListGroupRepositoriesRequestOrderByEnum struct {
	ID         ListGroupRepositoriesRequestOrderBy
	NAME       ListGroupRepositoriesRequestOrderBy
	CREATED_AT ListGroupRepositoriesRequestOrderBy
	UPDATED_AT ListGroupRepositoriesRequestOrderBy
}

func GetListGroupRepositoriesRequestOrderByEnum() ListGroupRepositoriesRequestOrderByEnum {
	return ListGroupRepositoriesRequestOrderByEnum{
		ID: ListGroupRepositoriesRequestOrderBy{
			value: "id",
		},
		NAME: ListGroupRepositoriesRequestOrderBy{
			value: "name",
		},
		CREATED_AT: ListGroupRepositoriesRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListGroupRepositoriesRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListGroupRepositoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListGroupRepositoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupRepositoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListGroupRepositoriesRequestSort struct {
	value string
}

type ListGroupRepositoriesRequestSortEnum struct {
	ASC  ListGroupRepositoriesRequestSort
	DESC ListGroupRepositoriesRequestSort
}

func GetListGroupRepositoriesRequestSortEnum() ListGroupRepositoriesRequestSortEnum {
	return ListGroupRepositoriesRequestSortEnum{
		ASC: ListGroupRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListGroupRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListGroupRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListGroupRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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
