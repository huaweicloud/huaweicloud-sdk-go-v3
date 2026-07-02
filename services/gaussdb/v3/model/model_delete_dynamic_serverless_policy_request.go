package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDynamicServerlessPolicyRequest Request Object
type DeleteDynamicServerlessPolicyRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us
	XLanguage *DeleteDynamicServerlessPolicyRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`
}

func (o DeleteDynamicServerlessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDynamicServerlessPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteDynamicServerlessPolicyRequest", string(data)}, " ")
}

type DeleteDynamicServerlessPolicyRequestXLanguage struct {
	value string
}

type DeleteDynamicServerlessPolicyRequestXLanguageEnum struct {
	ZH_CN DeleteDynamicServerlessPolicyRequestXLanguage
	EN_US DeleteDynamicServerlessPolicyRequestXLanguage
}

func GetDeleteDynamicServerlessPolicyRequestXLanguageEnum() DeleteDynamicServerlessPolicyRequestXLanguageEnum {
	return DeleteDynamicServerlessPolicyRequestXLanguageEnum{
		ZH_CN: DeleteDynamicServerlessPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteDynamicServerlessPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteDynamicServerlessPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteDynamicServerlessPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDynamicServerlessPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
