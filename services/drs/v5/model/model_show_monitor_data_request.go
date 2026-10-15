package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMonitorDataRequest Request Object
type ShowMonitorDataRequest struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId string `json:"job_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文 - zh-cn：中文  **默认取值**：  en-us
	XLanguage *ShowMonitorDataRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowMonitorDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitorDataRequest", string(data)}, " ")
}

type ShowMonitorDataRequestXLanguage struct {
	value string
}

type ShowMonitorDataRequestXLanguageEnum struct {
	EN_US ShowMonitorDataRequestXLanguage
	ZH_CN ShowMonitorDataRequestXLanguage
}

func GetShowMonitorDataRequestXLanguageEnum() ShowMonitorDataRequestXLanguageEnum {
	return ShowMonitorDataRequestXLanguageEnum{
		EN_US: ShowMonitorDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowMonitorDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowMonitorDataRequestXLanguage) Value() string {
	return c.value
}

func (c ShowMonitorDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitorDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
