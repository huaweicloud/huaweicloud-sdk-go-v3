package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReviewDto 检视意见详细信息
type ReviewDto struct {
	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// **参数解释：** 评论内容。
	Note *string `json:"note,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 是否为回复。
	IsReply *bool `json:"is_reply,omitempty"`

	ResolvedBy *UserBasicDto `json:"resolved_by,omitempty"`

	// **参数解释：** 检视意见id(主评论和回复共用)。
	DiscussionId *string `json:"discussion_id,omitempty"`

	// **参数解释：** 仓库路径。
	RepositoryPath *string `json:"repository_path,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 意见分类key。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 意见分类中文。
	ReviewCategoriesCn *string `json:"review_categories_cn,omitempty"`

	// **参数解释：** 意见分类英文。
	ReviewCategoriesEn *string `json:"review_categories_en,omitempty"`

	// **参数解释：** 检视意见模块。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 严重程度key。
	Severity *ReviewDtoSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **约束限制：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。
	SeverityEn *ReviewDtoSeverityEn `json:"severity_en,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`

	Reviewer *UserBasicDto `json:"reviewer,omitempty"`

	// **参数解释：** 是否已解决。
	Resolved *bool `json:"resolved,omitempty"`

	// **参数解释：** 解决时间。
	ResolvedAt *string `json:"resolved_at,omitempty"`

	// **参数解释：** 合并请求或commit链接。
	Link *string `json:"link,omitempty"`

	// **参数解释：** 内容审核结果。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 内容审核时间。
	ModerationTime *int64 `json:"moderation_time,omitempty"`

	// **参数解释：** 内容审核状态。
	ModerationStatus *int32 `json:"moderation_status,omitempty"`

	// **参数解释：** 合并请求id(noteable_type=MergRequest时返回)。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 合并请求iid(noteable_type=MergRequest时返回)。
	MergeRequestIid *int32 `json:"merge_request_iid,omitempty"`

	// **参数解释：** 合并请求标题(noteable_type=MergRequest时返回)。
	MergeRequestTitle *string `json:"merge_request_title,omitempty"`

	// **参数解释：** 合并请求状态(noteable_type=MergRequest时返回)。
	MergeRequestState *string `json:"merge_request_state,omitempty"`

	// **参数解释：** commit id(noteable_type=Commit时返回)。
	CommitId *string `json:"commit_id,omitempty"`
}

func (o ReviewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewDto struct{}"
	}

	return strings.Join([]string{"ReviewDto", string(data)}, " ")
}

type ReviewDtoSeverity struct {
	value string
}

type ReviewDtoSeverityEnum struct {
	SUGGESTION ReviewDtoSeverity
	MINOR      ReviewDtoSeverity
	MAJOR      ReviewDtoSeverity
	FATAL      ReviewDtoSeverity
}

func GetReviewDtoSeverityEnum() ReviewDtoSeverityEnum {
	return ReviewDtoSeverityEnum{
		SUGGESTION: ReviewDtoSeverity{
			value: "suggestion",
		},
		MINOR: ReviewDtoSeverity{
			value: "minor",
		},
		MAJOR: ReviewDtoSeverity{
			value: "major",
		},
		FATAL: ReviewDtoSeverity{
			value: "fatal",
		},
	}
}

func (c ReviewDtoSeverity) Value() string {
	return c.value
}

func (c ReviewDtoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReviewDtoSeverity) UnmarshalJSON(b []byte) error {
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

type ReviewDtoSeverityEn struct {
	value string
}

type ReviewDtoSeverityEnEnum struct {
	SUGGESTION ReviewDtoSeverityEn
	MINOR      ReviewDtoSeverityEn
	MAJOR      ReviewDtoSeverityEn
	FATAL      ReviewDtoSeverityEn
}

func GetReviewDtoSeverityEnEnum() ReviewDtoSeverityEnEnum {
	return ReviewDtoSeverityEnEnum{
		SUGGESTION: ReviewDtoSeverityEn{
			value: "Suggestion",
		},
		MINOR: ReviewDtoSeverityEn{
			value: "Minor",
		},
		MAJOR: ReviewDtoSeverityEn{
			value: "Major",
		},
		FATAL: ReviewDtoSeverityEn{
			value: "Fatal",
		},
	}
}

func (c ReviewDtoSeverityEn) Value() string {
	return c.value
}

func (c ReviewDtoSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReviewDtoSeverityEn) UnmarshalJSON(b []byte) error {
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
