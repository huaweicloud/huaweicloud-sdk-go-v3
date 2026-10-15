package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AutoscalingHistory **参数解释**：单条调优历史记录，记录一次扩容或缩容操作。
type AutoscalingHistory struct {

	// **参数解释**：调优记录唯一ID。 **取值范围**：UUID格式，长度为36个字符。
	TuningId string `json:"tuning_id"`

	// **参数解释**：Flink作业ID。 **取值范围**：UUIDv7格式，长度为36个字符。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：调优动作类型。 **取值范围**：   - SCALE_OUT：扩容（并行度增加）。   - SCALE_IN：缩容（并行度减少）。
	ActionType AutoscalingHistoryActionType `json:"action_type"`

	// **参数解释**：调优消息，国际化字符串，人类可读，不建议程序化解析。 **取值范围**：不涉及。
	TuningMessage *string `json:"tuning_message,omitempty"`

	OldResourceSetting *AutoscalingHistoryOldResourceSetting `json:"old_resource_setting,omitempty"`

	NewResourceSetting *AutoscalingHistoryNewResourceSetting `json:"new_resource_setting,omitempty"`

	// **参数解释**：调优触发时间，格式为ISO 8601标准的UTC日期时间字符串。 **取值范围**：不涉及。
	TriggerTime *sdktime.SdkTime `json:"trigger_time"`
}

func (o AutoscalingHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingHistory struct{}"
	}

	return strings.Join([]string{"AutoscalingHistory", string(data)}, " ")
}

type AutoscalingHistoryActionType struct {
	value string
}

type AutoscalingHistoryActionTypeEnum struct {
	SCALE_OUT AutoscalingHistoryActionType
	SCALE_IN  AutoscalingHistoryActionType
}

func GetAutoscalingHistoryActionTypeEnum() AutoscalingHistoryActionTypeEnum {
	return AutoscalingHistoryActionTypeEnum{
		SCALE_OUT: AutoscalingHistoryActionType{
			value: "SCALE_OUT",
		},
		SCALE_IN: AutoscalingHistoryActionType{
			value: "SCALE_IN",
		},
	}
}

func (c AutoscalingHistoryActionType) Value() string {
	return c.value
}

func (c AutoscalingHistoryActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoscalingHistoryActionType) UnmarshalJSON(b []byte) error {
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
