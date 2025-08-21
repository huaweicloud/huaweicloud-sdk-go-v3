package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListPersonalRepositoryImportRecordsRequest Request Object
type ListPersonalRepositoryImportRecordsRequest struct {

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 状态 **取值范围：** - finished, 导入成功 - fail, 导入失败 - importing, 导入中
	State *ListPersonalRepositoryImportRecordsRequestState `json:"state,omitempty"`

	// **参数解释：** 导入来源 **取值范围：** - gitee - self_managed_gitlab - gitlab - github - git - svn - coding - bitbucket - gerrit - codeup
	SourceType *ListPersonalRepositoryImportRecordsRequestSourceType `json:"source_type,omitempty"`

	// **参数解释：** 筛选在该时间之后导入的
	CreatedAfter *sdktime.SdkTime `json:"created_after,omitempty"`

	// **参数解释：** 筛选在该时间之前导入的
	CreatedBefore *sdktime.SdkTime `json:"created_before,omitempty"`

	// **参数解释：** 筛选在该时间之后导入完成的
	FinishedAfter *sdktime.SdkTime `json:"finished_after,omitempty"`

	// **参数解释：** 筛选在该时间之前导入完成的
	FinishedBefore *sdktime.SdkTime `json:"finished_before,omitempty"`

	// **参数解释：** 搜索仓库
	Search *string `json:"search,omitempty"`

	// **参数解释：** 排序方式。 **取值范围：** - created_at, 导入时间 - source_repo_name, 源仓库路径 - size, 源仓库容量
	OrderBy *ListPersonalRepositoryImportRecordsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 返回排序 - asc 正序返回 - desc 倒序返回
	Sort *ListPersonalRepositoryImportRecordsRequestSort `json:"sort,omitempty"`
}

func (o ListPersonalRepositoryImportRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonalRepositoryImportRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListPersonalRepositoryImportRecordsRequest", string(data)}, " ")
}

type ListPersonalRepositoryImportRecordsRequestState struct {
	value string
}

type ListPersonalRepositoryImportRecordsRequestStateEnum struct {
	FINISHED  ListPersonalRepositoryImportRecordsRequestState
	FAIL      ListPersonalRepositoryImportRecordsRequestState
	IMPORTING ListPersonalRepositoryImportRecordsRequestState
}

func GetListPersonalRepositoryImportRecordsRequestStateEnum() ListPersonalRepositoryImportRecordsRequestStateEnum {
	return ListPersonalRepositoryImportRecordsRequestStateEnum{
		FINISHED: ListPersonalRepositoryImportRecordsRequestState{
			value: "finished",
		},
		FAIL: ListPersonalRepositoryImportRecordsRequestState{
			value: "fail",
		},
		IMPORTING: ListPersonalRepositoryImportRecordsRequestState{
			value: "importing",
		},
	}
}

func (c ListPersonalRepositoryImportRecordsRequestState) Value() string {
	return c.value
}

func (c ListPersonalRepositoryImportRecordsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalRepositoryImportRecordsRequestState) UnmarshalJSON(b []byte) error {
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

type ListPersonalRepositoryImportRecordsRequestSourceType struct {
	value string
}

type ListPersonalRepositoryImportRecordsRequestSourceTypeEnum struct {
	GITEE               ListPersonalRepositoryImportRecordsRequestSourceType
	SELF_MANAGED_GITLAB ListPersonalRepositoryImportRecordsRequestSourceType
	GITLAB              ListPersonalRepositoryImportRecordsRequestSourceType
	GITHUB              ListPersonalRepositoryImportRecordsRequestSourceType
	GIT                 ListPersonalRepositoryImportRecordsRequestSourceType
	SVN                 ListPersonalRepositoryImportRecordsRequestSourceType
	CODING              ListPersonalRepositoryImportRecordsRequestSourceType
	BITBUCKET           ListPersonalRepositoryImportRecordsRequestSourceType
	GERRIT              ListPersonalRepositoryImportRecordsRequestSourceType
	CODEUP              ListPersonalRepositoryImportRecordsRequestSourceType
}

func GetListPersonalRepositoryImportRecordsRequestSourceTypeEnum() ListPersonalRepositoryImportRecordsRequestSourceTypeEnum {
	return ListPersonalRepositoryImportRecordsRequestSourceTypeEnum{
		GITEE: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "gitee",
		},
		SELF_MANAGED_GITLAB: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "self_managed_gitlab",
		},
		GITLAB: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "gitlab",
		},
		GITHUB: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "github",
		},
		GIT: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "git",
		},
		SVN: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "svn",
		},
		CODING: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "coding",
		},
		BITBUCKET: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "bitbucket",
		},
		GERRIT: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "gerrit",
		},
		CODEUP: ListPersonalRepositoryImportRecordsRequestSourceType{
			value: "codeup",
		},
	}
}

func (c ListPersonalRepositoryImportRecordsRequestSourceType) Value() string {
	return c.value
}

func (c ListPersonalRepositoryImportRecordsRequestSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalRepositoryImportRecordsRequestSourceType) UnmarshalJSON(b []byte) error {
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

type ListPersonalRepositoryImportRecordsRequestOrderBy struct {
	value string
}

type ListPersonalRepositoryImportRecordsRequestOrderByEnum struct {
	CREATED_AT       ListPersonalRepositoryImportRecordsRequestOrderBy
	SOURCE_REPO_NAME ListPersonalRepositoryImportRecordsRequestOrderBy
	SIZE             ListPersonalRepositoryImportRecordsRequestOrderBy
}

func GetListPersonalRepositoryImportRecordsRequestOrderByEnum() ListPersonalRepositoryImportRecordsRequestOrderByEnum {
	return ListPersonalRepositoryImportRecordsRequestOrderByEnum{
		CREATED_AT: ListPersonalRepositoryImportRecordsRequestOrderBy{
			value: "created_at",
		},
		SOURCE_REPO_NAME: ListPersonalRepositoryImportRecordsRequestOrderBy{
			value: "source_repo_name",
		},
		SIZE: ListPersonalRepositoryImportRecordsRequestOrderBy{
			value: "size",
		},
	}
}

func (c ListPersonalRepositoryImportRecordsRequestOrderBy) Value() string {
	return c.value
}

func (c ListPersonalRepositoryImportRecordsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalRepositoryImportRecordsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListPersonalRepositoryImportRecordsRequestSort struct {
	value string
}

type ListPersonalRepositoryImportRecordsRequestSortEnum struct {
	ASC  ListPersonalRepositoryImportRecordsRequestSort
	DESC ListPersonalRepositoryImportRecordsRequestSort
}

func GetListPersonalRepositoryImportRecordsRequestSortEnum() ListPersonalRepositoryImportRecordsRequestSortEnum {
	return ListPersonalRepositoryImportRecordsRequestSortEnum{
		ASC: ListPersonalRepositoryImportRecordsRequestSort{
			value: "asc",
		},
		DESC: ListPersonalRepositoryImportRecordsRequestSort{
			value: "desc",
		},
	}
}

func (c ListPersonalRepositoryImportRecordsRequestSort) Value() string {
	return c.value
}

func (c ListPersonalRepositoryImportRecordsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPersonalRepositoryImportRecordsRequestSort) UnmarshalJSON(b []byte) error {
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
