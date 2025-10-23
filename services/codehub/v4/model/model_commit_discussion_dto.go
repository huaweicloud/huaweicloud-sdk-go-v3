package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommitDiscussionDto commit单个检视意见信息。
type CommitDiscussionDto struct {

	// **参数解释：** 检视意见id(主评论和回复共用)。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 个人检视意见(不需要解决)。
	IndividualNote *bool `json:"individual_note,omitempty"`

	// **参数解释：** 评论列表(主评+回复)。
	Notes *[]CommitNoteDto `json:"notes,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 意见类型。
	NoteableType *CommitDiscussionDtoNoteableType `json:"noteable_type,omitempty"`

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

	// **参数解释：** 是否为修改文件。
	DiffFile *bool `json:"diff_file,omitempty"`

	// **参数解释：** 检视意见所在文件的新增行数量。
	AddedLines *int32 `json:"added_lines,omitempty"`

	// **参数解释：** 检视意见所在文件的删除行数量。
	RemovedLines *int32 `json:"removed_lines,omitempty"`

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
	Severity *CommitDiscussionDtoSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **约束限制：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。
	SeverityEn *CommitDiscussionDtoSeverityEn `json:"severity_en,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`
}

func (o CommitDiscussionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitDiscussionDto struct{}"
	}

	return strings.Join([]string{"CommitDiscussionDto", string(data)}, " ")
}

type CommitDiscussionDtoNoteableType struct {
	value string
}

type CommitDiscussionDtoNoteableTypeEnum struct {
	MERGE_REQUEST CommitDiscussionDtoNoteableType
	COMMIT        CommitDiscussionDtoNoteableType
}

func GetCommitDiscussionDtoNoteableTypeEnum() CommitDiscussionDtoNoteableTypeEnum {
	return CommitDiscussionDtoNoteableTypeEnum{
		MERGE_REQUEST: CommitDiscussionDtoNoteableType{
			value: "MergeRequest",
		},
		COMMIT: CommitDiscussionDtoNoteableType{
			value: "Commit",
		},
	}
}

func (c CommitDiscussionDtoNoteableType) Value() string {
	return c.value
}

func (c CommitDiscussionDtoNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitDiscussionDtoNoteableType) UnmarshalJSON(b []byte) error {
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

type CommitDiscussionDtoSeverity struct {
	value string
}

type CommitDiscussionDtoSeverityEnum struct {
	SUGGESTION CommitDiscussionDtoSeverity
	MINOR      CommitDiscussionDtoSeverity
	MAJOR      CommitDiscussionDtoSeverity
	FATAL      CommitDiscussionDtoSeverity
}

func GetCommitDiscussionDtoSeverityEnum() CommitDiscussionDtoSeverityEnum {
	return CommitDiscussionDtoSeverityEnum{
		SUGGESTION: CommitDiscussionDtoSeverity{
			value: "suggestion",
		},
		MINOR: CommitDiscussionDtoSeverity{
			value: "minor",
		},
		MAJOR: CommitDiscussionDtoSeverity{
			value: "major",
		},
		FATAL: CommitDiscussionDtoSeverity{
			value: "fatal",
		},
	}
}

func (c CommitDiscussionDtoSeverity) Value() string {
	return c.value
}

func (c CommitDiscussionDtoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitDiscussionDtoSeverity) UnmarshalJSON(b []byte) error {
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

type CommitDiscussionDtoSeverityEn struct {
	value string
}

type CommitDiscussionDtoSeverityEnEnum struct {
	SUGGESTION CommitDiscussionDtoSeverityEn
	MINOR      CommitDiscussionDtoSeverityEn
	MAJOR      CommitDiscussionDtoSeverityEn
	FATAL      CommitDiscussionDtoSeverityEn
}

func GetCommitDiscussionDtoSeverityEnEnum() CommitDiscussionDtoSeverityEnEnum {
	return CommitDiscussionDtoSeverityEnEnum{
		SUGGESTION: CommitDiscussionDtoSeverityEn{
			value: "Suggestion",
		},
		MINOR: CommitDiscussionDtoSeverityEn{
			value: "Minor",
		},
		MAJOR: CommitDiscussionDtoSeverityEn{
			value: "Major",
		},
		FATAL: CommitDiscussionDtoSeverityEn{
			value: "Fatal",
		},
	}
}

func (c CommitDiscussionDtoSeverityEn) Value() string {
	return c.value
}

func (c CommitDiscussionDtoSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitDiscussionDtoSeverityEn) UnmarshalJSON(b []byte) error {
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
