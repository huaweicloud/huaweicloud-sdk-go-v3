package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectMergeRequestsRequest Request Object
type ListProjectMergeRequestsRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk_ch)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 返回指定状态的合并请求。 **约束限制：** - all，表示所有状态。 - opened，表示开启中状态 - closed，表示已关闭状态 - merged，表示已合并状态
	State *ListProjectMergeRequestsRequestState `json:"state,omitempty"`

	// **参数解释：** 排序方式。 **取值范围：** - created_at，创建时间。 - updated_at，更新时间。
	OrderBy *ListProjectMergeRequestsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序方式。 **约束限制：** - asc，升序。 - desc，降序。
	Sort *ListProjectMergeRequestsRequestSort `json:"sort,omitempty"`

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

	// **参数解释：** 合并请求主题
	Topic *string `json:"topic,omitempty"`
}

func (o ListProjectMergeRequestsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMergeRequestsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectMergeRequestsRequest", string(data)}, " ")
}

type ListProjectMergeRequestsRequestState struct {
	value string
}

type ListProjectMergeRequestsRequestStateEnum struct {
	ALL    ListProjectMergeRequestsRequestState
	OPENED ListProjectMergeRequestsRequestState
	CLOSED ListProjectMergeRequestsRequestState
	MERGED ListProjectMergeRequestsRequestState
}

func GetListProjectMergeRequestsRequestStateEnum() ListProjectMergeRequestsRequestStateEnum {
	return ListProjectMergeRequestsRequestStateEnum{
		ALL: ListProjectMergeRequestsRequestState{
			value: "all",
		},
		OPENED: ListProjectMergeRequestsRequestState{
			value: "opened",
		},
		CLOSED: ListProjectMergeRequestsRequestState{
			value: "closed",
		},
		MERGED: ListProjectMergeRequestsRequestState{
			value: "merged",
		},
	}
}

func (c ListProjectMergeRequestsRequestState) Value() string {
	return c.value
}

func (c ListProjectMergeRequestsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectMergeRequestsRequestState) UnmarshalJSON(b []byte) error {
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

type ListProjectMergeRequestsRequestOrderBy struct {
	value string
}

type ListProjectMergeRequestsRequestOrderByEnum struct {
	CREATED_AT ListProjectMergeRequestsRequestOrderBy
	UPDATED_AT ListProjectMergeRequestsRequestOrderBy
}

func GetListProjectMergeRequestsRequestOrderByEnum() ListProjectMergeRequestsRequestOrderByEnum {
	return ListProjectMergeRequestsRequestOrderByEnum{
		CREATED_AT: ListProjectMergeRequestsRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListProjectMergeRequestsRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListProjectMergeRequestsRequestOrderBy) Value() string {
	return c.value
}

func (c ListProjectMergeRequestsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectMergeRequestsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListProjectMergeRequestsRequestSort struct {
	value string
}

type ListProjectMergeRequestsRequestSortEnum struct {
	ASC  ListProjectMergeRequestsRequestSort
	DESC ListProjectMergeRequestsRequestSort
}

func GetListProjectMergeRequestsRequestSortEnum() ListProjectMergeRequestsRequestSortEnum {
	return ListProjectMergeRequestsRequestSortEnum{
		ASC: ListProjectMergeRequestsRequestSort{
			value: "asc",
		},
		DESC: ListProjectMergeRequestsRequestSort{
			value: "desc",
		},
	}
}

func (c ListProjectMergeRequestsRequestSort) Value() string {
	return c.value
}

func (c ListProjectMergeRequestsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectMergeRequestsRequestSort) UnmarshalJSON(b []byte) error {
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
