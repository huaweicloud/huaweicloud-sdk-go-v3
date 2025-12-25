package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NoteDto 评论详细信息。
type NoteDto struct {

	// **参数解释：** 评论id(主评论和回复不共用)。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 评论类型。 **取值范围：** - DiscussionNote: 需要解决的关联代码行的评论。 - DiffNote: 一般。
	Type *string `json:"type,omitempty"`

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

	// **参数解释：** 意见类型。 **取值范围：** - MergeRequest: 合并请求下提的检视意见。 - Commit: 代码页或提交记录下提的检视意见。
	NoteableType *NoteDtoNoteableType `json:"noteable_type,omitempty"`

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
	Severity *NoteDtoSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **约束限制：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。
	SeverityEn *NoteDtoSeverityEn `json:"severity_en,omitempty"`

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
}

func (o NoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoteDto struct{}"
	}

	return strings.Join([]string{"NoteDto", string(data)}, " ")
}

type NoteDtoNoteableType struct {
	value string
}

type NoteDtoNoteableTypeEnum struct {
	MERGE_REQUEST NoteDtoNoteableType
	COMMIT        NoteDtoNoteableType
}

func GetNoteDtoNoteableTypeEnum() NoteDtoNoteableTypeEnum {
	return NoteDtoNoteableTypeEnum{
		MERGE_REQUEST: NoteDtoNoteableType{
			value: "MergeRequest",
		},
		COMMIT: NoteDtoNoteableType{
			value: "Commit",
		},
	}
}

func (c NoteDtoNoteableType) Value() string {
	return c.value
}

func (c NoteDtoNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoteDtoNoteableType) UnmarshalJSON(b []byte) error {
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

type NoteDtoSeverity struct {
	value string
}

type NoteDtoSeverityEnum struct {
	SUGGESTION NoteDtoSeverity
	MINOR      NoteDtoSeverity
	MAJOR      NoteDtoSeverity
	FATAL      NoteDtoSeverity
}

func GetNoteDtoSeverityEnum() NoteDtoSeverityEnum {
	return NoteDtoSeverityEnum{
		SUGGESTION: NoteDtoSeverity{
			value: "suggestion",
		},
		MINOR: NoteDtoSeverity{
			value: "minor",
		},
		MAJOR: NoteDtoSeverity{
			value: "major",
		},
		FATAL: NoteDtoSeverity{
			value: "fatal",
		},
	}
}

func (c NoteDtoSeverity) Value() string {
	return c.value
}

func (c NoteDtoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoteDtoSeverity) UnmarshalJSON(b []byte) error {
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

type NoteDtoSeverityEn struct {
	value string
}

type NoteDtoSeverityEnEnum struct {
	SUGGESTION NoteDtoSeverityEn
	MINOR      NoteDtoSeverityEn
	MAJOR      NoteDtoSeverityEn
	FATAL      NoteDtoSeverityEn
}

func GetNoteDtoSeverityEnEnum() NoteDtoSeverityEnEnum {
	return NoteDtoSeverityEnEnum{
		SUGGESTION: NoteDtoSeverityEn{
			value: "Suggestion",
		},
		MINOR: NoteDtoSeverityEn{
			value: "Minor",
		},
		MAJOR: NoteDtoSeverityEn{
			value: "Major",
		},
		FATAL: NoteDtoSeverityEn{
			value: "Fatal",
		},
	}
}

func (c NoteDtoSeverityEn) Value() string {
	return c.value
}

func (c NoteDtoSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoteDtoSeverityEn) UnmarshalJSON(b []byte) error {
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
