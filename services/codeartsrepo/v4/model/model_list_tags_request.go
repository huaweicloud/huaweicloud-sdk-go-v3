package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagsRequest Request Object
type ListTagsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 创建者，用户ID或者IamId。 **取值范围：** 字符串长度不少于1，不超过100000。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：** 排序方式。 **约束限制：** - asc，升序。 - desc，降序。
	Sort *ListTagsRequestSort `json:"sort,omitempty"`

	// **参数解释：** 搜索条件，按标签名称搜索。 **取值范围：** 字符串长度不少于1，不超过2000。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 排序方式。 **取值范围：** - name，名字。 - created，创建时间。 - updated，更新时间。
	OrderBy *ListTagsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 结果集展示内容。 **取值范围：** - detail，返回标签的所有信息。 - basic，返回标签的基本信息。 - simple，返回标签的简单信息。
	View *ListTagsRequestView `json:"view,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

type ListTagsRequestSort struct {
	value string
}

type ListTagsRequestSortEnum struct {
	ASC  ListTagsRequestSort
	DESC ListTagsRequestSort
}

func GetListTagsRequestSortEnum() ListTagsRequestSortEnum {
	return ListTagsRequestSortEnum{
		ASC: ListTagsRequestSort{
			value: "asc",
		},
		DESC: ListTagsRequestSort{
			value: "desc",
		},
	}
}

func (c ListTagsRequestSort) Value() string {
	return c.value
}

func (c ListTagsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestSort) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestOrderBy struct {
	value string
}

type ListTagsRequestOrderByEnum struct {
	NAME    ListTagsRequestOrderBy
	UPDATED ListTagsRequestOrderBy
	CREATED ListTagsRequestOrderBy
}

func GetListTagsRequestOrderByEnum() ListTagsRequestOrderByEnum {
	return ListTagsRequestOrderByEnum{
		NAME: ListTagsRequestOrderBy{
			value: "name",
		},
		UPDATED: ListTagsRequestOrderBy{
			value: "updated",
		},
		CREATED: ListTagsRequestOrderBy{
			value: "created",
		},
	}
}

func (c ListTagsRequestOrderBy) Value() string {
	return c.value
}

func (c ListTagsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestView struct {
	value string
}

type ListTagsRequestViewEnum struct {
	SIMPLE ListTagsRequestView
	BASIC  ListTagsRequestView
	DETAIL ListTagsRequestView
}

func GetListTagsRequestViewEnum() ListTagsRequestViewEnum {
	return ListTagsRequestViewEnum{
		SIMPLE: ListTagsRequestView{
			value: "simple",
		},
		BASIC: ListTagsRequestView{
			value: "basic",
		},
		DETAIL: ListTagsRequestView{
			value: "detail",
		},
	}
}

func (c ListTagsRequestView) Value() string {
	return c.value
}

func (c ListTagsRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestView) UnmarshalJSON(b []byte) error {
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
