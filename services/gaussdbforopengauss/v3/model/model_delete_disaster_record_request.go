package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDisasterRecordRequest Request Object
type DeleteDisasterRecordRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *DeleteDisasterRecordRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteDisasterRecordRequestBody `json:"body,omitempty"`
}

func (o DeleteDisasterRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecordRequest struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecordRequest", string(data)}, " ")
}

type DeleteDisasterRecordRequestXLanguage struct {
	value string
}

type DeleteDisasterRecordRequestXLanguageEnum struct {
	ZH_CN DeleteDisasterRecordRequestXLanguage
	EN_US DeleteDisasterRecordRequestXLanguage
}

func GetDeleteDisasterRecordRequestXLanguageEnum() DeleteDisasterRecordRequestXLanguageEnum {
	return DeleteDisasterRecordRequestXLanguageEnum{
		ZH_CN: DeleteDisasterRecordRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteDisasterRecordRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteDisasterRecordRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteDisasterRecordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDisasterRecordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
