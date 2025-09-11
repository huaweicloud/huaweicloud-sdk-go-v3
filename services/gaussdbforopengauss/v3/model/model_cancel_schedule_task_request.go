package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CancelScheduleTaskRequest Request Object
type CancelScheduleTaskRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *CancelScheduleTaskRequestXLanguage `json:"X-Language,omitempty"`

	// 任务id。
	TaskId string `json:"task_id"`
}

func (o CancelScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelScheduleTaskRequest", string(data)}, " ")
}

type CancelScheduleTaskRequestXLanguage struct {
	value string
}

type CancelScheduleTaskRequestXLanguageEnum struct {
	ZH_CN CancelScheduleTaskRequestXLanguage
	EN_US CancelScheduleTaskRequestXLanguage
}

func GetCancelScheduleTaskRequestXLanguageEnum() CancelScheduleTaskRequestXLanguageEnum {
	return CancelScheduleTaskRequestXLanguageEnum{
		ZH_CN: CancelScheduleTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CancelScheduleTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CancelScheduleTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CancelScheduleTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelScheduleTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
