package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryMergeRequestsRequest Request Object
type ListRepositoryMergeRequestsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 返回指定状态的合并请求。 **约束限制：** - all，表示所有状态。 - opened，表示开启中状态 - closed，表示已关闭状态 - merged，表示已合并状态
	State *ListRepositoryMergeRequestsRequestState `json:"state,omitempty"`

	// **参数解释：** 排序方式。 **取值范围：** - created_at，创建时间。 - updated_at，更新时间。
	OrderBy *ListRepositoryMergeRequestsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序方式。 **约束限制：** - asc，升序。 - desc，降序。
	Sort *ListRepositoryMergeRequestsRequestSort `json:"sort,omitempty"`

	// **参数解释：** 返回由指定ID用户创建的合并请求。 多个ID以逗号','分隔，返回满足条件的合并请求并集。
	AuthorId *string `json:"author_id,omitempty"`

	// **参数解释：** 返回指定源分支的合并请求。
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释：** 返回指定目标分支的合并请求。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 合并请求关键字搜索。 返回标题或者描述包含对应关键字的合并请求。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 查询指定源仓库的数据。
	SourceRepositoryId *int32 `json:"source_repository_id,omitempty"`

	// **参数解释：** 是否仅返回合并请求计数。 **取值范围：** - true，仅返回合并请求计数。 - false，返回合并请求详细信息。
	OnlyCount *bool `json:"only_count,omitempty"`

	// **参数解释：** 查询包含指定labels的合并请求。
	Labels *string `json:"labels,omitempty"`
}

func (o ListRepositoryMergeRequestsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryMergeRequestsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryMergeRequestsRequest", string(data)}, " ")
}

type ListRepositoryMergeRequestsRequestState struct {
	value string
}

type ListRepositoryMergeRequestsRequestStateEnum struct {
	ALL    ListRepositoryMergeRequestsRequestState
	OPENED ListRepositoryMergeRequestsRequestState
	CLOSED ListRepositoryMergeRequestsRequestState
	MERGED ListRepositoryMergeRequestsRequestState
}

func GetListRepositoryMergeRequestsRequestStateEnum() ListRepositoryMergeRequestsRequestStateEnum {
	return ListRepositoryMergeRequestsRequestStateEnum{
		ALL: ListRepositoryMergeRequestsRequestState{
			value: "all",
		},
		OPENED: ListRepositoryMergeRequestsRequestState{
			value: "opened",
		},
		CLOSED: ListRepositoryMergeRequestsRequestState{
			value: "closed",
		},
		MERGED: ListRepositoryMergeRequestsRequestState{
			value: "merged",
		},
	}
}

func (c ListRepositoryMergeRequestsRequestState) Value() string {
	return c.value
}

func (c ListRepositoryMergeRequestsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryMergeRequestsRequestState) UnmarshalJSON(b []byte) error {
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

type ListRepositoryMergeRequestsRequestOrderBy struct {
	value string
}

type ListRepositoryMergeRequestsRequestOrderByEnum struct {
	CREATED_AT ListRepositoryMergeRequestsRequestOrderBy
	UPDATED_AT ListRepositoryMergeRequestsRequestOrderBy
}

func GetListRepositoryMergeRequestsRequestOrderByEnum() ListRepositoryMergeRequestsRequestOrderByEnum {
	return ListRepositoryMergeRequestsRequestOrderByEnum{
		CREATED_AT: ListRepositoryMergeRequestsRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListRepositoryMergeRequestsRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListRepositoryMergeRequestsRequestOrderBy) Value() string {
	return c.value
}

func (c ListRepositoryMergeRequestsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryMergeRequestsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListRepositoryMergeRequestsRequestSort struct {
	value string
}

type ListRepositoryMergeRequestsRequestSortEnum struct {
	ASC  ListRepositoryMergeRequestsRequestSort
	DESC ListRepositoryMergeRequestsRequestSort
}

func GetListRepositoryMergeRequestsRequestSortEnum() ListRepositoryMergeRequestsRequestSortEnum {
	return ListRepositoryMergeRequestsRequestSortEnum{
		ASC: ListRepositoryMergeRequestsRequestSort{
			value: "asc",
		},
		DESC: ListRepositoryMergeRequestsRequestSort{
			value: "desc",
		},
	}
}

func (c ListRepositoryMergeRequestsRequestSort) Value() string {
	return c.value
}

func (c ListRepositoryMergeRequestsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryMergeRequestsRequestSort) UnmarshalJSON(b []byte) error {
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
