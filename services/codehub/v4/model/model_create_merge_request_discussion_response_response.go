package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMergeRequestDiscussionResponseResponse Response Object
type CreateMergeRequestDiscussionResponseResponse struct {

	// **参数解释：** 评论id(主评论和回复不共用)。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 类型(普通评论、需要解决的普通评论、需要解决的关联代码行的评论)。
	Type *CreateMergeRequestDiscussionResponseResponseType `json:"type,omitempty"`

	// **参数解释：** 评论内容。
	Body *string `json:"body,omitempty"`

	// **参数解释：** 附件(弃用)。
	Attachment *string `json:"attachment,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 是否为系统添加的。
	System *bool `json:"system,omitempty"`

	// **参数解释：** 合并请求id或issue id。
	NoteableId *int32 `json:"noteable_id,omitempty"`

	// **参数解释：** 意见类型。
	NoteableType *CreateMergeRequestDiscussionResponseResponseNoteableType `json:"noteable_type,omitempty"`

	// **参数解释：** 提交记录id。
	CommitId *string `json:"commit_id,omitempty"`

	// **参数解释：** 是否需要解决。
	Resolvable *bool `json:"resolvable,omitempty"`

	// **参数解释：** 是否为回复。
	IsReply *bool `json:"is_reply,omitempty"`

	ResolvedBy *UserBasicDto `json:"resolved_by,omitempty"`

	// **参数解释：** 合并请求iid或issue iid。
	NoteableIid *int32 `json:"noteable_iid,omitempty"`

	// **参数解释：** 检视意见id(主评论和回复共用)。
	DiscussionId *string `json:"discussion_id,omitempty"`

	// **参数解释：** 仓库路径。
	Repository *string `json:"repository,omitempty"`

	// **参数解释：** 关联代码行所在文件的文件名。
	DiffFile *string `json:"diff_file,omitempty"`

	// **参数解释：** 关联代码行的代码片段。
	Diff *string `json:"diff,omitempty"`

	// **参数解释：** 是否已归档。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 意见分类key。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 意见分类中文名。
	ReviewCategoriesCn *string `json:"review_categories_cn,omitempty"`

	// **参数解释：** 合并请求版本信息。
	ReviewCategoriesEn *string `json:"review_categories_en,omitempty"`

	// **参数解释：** 合并请求版本信息。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 严重程度key。
	Severity *CreateMergeRequestDiscussionResponseResponseSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **约束限制：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。
	SeverityEn *CreateMergeRequestDiscussionResponseResponseSeverityEn `json:"severity_en,omitempty"`

	// **参数解释：** 文件路径(弃用)。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 行号(弃用)。
	Line *string `json:"line,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`

	Position *PositionDto `json:"position,omitempty"`

	// **参数解释：** 是否已解决。
	Resolved *bool `json:"resolved,omitempty"`

	// **参数解释：** 是否已过期。
	IsOutdated *bool `json:"is_outdated,omitempty"`

	// **参数解释：** 内容审核结果。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 内容审核时间。
	ModerationTime *int64 `json:"moderation_time,omitempty"`

	// **参数解释：** 内容审核状态。
	ModerationStatus *int32 `json:"moderation_status,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateMergeRequestDiscussionResponseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionResponseResponse struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionResponseResponse", string(data)}, " ")
}

type CreateMergeRequestDiscussionResponseResponseType struct {
	value string
}

type CreateMergeRequestDiscussionResponseResponseTypeEnum struct {
	DISCUSSION_NOTE CreateMergeRequestDiscussionResponseResponseType
	DIFF_NOTE       CreateMergeRequestDiscussionResponseResponseType
}

func GetCreateMergeRequestDiscussionResponseResponseTypeEnum() CreateMergeRequestDiscussionResponseResponseTypeEnum {
	return CreateMergeRequestDiscussionResponseResponseTypeEnum{
		DISCUSSION_NOTE: CreateMergeRequestDiscussionResponseResponseType{
			value: "DiscussionNote",
		},
		DIFF_NOTE: CreateMergeRequestDiscussionResponseResponseType{
			value: "DiffNote",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseResponseType) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseResponseType) UnmarshalJSON(b []byte) error {
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

type CreateMergeRequestDiscussionResponseResponseNoteableType struct {
	value string
}

type CreateMergeRequestDiscussionResponseResponseNoteableTypeEnum struct {
	MERGE_REQUEST CreateMergeRequestDiscussionResponseResponseNoteableType
	COMMIT        CreateMergeRequestDiscussionResponseResponseNoteableType
}

func GetCreateMergeRequestDiscussionResponseResponseNoteableTypeEnum() CreateMergeRequestDiscussionResponseResponseNoteableTypeEnum {
	return CreateMergeRequestDiscussionResponseResponseNoteableTypeEnum{
		MERGE_REQUEST: CreateMergeRequestDiscussionResponseResponseNoteableType{
			value: "MergeRequest",
		},
		COMMIT: CreateMergeRequestDiscussionResponseResponseNoteableType{
			value: "Commit",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseResponseNoteableType) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseResponseNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseResponseNoteableType) UnmarshalJSON(b []byte) error {
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

type CreateMergeRequestDiscussionResponseResponseSeverity struct {
	value string
}

type CreateMergeRequestDiscussionResponseResponseSeverityEnum struct {
	SUGGESTION CreateMergeRequestDiscussionResponseResponseSeverity
	MINOR      CreateMergeRequestDiscussionResponseResponseSeverity
	MAJOR      CreateMergeRequestDiscussionResponseResponseSeverity
	FATAL      CreateMergeRequestDiscussionResponseResponseSeverity
}

func GetCreateMergeRequestDiscussionResponseResponseSeverityEnum() CreateMergeRequestDiscussionResponseResponseSeverityEnum {
	return CreateMergeRequestDiscussionResponseResponseSeverityEnum{
		SUGGESTION: CreateMergeRequestDiscussionResponseResponseSeverity{
			value: "suggestion",
		},
		MINOR: CreateMergeRequestDiscussionResponseResponseSeverity{
			value: "minor",
		},
		MAJOR: CreateMergeRequestDiscussionResponseResponseSeverity{
			value: "major",
		},
		FATAL: CreateMergeRequestDiscussionResponseResponseSeverity{
			value: "fatal",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseResponseSeverity) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseResponseSeverity) UnmarshalJSON(b []byte) error {
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

type CreateMergeRequestDiscussionResponseResponseSeverityEn struct {
	value string
}

type CreateMergeRequestDiscussionResponseResponseSeverityEnEnum struct {
	SUGGESTION CreateMergeRequestDiscussionResponseResponseSeverityEn
	MINOR      CreateMergeRequestDiscussionResponseResponseSeverityEn
	MAJOR      CreateMergeRequestDiscussionResponseResponseSeverityEn
	FATAL      CreateMergeRequestDiscussionResponseResponseSeverityEn
}

func GetCreateMergeRequestDiscussionResponseResponseSeverityEnEnum() CreateMergeRequestDiscussionResponseResponseSeverityEnEnum {
	return CreateMergeRequestDiscussionResponseResponseSeverityEnEnum{
		SUGGESTION: CreateMergeRequestDiscussionResponseResponseSeverityEn{
			value: "Suggestion",
		},
		MINOR: CreateMergeRequestDiscussionResponseResponseSeverityEn{
			value: "Minor",
		},
		MAJOR: CreateMergeRequestDiscussionResponseResponseSeverityEn{
			value: "Major",
		},
		FATAL: CreateMergeRequestDiscussionResponseResponseSeverityEn{
			value: "Fatal",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseResponseSeverityEn) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseResponseSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseResponseSeverityEn) UnmarshalJSON(b []byte) error {
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
