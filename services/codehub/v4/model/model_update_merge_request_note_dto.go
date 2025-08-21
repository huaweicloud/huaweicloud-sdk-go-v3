package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateMergeRequestNoteDto 更新合并请求检视意见请求体
type UpdateMergeRequestNoteDto struct {

	// **参数解释：** 评论内容。
	Body *string `json:"body,omitempty"`

	// **参数解释：** 检视意见严重程度。
	Severity *UpdateMergeRequestNoteDtoSeverity `json:"severity,omitempty"`

	// **参数解释：** 检视意见指派人id。
	AssigneeId *string `json:"assignee_id,omitempty"`

	// **参数解释：** 检视意见分类。
	ReviewCategories *string `json:"review_categories,omitempty"`

	// **参数解释：** 检视意见模块。
	ReviewModules *string `json:"review_modules,omitempty"`

	// **参数解释：** 检视意见检视人id。
	ProposerId *string `json:"proposer_id,omitempty"`

	// **参数解释：** 解决或取消解决检视意见(使用时需仅传此参数)。
	Resolved *bool `json:"resolved,omitempty"`
}

func (o UpdateMergeRequestNoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestNoteDto struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestNoteDto", string(data)}, " ")
}

type UpdateMergeRequestNoteDtoSeverity struct {
	value string
}

type UpdateMergeRequestNoteDtoSeverityEnum struct {
	SUGGESTION UpdateMergeRequestNoteDtoSeverity
	MINOR      UpdateMergeRequestNoteDtoSeverity
	MAJOR      UpdateMergeRequestNoteDtoSeverity
	FATAL      UpdateMergeRequestNoteDtoSeverity
}

func GetUpdateMergeRequestNoteDtoSeverityEnum() UpdateMergeRequestNoteDtoSeverityEnum {
	return UpdateMergeRequestNoteDtoSeverityEnum{
		SUGGESTION: UpdateMergeRequestNoteDtoSeverity{
			value: "suggestion",
		},
		MINOR: UpdateMergeRequestNoteDtoSeverity{
			value: "minor",
		},
		MAJOR: UpdateMergeRequestNoteDtoSeverity{
			value: "major",
		},
		FATAL: UpdateMergeRequestNoteDtoSeverity{
			value: "fatal",
		},
	}
}

func (c UpdateMergeRequestNoteDtoSeverity) Value() string {
	return c.value
}

func (c UpdateMergeRequestNoteDtoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMergeRequestNoteDtoSeverity) UnmarshalJSON(b []byte) error {
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
