package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMergeRequestDiscussionResponse Response Object
type CreateMergeRequestDiscussionResponse struct {

	// **参数解释：** 检视意见id(主评论和回复共用)。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 个人检视意见(不需要解决)。
	IndividualNote *bool `json:"individual_note,omitempty"`

	// **参数解释：** 评论列表(主评+回复)。
	Notes *[]NoteDto `json:"notes,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 意见类型。
	NoteableType *CreateMergeRequestDiscussionResponseNoteableType `json:"noteable_type,omitempty"`

	// **参数解释：** 提交记录id。
	CommitId *string `json:"commit_id,omitempty"`

	// **参数解释：** 仓库路径。
	RepositoryFullPath *string `json:"repository_full_path,omitempty"`

	// **参数解释：** 文件旧权限(默认100644)。
	AMode *string `json:"a_mode,omitempty"`

	// **参数解释：** 文件新权限(默认100644)。
	BMode *string `json:"b_mode,omitempty"`

	// **参数解释：** 是否为删除文件。
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// **参数解释：** 是否为新增文件。
	NewFile *bool `json:"new_file,omitempty"`

	// **参数解释：** 是否已解决。
	Resolved *bool `json:"resolved,omitempty"`

	// **参数解释：** 是否已归档。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 意见分类key。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 意见分类中文。
	ReviewCategoriesCn *string `json:"review_categories_cn,omitempty"`

	// **参数解释：** 意见分类英文。
	ReviewCategoriesEn *string `json:"review_categories_en,omitempty"`

	// **参数解释：** 意见模块。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 严重程度key。
	Severity *CreateMergeRequestDiscussionResponseSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **约束限制：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。
	SeverityEn *CreateMergeRequestDiscussionResponseSeverityEn `json:"severity_en,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`

	MergeRequestVersionParams *MergeRequestVersionParamsDto `json:"merge_request_version_params,omitempty"`

	// **参数解释：** 变更页检视意见的代码片段。
	DiffFile *string `json:"diff_file,omitempty"`

	// **参数解释：** 检视意见所在文件的新增行数量。
	AddedLines *int32 `json:"added_lines,omitempty"`

	// **参数解释：** 检视意见所在文件的删除行数量。
	RemovedLines   *int32 `json:"removed_lines,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateMergeRequestDiscussionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionResponse struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionResponse", string(data)}, " ")
}

type CreateMergeRequestDiscussionResponseNoteableType struct {
	value string
}

type CreateMergeRequestDiscussionResponseNoteableTypeEnum struct {
	MERGE_REQUEST CreateMergeRequestDiscussionResponseNoteableType
	COMMIT        CreateMergeRequestDiscussionResponseNoteableType
}

func GetCreateMergeRequestDiscussionResponseNoteableTypeEnum() CreateMergeRequestDiscussionResponseNoteableTypeEnum {
	return CreateMergeRequestDiscussionResponseNoteableTypeEnum{
		MERGE_REQUEST: CreateMergeRequestDiscussionResponseNoteableType{
			value: "MergeRequest",
		},
		COMMIT: CreateMergeRequestDiscussionResponseNoteableType{
			value: "Commit",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseNoteableType) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseNoteableType) UnmarshalJSON(b []byte) error {
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

type CreateMergeRequestDiscussionResponseSeverity struct {
	value string
}

type CreateMergeRequestDiscussionResponseSeverityEnum struct {
	SUGGESTION CreateMergeRequestDiscussionResponseSeverity
	MINOR      CreateMergeRequestDiscussionResponseSeverity
	MAJOR      CreateMergeRequestDiscussionResponseSeverity
	FATAL      CreateMergeRequestDiscussionResponseSeverity
}

func GetCreateMergeRequestDiscussionResponseSeverityEnum() CreateMergeRequestDiscussionResponseSeverityEnum {
	return CreateMergeRequestDiscussionResponseSeverityEnum{
		SUGGESTION: CreateMergeRequestDiscussionResponseSeverity{
			value: "suggestion",
		},
		MINOR: CreateMergeRequestDiscussionResponseSeverity{
			value: "minor",
		},
		MAJOR: CreateMergeRequestDiscussionResponseSeverity{
			value: "major",
		},
		FATAL: CreateMergeRequestDiscussionResponseSeverity{
			value: "fatal",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseSeverity) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseSeverity) UnmarshalJSON(b []byte) error {
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

type CreateMergeRequestDiscussionResponseSeverityEn struct {
	value string
}

type CreateMergeRequestDiscussionResponseSeverityEnEnum struct {
	SUGGESTION CreateMergeRequestDiscussionResponseSeverityEn
	MINOR      CreateMergeRequestDiscussionResponseSeverityEn
	MAJOR      CreateMergeRequestDiscussionResponseSeverityEn
	FATAL      CreateMergeRequestDiscussionResponseSeverityEn
}

func GetCreateMergeRequestDiscussionResponseSeverityEnEnum() CreateMergeRequestDiscussionResponseSeverityEnEnum {
	return CreateMergeRequestDiscussionResponseSeverityEnEnum{
		SUGGESTION: CreateMergeRequestDiscussionResponseSeverityEn{
			value: "Suggestion",
		},
		MINOR: CreateMergeRequestDiscussionResponseSeverityEn{
			value: "Minor",
		},
		MAJOR: CreateMergeRequestDiscussionResponseSeverityEn{
			value: "Major",
		},
		FATAL: CreateMergeRequestDiscussionResponseSeverityEn{
			value: "Fatal",
		},
	}
}

func (c CreateMergeRequestDiscussionResponseSeverityEn) Value() string {
	return c.value
}

func (c CreateMergeRequestDiscussionResponseSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestDiscussionResponseSeverityEn) UnmarshalJSON(b []byte) error {
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
