package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMergeRequestDiscussionResponse Response Object
type ShowMergeRequestDiscussionResponse struct {

	// **参数解释：** 检视意见id(主评论和回复共用)。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 个人检视意见(不需要解决)。
	IndividualNote *bool `json:"individual_note,omitempty"`

	// **参数解释：** 评论列表(主评+回复)。
	Notes *[]NoteDto `json:"notes,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 意见类型。 **取值范围：** - MergeRequest: 合并请求下提的检视意见。 - Commit: 代码页或提交记录下提的检视意见。
	NoteableType *ShowMergeRequestDiscussionResponseNoteableType `json:"noteable_type,omitempty"`

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

	// **参数解释：** 严重程度key。 **取值范围：** - suggestion: 建议。 - minor: 一般。 - major: 严重。 - fatal: 致命。
	Severity *ShowMergeRequestDiscussionResponseSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **取值范围：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。 **取值范围：** - Suggestion: 建议。 - Minor: 一般。 - major: 严重。 - fatal: 致命。
	SeverityEn *ShowMergeRequestDiscussionResponseSeverityEn `json:"severity_en,omitempty"`

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

func (o ShowMergeRequestDiscussionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestDiscussionResponse struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestDiscussionResponse", string(data)}, " ")
}

type ShowMergeRequestDiscussionResponseNoteableType struct {
	value string
}

type ShowMergeRequestDiscussionResponseNoteableTypeEnum struct {
	MERGE_REQUEST ShowMergeRequestDiscussionResponseNoteableType
	COMMIT        ShowMergeRequestDiscussionResponseNoteableType
}

func GetShowMergeRequestDiscussionResponseNoteableTypeEnum() ShowMergeRequestDiscussionResponseNoteableTypeEnum {
	return ShowMergeRequestDiscussionResponseNoteableTypeEnum{
		MERGE_REQUEST: ShowMergeRequestDiscussionResponseNoteableType{
			value: "MergeRequest",
		},
		COMMIT: ShowMergeRequestDiscussionResponseNoteableType{
			value: "Commit",
		},
	}
}

func (c ShowMergeRequestDiscussionResponseNoteableType) Value() string {
	return c.value
}

func (c ShowMergeRequestDiscussionResponseNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestDiscussionResponseNoteableType) UnmarshalJSON(b []byte) error {
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

type ShowMergeRequestDiscussionResponseSeverity struct {
	value string
}

type ShowMergeRequestDiscussionResponseSeverityEnum struct {
	SUGGESTION ShowMergeRequestDiscussionResponseSeverity
	MINOR      ShowMergeRequestDiscussionResponseSeverity
	MAJOR      ShowMergeRequestDiscussionResponseSeverity
	FATAL      ShowMergeRequestDiscussionResponseSeverity
}

func GetShowMergeRequestDiscussionResponseSeverityEnum() ShowMergeRequestDiscussionResponseSeverityEnum {
	return ShowMergeRequestDiscussionResponseSeverityEnum{
		SUGGESTION: ShowMergeRequestDiscussionResponseSeverity{
			value: "suggestion",
		},
		MINOR: ShowMergeRequestDiscussionResponseSeverity{
			value: "minor",
		},
		MAJOR: ShowMergeRequestDiscussionResponseSeverity{
			value: "major",
		},
		FATAL: ShowMergeRequestDiscussionResponseSeverity{
			value: "fatal",
		},
	}
}

func (c ShowMergeRequestDiscussionResponseSeverity) Value() string {
	return c.value
}

func (c ShowMergeRequestDiscussionResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestDiscussionResponseSeverity) UnmarshalJSON(b []byte) error {
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

type ShowMergeRequestDiscussionResponseSeverityEn struct {
	value string
}

type ShowMergeRequestDiscussionResponseSeverityEnEnum struct {
	SUGGESTION ShowMergeRequestDiscussionResponseSeverityEn
	MINOR      ShowMergeRequestDiscussionResponseSeverityEn
	MAJOR      ShowMergeRequestDiscussionResponseSeverityEn
	FATAL      ShowMergeRequestDiscussionResponseSeverityEn
}

func GetShowMergeRequestDiscussionResponseSeverityEnEnum() ShowMergeRequestDiscussionResponseSeverityEnEnum {
	return ShowMergeRequestDiscussionResponseSeverityEnEnum{
		SUGGESTION: ShowMergeRequestDiscussionResponseSeverityEn{
			value: "Suggestion",
		},
		MINOR: ShowMergeRequestDiscussionResponseSeverityEn{
			value: "Minor",
		},
		MAJOR: ShowMergeRequestDiscussionResponseSeverityEn{
			value: "Major",
		},
		FATAL: ShowMergeRequestDiscussionResponseSeverityEn{
			value: "Fatal",
		},
	}
}

func (c ShowMergeRequestDiscussionResponseSeverityEn) Value() string {
	return c.value
}

func (c ShowMergeRequestDiscussionResponseSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestDiscussionResponseSeverityEn) UnmarshalJSON(b []byte) error {
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
