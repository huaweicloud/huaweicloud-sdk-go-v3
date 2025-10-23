package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommitNoteDto 评论详细信息。
type CommitNoteDto struct {

	// **参数解释：** 评论id(主评论和回复不共用)。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 意见类型。
	NoteableType *CommitNoteDtoNoteableType `json:"noteable_type,omitempty"`

	// **参数解释：** 提交记录id。
	CommitId *string `json:"commit_id,omitempty"`

	// **参数解释：** 检视意见id(主评论和回复共用)。
	DiscussionId *string `json:"discussion_id,omitempty"`

	// **参数解释：** 类型(普通评论、需要解决的普通评论、需要解决的关联代码行的评论)。
	Type *CommitNoteDtoType `json:"type,omitempty"`

	// **参数解释：** 评论内容。
	Body *string `json:"body,omitempty"`

	// **参数解释：** 关联代码行所在文件的文件名。
	DiffFile *string `json:"diff_file,omitempty"`

	// **参数解释：** 关联代码行的代码片段。
	Diff *string `json:"diff,omitempty"`

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

	Position *PositionDto `json:"position,omitempty"`

	// **参数解释：** 是否需要解决。
	Resolvable *bool `json:"resolvable,omitempty"`

	// **参数解释：** 是否已解决。
	Resolved *bool `json:"resolved,omitempty"`

	ResolvedBy *UserBasicDto `json:"resolved_by,omitempty"`

	// **参数解释：** 是否已归档。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 合并请求iid或issue iid。
	NoteableIid *int32 `json:"noteable_iid,omitempty"`

	// **参数解释：** 意见分类key。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 意见分类中文名。
	ReviewCategoriesCn *string `json:"review_categories_cn,omitempty"`

	// **参数解释：** 合并请求版本信息。
	ReviewCategoriesEn *string `json:"review_categories_en,omitempty"`

	// **参数解释：** 合并请求版本信息。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 严重程度key。
	Severity *CommitNoteDtoSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **约束限制：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。
	SeverityEn *CommitNoteDtoSeverityEn `json:"severity_en,omitempty"`

	// **参数解释：** 文件路径(弃用)。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 行号(弃用)。
	Line *string `json:"line,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`

	// **参数解释：** 是否为回复。
	IsReply *bool `json:"is_reply,omitempty"`

	// **参数解释：** 内容审核结果。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 内容审核时间。
	ModerationTime *int64 `json:"moderation_time,omitempty"`

	// **参数解释：** 内容审核状态。
	ModerationStatus *int32 `json:"moderation_status,omitempty"`
}

func (o CommitNoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitNoteDto struct{}"
	}

	return strings.Join([]string{"CommitNoteDto", string(data)}, " ")
}

type CommitNoteDtoNoteableType struct {
	value string
}

type CommitNoteDtoNoteableTypeEnum struct {
	MERGE_REQUEST CommitNoteDtoNoteableType
	COMMIT        CommitNoteDtoNoteableType
}

func GetCommitNoteDtoNoteableTypeEnum() CommitNoteDtoNoteableTypeEnum {
	return CommitNoteDtoNoteableTypeEnum{
		MERGE_REQUEST: CommitNoteDtoNoteableType{
			value: "MergeRequest",
		},
		COMMIT: CommitNoteDtoNoteableType{
			value: "Commit",
		},
	}
}

func (c CommitNoteDtoNoteableType) Value() string {
	return c.value
}

func (c CommitNoteDtoNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitNoteDtoNoteableType) UnmarshalJSON(b []byte) error {
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

type CommitNoteDtoType struct {
	value string
}

type CommitNoteDtoTypeEnum struct {
	DISCUSSION_NOTE CommitNoteDtoType
	DIFF_NOTE       CommitNoteDtoType
}

func GetCommitNoteDtoTypeEnum() CommitNoteDtoTypeEnum {
	return CommitNoteDtoTypeEnum{
		DISCUSSION_NOTE: CommitNoteDtoType{
			value: "DiscussionNote",
		},
		DIFF_NOTE: CommitNoteDtoType{
			value: "DiffNote",
		},
	}
}

func (c CommitNoteDtoType) Value() string {
	return c.value
}

func (c CommitNoteDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitNoteDtoType) UnmarshalJSON(b []byte) error {
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

type CommitNoteDtoSeverity struct {
	value string
}

type CommitNoteDtoSeverityEnum struct {
	SUGGESTION CommitNoteDtoSeverity
	MINOR      CommitNoteDtoSeverity
	MAJOR      CommitNoteDtoSeverity
	FATAL      CommitNoteDtoSeverity
}

func GetCommitNoteDtoSeverityEnum() CommitNoteDtoSeverityEnum {
	return CommitNoteDtoSeverityEnum{
		SUGGESTION: CommitNoteDtoSeverity{
			value: "suggestion",
		},
		MINOR: CommitNoteDtoSeverity{
			value: "minor",
		},
		MAJOR: CommitNoteDtoSeverity{
			value: "major",
		},
		FATAL: CommitNoteDtoSeverity{
			value: "fatal",
		},
	}
}

func (c CommitNoteDtoSeverity) Value() string {
	return c.value
}

func (c CommitNoteDtoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitNoteDtoSeverity) UnmarshalJSON(b []byte) error {
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

type CommitNoteDtoSeverityEn struct {
	value string
}

type CommitNoteDtoSeverityEnEnum struct {
	SUGGESTION CommitNoteDtoSeverityEn
	MINOR      CommitNoteDtoSeverityEn
	MAJOR      CommitNoteDtoSeverityEn
	FATAL      CommitNoteDtoSeverityEn
}

func GetCommitNoteDtoSeverityEnEnum() CommitNoteDtoSeverityEnEnum {
	return CommitNoteDtoSeverityEnEnum{
		SUGGESTION: CommitNoteDtoSeverityEn{
			value: "Suggestion",
		},
		MINOR: CommitNoteDtoSeverityEn{
			value: "Minor",
		},
		MAJOR: CommitNoteDtoSeverityEn{
			value: "Major",
		},
		FATAL: CommitNoteDtoSeverityEn{
			value: "Fatal",
		},
	}
}

func (c CommitNoteDtoSeverityEn) Value() string {
	return c.value
}

func (c CommitNoteDtoSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitNoteDtoSeverityEn) UnmarshalJSON(b []byte) error {
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
