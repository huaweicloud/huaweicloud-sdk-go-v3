package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckScheduleTaskExistRequest Request Object
type CheckScheduleTaskExistRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *CheckScheduleTaskExistRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *CheckScheduleTaskExistRequestBody `json:"body,omitempty"`
}

func (o CheckScheduleTaskExistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckScheduleTaskExistRequest struct{}"
	}

	return strings.Join([]string{"CheckScheduleTaskExistRequest", string(data)}, " ")
}

type CheckScheduleTaskExistRequestXLanguage struct {
	value string
}

type CheckScheduleTaskExistRequestXLanguageEnum struct {
	ZH_CN CheckScheduleTaskExistRequestXLanguage
	EN_US CheckScheduleTaskExistRequestXLanguage
}

func GetCheckScheduleTaskExistRequestXLanguageEnum() CheckScheduleTaskExistRequestXLanguageEnum {
	return CheckScheduleTaskExistRequestXLanguageEnum{
		ZH_CN: CheckScheduleTaskExistRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CheckScheduleTaskExistRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CheckScheduleTaskExistRequestXLanguage) Value() string {
	return c.value
}

func (c CheckScheduleTaskExistRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckScheduleTaskExistRequestXLanguage) UnmarshalJSON(b []byte) error {
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
