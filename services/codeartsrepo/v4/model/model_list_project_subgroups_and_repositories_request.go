package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectSubgroupsAndRepositoriesRequest Request Object
type ListProjectSubgroupsAndRepositoriesRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 检索条件，名称。
	Filter *int32 `json:"filter,omitempty"`

	// **参数解释：** 排序字段 id 唯一标识 name 名称 created_at 创建时间 updated_at 更新时间
	OrderBy *ListProjectSubgroupsAndRepositoriesRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序顺序 asc顺序 desc逆序
	Sort *ListProjectSubgroupsAndRepositoriesRequestSort `json:"sort,omitempty"`

	// **参数解释：** 是否归档
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListProjectSubgroupsAndRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectSubgroupsAndRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListProjectSubgroupsAndRepositoriesRequest", string(data)}, " ")
}

type ListProjectSubgroupsAndRepositoriesRequestOrderBy struct {
	value string
}

type ListProjectSubgroupsAndRepositoriesRequestOrderByEnum struct {
	ID         ListProjectSubgroupsAndRepositoriesRequestOrderBy
	NAME       ListProjectSubgroupsAndRepositoriesRequestOrderBy
	CREATED_AT ListProjectSubgroupsAndRepositoriesRequestOrderBy
	UPDATED_AT ListProjectSubgroupsAndRepositoriesRequestOrderBy
}

func GetListProjectSubgroupsAndRepositoriesRequestOrderByEnum() ListProjectSubgroupsAndRepositoriesRequestOrderByEnum {
	return ListProjectSubgroupsAndRepositoriesRequestOrderByEnum{
		ID: ListProjectSubgroupsAndRepositoriesRequestOrderBy{
			value: "id",
		},
		NAME: ListProjectSubgroupsAndRepositoriesRequestOrderBy{
			value: "name",
		},
		CREATED_AT: ListProjectSubgroupsAndRepositoriesRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListProjectSubgroupsAndRepositoriesRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListProjectSubgroupsAndRepositoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListProjectSubgroupsAndRepositoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectSubgroupsAndRepositoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListProjectSubgroupsAndRepositoriesRequestSort struct {
	value string
}

type ListProjectSubgroupsAndRepositoriesRequestSortEnum struct {
	ASC  ListProjectSubgroupsAndRepositoriesRequestSort
	DESC ListProjectSubgroupsAndRepositoriesRequestSort
}

func GetListProjectSubgroupsAndRepositoriesRequestSortEnum() ListProjectSubgroupsAndRepositoriesRequestSortEnum {
	return ListProjectSubgroupsAndRepositoriesRequestSortEnum{
		ASC: ListProjectSubgroupsAndRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListProjectSubgroupsAndRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListProjectSubgroupsAndRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListProjectSubgroupsAndRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectSubgroupsAndRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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
