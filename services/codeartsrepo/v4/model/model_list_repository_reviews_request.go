package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListRepositoryReviewsRequest Request Object
type ListRepositoryReviewsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 意见类型。 **取值范围：** - Commit，提交。 - MergeRequest，合并请求。
	NoteableType ListRepositoryReviewsRequestNoteableType `json:"noteable_type"`

	// **参数解释：** 查询评论内容。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 开始日期。
	StartDate *sdktime.SdkTime `json:"start_date,omitempty"`

	// **参数解释：** 结束日期。
	EndDate *sdktime.SdkTime `json:"end_date,omitempty"`

	// **参数解释：** 是否仅返回带有提交计数和diffs计数的结果。 **取值范围：** - true，仅返回带有提交计数和diffs计数的结果。 - false，按照compare_view参数返回结果信息。
	OnlyCount *bool `json:"only_count,omitempty"`

	// **参数解释：** 搜索的检视意见分类。 **取值范围：** 字符串长度不少于1，不超过200。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 搜索的检视意见模块。 **取值范围：** 字符串长度不少于1，不超过200。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 检视意见严重程度
	Severity *ListRepositoryReviewsRequestSeverity `json:"severity,omitempty"`

	// **参数解释：** 检视意见指派人id。
	AssigneeId *int32 `json:"assignee_id,omitempty"`

	// **参数解释：** 检视意见检视人id。
	ProposerId *int32 `json:"proposer_id,omitempty"`

	// **参数解释：** 目标分支名称。 **取值范围：** 字符串长度不少于1，不超过2000。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 是否包含回复。 **取值范围：** - true，包含。 - false，不包含。
	IncludeReply *bool `json:"include_reply,omitempty"`

	// **参数解释：** 排序方式。 **取值范围：** - created，创建时间。 - updated，更新时间。
	OrderBy *ListRepositoryReviewsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 检视意见返回排序 - asc 按创建时间正序返回 - desc 按创建时间倒序返回
	Sort *string `json:"sort,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRepositoryReviewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryReviewsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryReviewsRequest", string(data)}, " ")
}

type ListRepositoryReviewsRequestNoteableType struct {
	value string
}

type ListRepositoryReviewsRequestNoteableTypeEnum struct {
	COMMIT        ListRepositoryReviewsRequestNoteableType
	MERGE_REQUEST ListRepositoryReviewsRequestNoteableType
}

func GetListRepositoryReviewsRequestNoteableTypeEnum() ListRepositoryReviewsRequestNoteableTypeEnum {
	return ListRepositoryReviewsRequestNoteableTypeEnum{
		COMMIT: ListRepositoryReviewsRequestNoteableType{
			value: "Commit",
		},
		MERGE_REQUEST: ListRepositoryReviewsRequestNoteableType{
			value: "MergeRequest",
		},
	}
}

func (c ListRepositoryReviewsRequestNoteableType) Value() string {
	return c.value
}

func (c ListRepositoryReviewsRequestNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryReviewsRequestNoteableType) UnmarshalJSON(b []byte) error {
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

type ListRepositoryReviewsRequestSeverity struct {
	value string
}

type ListRepositoryReviewsRequestSeverityEnum struct {
	SUGGESTION ListRepositoryReviewsRequestSeverity
	MINOR      ListRepositoryReviewsRequestSeverity
	MAJOR      ListRepositoryReviewsRequestSeverity
	FATAL      ListRepositoryReviewsRequestSeverity
}

func GetListRepositoryReviewsRequestSeverityEnum() ListRepositoryReviewsRequestSeverityEnum {
	return ListRepositoryReviewsRequestSeverityEnum{
		SUGGESTION: ListRepositoryReviewsRequestSeverity{
			value: "suggestion",
		},
		MINOR: ListRepositoryReviewsRequestSeverity{
			value: "minor",
		},
		MAJOR: ListRepositoryReviewsRequestSeverity{
			value: "major",
		},
		FATAL: ListRepositoryReviewsRequestSeverity{
			value: "fatal",
		},
	}
}

func (c ListRepositoryReviewsRequestSeverity) Value() string {
	return c.value
}

func (c ListRepositoryReviewsRequestSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryReviewsRequestSeverity) UnmarshalJSON(b []byte) error {
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

type ListRepositoryReviewsRequestOrderBy struct {
	value string
}

type ListRepositoryReviewsRequestOrderByEnum struct {
	CREATED_AT ListRepositoryReviewsRequestOrderBy
	UPDATED_AT ListRepositoryReviewsRequestOrderBy
}

func GetListRepositoryReviewsRequestOrderByEnum() ListRepositoryReviewsRequestOrderByEnum {
	return ListRepositoryReviewsRequestOrderByEnum{
		CREATED_AT: ListRepositoryReviewsRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListRepositoryReviewsRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListRepositoryReviewsRequestOrderBy) Value() string {
	return c.value
}

func (c ListRepositoryReviewsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryReviewsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
