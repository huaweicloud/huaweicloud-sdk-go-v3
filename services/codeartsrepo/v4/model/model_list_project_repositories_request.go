package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectRepositoriesRequest Request Object
type ListProjectRepositoriesRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 仓库名称搜索关键字。 **取值范围：** 不涉及。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段。 **取值范围：** - id，仓库ID。 - name，仓库名称。 - created_at，创建时间。 - updated_at，更新时间。 **约束限制：** 不涉及。 **默认取值：** updated_at。
	OrderBy *ListProjectRepositoriesRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序字段。 **取值范围：** - asc，升序。 - desc，逆序。 **约束限制：** 不涉及。 **默认取值：** desc。
	Sort *ListProjectRepositoriesRequestSort `json:"sort,omitempty"`
}

func (o ListProjectRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListProjectRepositoriesRequest", string(data)}, " ")
}

type ListProjectRepositoriesRequestOrderBy struct {
	value string
}

type ListProjectRepositoriesRequestOrderByEnum struct {
	ID         ListProjectRepositoriesRequestOrderBy
	NAME       ListProjectRepositoriesRequestOrderBy
	CREATED_AT ListProjectRepositoriesRequestOrderBy
	UPDATED_AT ListProjectRepositoriesRequestOrderBy
}

func GetListProjectRepositoriesRequestOrderByEnum() ListProjectRepositoriesRequestOrderByEnum {
	return ListProjectRepositoriesRequestOrderByEnum{
		ID: ListProjectRepositoriesRequestOrderBy{
			value: "id",
		},
		NAME: ListProjectRepositoriesRequestOrderBy{
			value: "name",
		},
		CREATED_AT: ListProjectRepositoriesRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListProjectRepositoriesRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListProjectRepositoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListProjectRepositoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectRepositoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListProjectRepositoriesRequestSort struct {
	value string
}

type ListProjectRepositoriesRequestSortEnum struct {
	ASC  ListProjectRepositoriesRequestSort
	DESC ListProjectRepositoriesRequestSort
}

func GetListProjectRepositoriesRequestSortEnum() ListProjectRepositoriesRequestSortEnum {
	return ListProjectRepositoriesRequestSortEnum{
		ASC: ListProjectRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListProjectRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListProjectRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListProjectRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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
