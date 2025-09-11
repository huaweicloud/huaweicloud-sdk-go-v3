package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteScheduleTaskRequest Request Object
type DeleteScheduleTaskRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *DeleteScheduleTaskRequestXLanguage `json:"X-Language,omitempty"`

	// 任务id。
	TaskId string `json:"task_id"`
}

func (o DeleteScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteScheduleTaskRequest", string(data)}, " ")
}

type DeleteScheduleTaskRequestXLanguage struct {
	value string
}

type DeleteScheduleTaskRequestXLanguageEnum struct {
	ZH_CN DeleteScheduleTaskRequestXLanguage
	EN_US DeleteScheduleTaskRequestXLanguage
}

func GetDeleteScheduleTaskRequestXLanguageEnum() DeleteScheduleTaskRequestXLanguageEnum {
	return DeleteScheduleTaskRequestXLanguageEnum{
		ZH_CN: DeleteScheduleTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteScheduleTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteScheduleTaskRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteScheduleTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteScheduleTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
