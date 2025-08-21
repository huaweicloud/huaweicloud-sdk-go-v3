package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBranchesRequest Request Object
type ListBranchesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支类型，默认返回全部分支，not_protect，返回非保护分支。 **约束限制：** 不涉及。 **取值范围：** 不涉及。
	BranchType *ListBranchesRequestBranchType `json:"branch_type,omitempty"`

	// **参数解释：** 创建者，用户ID或者IamId。 **取值范围：** 字符串长度不少于1，不超过100000。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：** 排序方式 **约束限制：** - name，分支名倒序。 - updated_desc，更新时间倒序。 - updated_asc，更新时间正序。
	Sort *ListBranchesRequestSort `json:"sort,omitempty"`

	// **参数解释：** 分支类型 **约束限制：** - my_branches，个人分支。 - active，活跃分支。 - stale，过时分支。 - all_branches，所有分支。
	QueryView *ListBranchesRequestQueryView `json:"query_view,omitempty"`

	// **参数解释：** 结果集属性，根据给定的参数返回不同的结果 **约束限制：** - all，返回分支的所有信息。 - basic，返回分支的基本信息。 - simple，返回分支的简单信息。
	View *ListBranchesRequestView `json:"view,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListBranchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListBranchesRequest", string(data)}, " ")
}

type ListBranchesRequestBranchType struct {
	value string
}

type ListBranchesRequestBranchTypeEnum struct {
	NOT_PROTECT ListBranchesRequestBranchType
}

func GetListBranchesRequestBranchTypeEnum() ListBranchesRequestBranchTypeEnum {
	return ListBranchesRequestBranchTypeEnum{
		NOT_PROTECT: ListBranchesRequestBranchType{
			value: "not_protect",
		},
	}
}

func (c ListBranchesRequestBranchType) Value() string {
	return c.value
}

func (c ListBranchesRequestBranchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBranchesRequestBranchType) UnmarshalJSON(b []byte) error {
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

type ListBranchesRequestSort struct {
	value string
}

type ListBranchesRequestSortEnum struct {
	NAME         ListBranchesRequestSort
	UPDATED_DESC ListBranchesRequestSort
	UPDATED_ASC  ListBranchesRequestSort
}

func GetListBranchesRequestSortEnum() ListBranchesRequestSortEnum {
	return ListBranchesRequestSortEnum{
		NAME: ListBranchesRequestSort{
			value: "name",
		},
		UPDATED_DESC: ListBranchesRequestSort{
			value: "updated_desc",
		},
		UPDATED_ASC: ListBranchesRequestSort{
			value: "updated_asc",
		},
	}
}

func (c ListBranchesRequestSort) Value() string {
	return c.value
}

func (c ListBranchesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBranchesRequestSort) UnmarshalJSON(b []byte) error {
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

type ListBranchesRequestQueryView struct {
	value string
}

type ListBranchesRequestQueryViewEnum struct {
	MY_BRANCHES  ListBranchesRequestQueryView
	ACTIVE       ListBranchesRequestQueryView
	STALE        ListBranchesRequestQueryView
	ALL_BRANCHES ListBranchesRequestQueryView
}

func GetListBranchesRequestQueryViewEnum() ListBranchesRequestQueryViewEnum {
	return ListBranchesRequestQueryViewEnum{
		MY_BRANCHES: ListBranchesRequestQueryView{
			value: "my_branches",
		},
		ACTIVE: ListBranchesRequestQueryView{
			value: "active",
		},
		STALE: ListBranchesRequestQueryView{
			value: "stale",
		},
		ALL_BRANCHES: ListBranchesRequestQueryView{
			value: "all_branches",
		},
	}
}

func (c ListBranchesRequestQueryView) Value() string {
	return c.value
}

func (c ListBranchesRequestQueryView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBranchesRequestQueryView) UnmarshalJSON(b []byte) error {
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

type ListBranchesRequestView struct {
	value string
}

type ListBranchesRequestViewEnum struct {
	SIMPLE ListBranchesRequestView
	BASIC  ListBranchesRequestView
	ALL    ListBranchesRequestView
}

func GetListBranchesRequestViewEnum() ListBranchesRequestViewEnum {
	return ListBranchesRequestViewEnum{
		SIMPLE: ListBranchesRequestView{
			value: "simple",
		},
		BASIC: ListBranchesRequestView{
			value: "basic",
		},
		ALL: ListBranchesRequestView{
			value: "all",
		},
	}
}

func (c ListBranchesRequestView) Value() string {
	return c.value
}

func (c ListBranchesRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBranchesRequestView) UnmarshalJSON(b []byte) error {
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
