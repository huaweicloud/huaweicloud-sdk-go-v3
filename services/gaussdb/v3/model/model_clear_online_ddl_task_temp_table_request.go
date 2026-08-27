package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClearOnlineDdlTaskTempTableRequest Request Object
type ClearOnlineDdlTaskTempTableRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *ClearOnlineDdlTaskTempTableRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。   **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *ClearOnlineDdlTaskTempTableRequestV3 `json:"body,omitempty"`
}

func (o ClearOnlineDdlTaskTempTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearOnlineDdlTaskTempTableRequest struct{}"
	}

	return strings.Join([]string{"ClearOnlineDdlTaskTempTableRequest", string(data)}, " ")
}

type ClearOnlineDdlTaskTempTableRequestXLanguage struct {
	value string
}

type ClearOnlineDdlTaskTempTableRequestXLanguageEnum struct {
	ZH_CN ClearOnlineDdlTaskTempTableRequestXLanguage
	EN_US ClearOnlineDdlTaskTempTableRequestXLanguage
}

func GetClearOnlineDdlTaskTempTableRequestXLanguageEnum() ClearOnlineDdlTaskTempTableRequestXLanguageEnum {
	return ClearOnlineDdlTaskTempTableRequestXLanguageEnum{
		ZH_CN: ClearOnlineDdlTaskTempTableRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ClearOnlineDdlTaskTempTableRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ClearOnlineDdlTaskTempTableRequestXLanguage) Value() string {
	return c.value
}

func (c ClearOnlineDdlTaskTempTableRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClearOnlineDdlTaskTempTableRequestXLanguage) UnmarshalJSON(b []byte) error {
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
