package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListPersonalMergeRequestsRequest Request Object
type ListPersonalMergeRequestsRequest struct {

	// **参数解释：** 合并请求状态 **约束限制：** - all，返回所有状态的合并请求。 - opened，返回开启中的合并请求。 - closed，返回关闭的合并请求。 - locked，返回锁定的合并请求。 - merged，返回已合并的合并请求。
	State *ListPersonalMergeRequestsRequestState `json:"state,omitempty"`

	// **参数解释：** 排序方式 **约束限制：** - created_at，根据创建时间排序。 - updated_at，根据更新时间排序。 - merged_at，根据合并时间排序。
	OrderBy *ListPersonalMergeRequestsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序顺序 **约束限制：**   - asc，正序排序。   - desc，倒序排序。
	Sort *ListPersonalMergeRequestsRequestSort `json:"sort,omitempty"`

	// **参数解释：** 合并请求关联标签
	Labels *string `json:"labels,omitempty"`

	// **参数解释：** 指定时间前创建
	CreatedBefore *sdktime.SdkTime `json:"created_before,omitempty"`

	// **参数解释：** 指定时间后创建
	CreatedAfter *sdktime.SdkTime `json:"created_after,omitempty"`

	// **参数解释：** 指定时间后更新
	UpdatedAfter *sdktime.SdkTime `json:"updated_after,omitempty"`

	// **参数解释：** 指定时间前更新
	UpdatedBefore *sdktime.SdkTime `json:"updated_before,omitempty"`

	// **参数解释：** 结果集属性，根据给定的参数返回不同的结果， simple，返回简单数据，basic返回基本数据。
	View *ListPersonalMergeRequestsRequestView `json:"view,omitempty"`

	// **参数解释：** 合并请求创建人
	AuthorId *string `json:"author_id,omitempty"`

	// **参数解释：**   - created_by_me 我创建的合并请求   - assigned_to_me 待我合并的合并请求   - need_my_review 待我检视的合并请求   - need_my_approve 待我审核的合并请求   - all 所有有权限访问的合并请求
	Scope *ListPersonalMergeRequestsRequestScope `json:"scope,omitempty"`

	// **参数解释：** 合并请求原分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释：** 合并请求目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 根据合并请求标题、描述信息过滤关键字
	Search *string `json:"search,omitempty"`

	// **参数解释：** 合并请求标题是否有WIP关键字
	Wip *ListPersonalMergeRequestsRequestWip `json:"wip,omitempty"`

	// **参数解释：** 合并请求合并人
	MergedBy *string `json:"merged_by,omitempty"`

	// **参数解释：** 指定时间后合入
	MergedAfter *sdktime.SdkTime `json:"merged_after,omitempty"`

	// **参数解释：** 指定时间前合入
	MergedBefore *sdktime.SdkTime `json:"merged_before,omitempty"`

	// **参数解释：** 分页参数偏移量
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 分页参数结果数量限制
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 是否只返回合并请求总数
	OnlyCount *bool `json:"only_count,omitempty"`
}

func (o ListPersonalMergeRequestsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonalMergeRequestsRequest struct{}"
	}

	return strings.Join([]string{"ListPersonalMergeRequestsRequest", string(data)}, " ")
}

type ListPersonalMergeRequestsRequestState struct {
	value string
}

type ListPersonalMergeRequestsRequestStateEnum struct {
	OPENED ListPersonalMergeRequestsRequestState
	CLOSED ListPersonalMergeRequestsRequestState
	LOCKED ListPersonalMergeRequestsRequestState
	MERGED ListPersonalMergeRequestsRequestState
	ALL    ListPersonalMergeRequestsRequestState
}

func GetListPersonalMergeRequestsRequestStateEnum() ListPersonalMergeRequestsRequestStateEnum {
	return ListPersonalMergeRequestsRequestStateEnum{
		OPENED: ListPersonalMergeRequestsRequestState{
			value: "opened",
		},
		CLOSED: ListPersonalMergeRequestsRequestState{
			value: "closed",
		},
		LOCKED: ListPersonalMergeRequestsRequestState{
			value: "locked",
		},
		MERGED: ListPersonalMergeRequestsRequestState{
			value: "merged",
		},
		ALL: ListPersonalMergeRequestsRequestState{
			value: "all",
		},
	}
}

func (c ListPersonalMergeRequestsRequestState) Value() string {
	return c.value
}

func (c ListPersonalMergeRequestsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalMergeRequestsRequestState) UnmarshalJSON(b []byte) error {
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

type ListPersonalMergeRequestsRequestOrderBy struct {
	value string
}

type ListPersonalMergeRequestsRequestOrderByEnum struct {
	CREATED_AT ListPersonalMergeRequestsRequestOrderBy
	UPDATED_AT ListPersonalMergeRequestsRequestOrderBy
	MERGED_AT  ListPersonalMergeRequestsRequestOrderBy
}

func GetListPersonalMergeRequestsRequestOrderByEnum() ListPersonalMergeRequestsRequestOrderByEnum {
	return ListPersonalMergeRequestsRequestOrderByEnum{
		CREATED_AT: ListPersonalMergeRequestsRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListPersonalMergeRequestsRequestOrderBy{
			value: "updated_at",
		},
		MERGED_AT: ListPersonalMergeRequestsRequestOrderBy{
			value: "merged_at",
		},
	}
}

func (c ListPersonalMergeRequestsRequestOrderBy) Value() string {
	return c.value
}

func (c ListPersonalMergeRequestsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalMergeRequestsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListPersonalMergeRequestsRequestSort struct {
	value string
}

type ListPersonalMergeRequestsRequestSortEnum struct {
	ASC  ListPersonalMergeRequestsRequestSort
	DESC ListPersonalMergeRequestsRequestSort
}

func GetListPersonalMergeRequestsRequestSortEnum() ListPersonalMergeRequestsRequestSortEnum {
	return ListPersonalMergeRequestsRequestSortEnum{
		ASC: ListPersonalMergeRequestsRequestSort{
			value: "asc",
		},
		DESC: ListPersonalMergeRequestsRequestSort{
			value: "desc",
		},
	}
}

func (c ListPersonalMergeRequestsRequestSort) Value() string {
	return c.value
}

func (c ListPersonalMergeRequestsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalMergeRequestsRequestSort) UnmarshalJSON(b []byte) error {
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

type ListPersonalMergeRequestsRequestView struct {
	value string
}

type ListPersonalMergeRequestsRequestViewEnum struct {
	SIMPLE ListPersonalMergeRequestsRequestView
	BASIC  ListPersonalMergeRequestsRequestView
}

func GetListPersonalMergeRequestsRequestViewEnum() ListPersonalMergeRequestsRequestViewEnum {
	return ListPersonalMergeRequestsRequestViewEnum{
		SIMPLE: ListPersonalMergeRequestsRequestView{
			value: "simple",
		},
		BASIC: ListPersonalMergeRequestsRequestView{
			value: "basic",
		},
	}
}

func (c ListPersonalMergeRequestsRequestView) Value() string {
	return c.value
}

func (c ListPersonalMergeRequestsRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalMergeRequestsRequestView) UnmarshalJSON(b []byte) error {
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

type ListPersonalMergeRequestsRequestScope struct {
	value string
}

type ListPersonalMergeRequestsRequestScopeEnum struct {
	CREATED_BY_ME   ListPersonalMergeRequestsRequestScope
	ASSIGNED_TO_ME  ListPersonalMergeRequestsRequestScope
	NEED_MY_REVIEW  ListPersonalMergeRequestsRequestScope
	NEED_MY_APPROVE ListPersonalMergeRequestsRequestScope
	ALL             ListPersonalMergeRequestsRequestScope
}

func GetListPersonalMergeRequestsRequestScopeEnum() ListPersonalMergeRequestsRequestScopeEnum {
	return ListPersonalMergeRequestsRequestScopeEnum{
		CREATED_BY_ME: ListPersonalMergeRequestsRequestScope{
			value: "created_by_me",
		},
		ASSIGNED_TO_ME: ListPersonalMergeRequestsRequestScope{
			value: "assigned_to_me",
		},
		NEED_MY_REVIEW: ListPersonalMergeRequestsRequestScope{
			value: "need_my_review",
		},
		NEED_MY_APPROVE: ListPersonalMergeRequestsRequestScope{
			value: "need_my_approve",
		},
		ALL: ListPersonalMergeRequestsRequestScope{
			value: "all",
		},
	}
}

func (c ListPersonalMergeRequestsRequestScope) Value() string {
	return c.value
}

func (c ListPersonalMergeRequestsRequestScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalMergeRequestsRequestScope) UnmarshalJSON(b []byte) error {
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

type ListPersonalMergeRequestsRequestWip struct {
	value string
}

type ListPersonalMergeRequestsRequestWipEnum struct {
	TRUE  ListPersonalMergeRequestsRequestWip
	FALSE ListPersonalMergeRequestsRequestWip
}

func GetListPersonalMergeRequestsRequestWipEnum() ListPersonalMergeRequestsRequestWipEnum {
	return ListPersonalMergeRequestsRequestWipEnum{
		TRUE: ListPersonalMergeRequestsRequestWip{
			value: "true",
		},
		FALSE: ListPersonalMergeRequestsRequestWip{
			value: "false",
		},
	}
}

func (c ListPersonalMergeRequestsRequestWip) Value() string {
	return c.value
}

func (c ListPersonalMergeRequestsRequestWip) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalMergeRequestsRequestWip) UnmarshalJSON(b []byte) error {
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
