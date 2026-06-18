package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateMergeRequestDiscussionInfoResponse Response Object
type UpdateMergeRequestDiscussionInfoResponse struct {

	// **参数解释：** 检视意见id(主评论和回复共用)。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 个人检视意见(不需要解决)。
	IndividualNote *bool `json:"individual_note,omitempty"`

	// **参数解释：** 评论列表(主评+回复)。
	Notes *[]NoteDto `json:"notes,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 意见类型。 **取值范围：** - MergeRequest: 合并请求下提的检视意见。 - Commit: 代码页或提交记录下提的检视意见。
	NoteableType *UpdateMergeRequestDiscussionInfoResponseNoteableType `json:"noteable_type,omitempty"`

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
	Severity *UpdateMergeRequestDiscussionInfoResponseSeverity `json:"severity,omitempty"`

	// **参数解释：** 严重程度中文。 **取值范围：** - 建议 - 一般 - 严重 - 致命
	SeverityCn *string `json:"severity_cn,omitempty"`

	// **参数解释：** 严重程度英文。 **取值范围：** - Suggestion: 建议。 - Minor: 一般。 - major: 严重。 - fatal: 致命。
	SeverityEn *UpdateMergeRequestDiscussionInfoResponseSeverityEn `json:"severity_en,omitempty"`

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

func (o UpdateMergeRequestDiscussionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestDiscussionInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestDiscussionInfoResponse", string(data)}, " ")
}

type UpdateMergeRequestDiscussionInfoResponseNoteableType struct {
	value string
}

type UpdateMergeRequestDiscussionInfoResponseNoteableTypeEnum struct {
	MERGE_REQUEST UpdateMergeRequestDiscussionInfoResponseNoteableType
	COMMIT        UpdateMergeRequestDiscussionInfoResponseNoteableType
}

func GetUpdateMergeRequestDiscussionInfoResponseNoteableTypeEnum() UpdateMergeRequestDiscussionInfoResponseNoteableTypeEnum {
	return UpdateMergeRequestDiscussionInfoResponseNoteableTypeEnum{
		MERGE_REQUEST: UpdateMergeRequestDiscussionInfoResponseNoteableType{
			value: "MergeRequest",
		},
		COMMIT: UpdateMergeRequestDiscussionInfoResponseNoteableType{
			value: "Commit",
		},
	}
}

func (c UpdateMergeRequestDiscussionInfoResponseNoteableType) Value() string {
	return c.value
}

func (c UpdateMergeRequestDiscussionInfoResponseNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMergeRequestDiscussionInfoResponseNoteableType) UnmarshalJSON(b []byte) error {
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

type UpdateMergeRequestDiscussionInfoResponseSeverity struct {
	value string
}

type UpdateMergeRequestDiscussionInfoResponseSeverityEnum struct {
	SUGGESTION UpdateMergeRequestDiscussionInfoResponseSeverity
	MINOR      UpdateMergeRequestDiscussionInfoResponseSeverity
	MAJOR      UpdateMergeRequestDiscussionInfoResponseSeverity
	FATAL      UpdateMergeRequestDiscussionInfoResponseSeverity
}

func GetUpdateMergeRequestDiscussionInfoResponseSeverityEnum() UpdateMergeRequestDiscussionInfoResponseSeverityEnum {
	return UpdateMergeRequestDiscussionInfoResponseSeverityEnum{
		SUGGESTION: UpdateMergeRequestDiscussionInfoResponseSeverity{
			value: "suggestion",
		},
		MINOR: UpdateMergeRequestDiscussionInfoResponseSeverity{
			value: "minor",
		},
		MAJOR: UpdateMergeRequestDiscussionInfoResponseSeverity{
			value: "major",
		},
		FATAL: UpdateMergeRequestDiscussionInfoResponseSeverity{
			value: "fatal",
		},
	}
}

func (c UpdateMergeRequestDiscussionInfoResponseSeverity) Value() string {
	return c.value
}

func (c UpdateMergeRequestDiscussionInfoResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMergeRequestDiscussionInfoResponseSeverity) UnmarshalJSON(b []byte) error {
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

type UpdateMergeRequestDiscussionInfoResponseSeverityEn struct {
	value string
}

type UpdateMergeRequestDiscussionInfoResponseSeverityEnEnum struct {
	SUGGESTION UpdateMergeRequestDiscussionInfoResponseSeverityEn
	MINOR      UpdateMergeRequestDiscussionInfoResponseSeverityEn
	MAJOR      UpdateMergeRequestDiscussionInfoResponseSeverityEn
	FATAL      UpdateMergeRequestDiscussionInfoResponseSeverityEn
}

func GetUpdateMergeRequestDiscussionInfoResponseSeverityEnEnum() UpdateMergeRequestDiscussionInfoResponseSeverityEnEnum {
	return UpdateMergeRequestDiscussionInfoResponseSeverityEnEnum{
		SUGGESTION: UpdateMergeRequestDiscussionInfoResponseSeverityEn{
			value: "Suggestion",
		},
		MINOR: UpdateMergeRequestDiscussionInfoResponseSeverityEn{
			value: "Minor",
		},
		MAJOR: UpdateMergeRequestDiscussionInfoResponseSeverityEn{
			value: "Major",
		},
		FATAL: UpdateMergeRequestDiscussionInfoResponseSeverityEn{
			value: "Fatal",
		},
	}
}

func (c UpdateMergeRequestDiscussionInfoResponseSeverityEn) Value() string {
	return c.value
}

func (c UpdateMergeRequestDiscussionInfoResponseSeverityEn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMergeRequestDiscussionInfoResponseSeverityEn) UnmarshalJSON(b []byte) error {
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
