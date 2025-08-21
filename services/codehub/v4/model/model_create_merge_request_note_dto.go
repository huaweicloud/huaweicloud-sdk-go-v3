package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMergeRequestNoteDto 创建合并请求检视意见请求体
type CreateMergeRequestNoteDto struct {

	// **参数解释：** 评论内容。
	Body *string `json:"body,omitempty"`

	// **参数解释：** 检视意见严重程度。
	Severity *CreateMergeRequestNoteDtoSeverity `json:"severity,omitempty"`

	// **参数解释：** 检视意见指派人id。
	AssigneeId *string `json:"assignee_id,omitempty"`

	// **参数解释：** 检视意见分类。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 检视意见模块。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 检视人。
	ProposerId *string `json:"proposer_id,omitempty"`

	// **参数解释：** 检视意见所在行类型(代码行评论需要)。
	LineTypes *CreateMergeRequestNoteDtoLineTypes `json:"line_types,omitempty"`

	Position *PositionDto `json:"position,omitempty"`
}

func (o CreateMergeRequestNoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestNoteDto struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestNoteDto", string(data)}, " ")
}

type CreateMergeRequestNoteDtoSeverity struct {
	value string
}

type CreateMergeRequestNoteDtoSeverityEnum struct {
	SUGGESTION CreateMergeRequestNoteDtoSeverity
	MINOR      CreateMergeRequestNoteDtoSeverity
	MAJOR      CreateMergeRequestNoteDtoSeverity
	FATAL      CreateMergeRequestNoteDtoSeverity
}

func GetCreateMergeRequestNoteDtoSeverityEnum() CreateMergeRequestNoteDtoSeverityEnum {
	return CreateMergeRequestNoteDtoSeverityEnum{
		SUGGESTION: CreateMergeRequestNoteDtoSeverity{
			value: "suggestion",
		},
		MINOR: CreateMergeRequestNoteDtoSeverity{
			value: "minor",
		},
		MAJOR: CreateMergeRequestNoteDtoSeverity{
			value: "major",
		},
		FATAL: CreateMergeRequestNoteDtoSeverity{
			value: "fatal",
		},
	}
}

func (c CreateMergeRequestNoteDtoSeverity) Value() string {
	return c.value
}

func (c CreateMergeRequestNoteDtoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestNoteDtoSeverity) UnmarshalJSON(b []byte) error {
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

type CreateMergeRequestNoteDtoLineTypes struct {
	value string
}

type CreateMergeRequestNoteDtoLineTypesEnum struct {
	NEW CreateMergeRequestNoteDtoLineTypes
	OLD CreateMergeRequestNoteDtoLineTypes
}

func GetCreateMergeRequestNoteDtoLineTypesEnum() CreateMergeRequestNoteDtoLineTypesEnum {
	return CreateMergeRequestNoteDtoLineTypesEnum{
		NEW: CreateMergeRequestNoteDtoLineTypes{
			value: "new",
		},
		OLD: CreateMergeRequestNoteDtoLineTypes{
			value: "old",
		},
	}
}

func (c CreateMergeRequestNoteDtoLineTypes) Value() string {
	return c.value
}

func (c CreateMergeRequestNoteDtoLineTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMergeRequestNoteDtoLineTypes) UnmarshalJSON(b []byte) error {
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
