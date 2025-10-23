package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCurrentUserRepositoriesRequest Request Object
type ListCurrentUserRepositoriesRequest struct {

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段。 **取值范围：** - created_at，创建时间。 - updated_at，更新时间。 **约束限制：** 不涉及。 **默认取值：** created_at。
	OrderBy *ListCurrentUserRepositoriesRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序字段。 **取值范围：** - asc，升序。 - desc，逆序。 **约束限制：** 不涉及。 **默认取值：** desc。
	Sort *ListCurrentUserRepositoriesRequestSort `json:"sort,omitempty"`

	// **参数解释：** 是否归档。 **取值范围：** - true，归档。 - false，未归档。 **约束限制：** 不涉及。 **默认取值：** false。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 仓库搜索。 **取值范围：** 1-512。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 关注的仓库。 **取值范围：** - true，关注。 - false，未关注。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Starred *bool `json:"starred,omitempty"`

	// **参数解释：** 参与的仓库（有仓库成员身份）。 **取值范围：** - true，筛选我参与的仓库。 - false，不筛选。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Membership *bool `json:"membership,omitempty"`

	// **参数解释：** 创建的仓库。 **取值范围：** - true，筛选我创建的仓库。 - false，不筛选。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	UserCreated *bool `json:"user_created,omitempty"`

	// **参数解释：** 包含异常状态的仓库。 **取值范围：** - true，筛选包含异常状态的仓库。 - false，不筛选。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	IncludeAbnormal *bool `json:"include_abnormal,omitempty"`
}

func (o ListCurrentUserRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCurrentUserRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListCurrentUserRepositoriesRequest", string(data)}, " ")
}

type ListCurrentUserRepositoriesRequestOrderBy struct {
	value string
}

type ListCurrentUserRepositoriesRequestOrderByEnum struct {
	CREATED_AT ListCurrentUserRepositoriesRequestOrderBy
	UPDATED_AT ListCurrentUserRepositoriesRequestOrderBy
}

func GetListCurrentUserRepositoriesRequestOrderByEnum() ListCurrentUserRepositoriesRequestOrderByEnum {
	return ListCurrentUserRepositoriesRequestOrderByEnum{
		CREATED_AT: ListCurrentUserRepositoriesRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListCurrentUserRepositoriesRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListCurrentUserRepositoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListCurrentUserRepositoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCurrentUserRepositoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListCurrentUserRepositoriesRequestSort struct {
	value string
}

type ListCurrentUserRepositoriesRequestSortEnum struct {
	ASC  ListCurrentUserRepositoriesRequestSort
	DESC ListCurrentUserRepositoriesRequestSort
}

func GetListCurrentUserRepositoriesRequestSortEnum() ListCurrentUserRepositoriesRequestSortEnum {
	return ListCurrentUserRepositoriesRequestSortEnum{
		ASC: ListCurrentUserRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListCurrentUserRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListCurrentUserRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListCurrentUserRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCurrentUserRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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
