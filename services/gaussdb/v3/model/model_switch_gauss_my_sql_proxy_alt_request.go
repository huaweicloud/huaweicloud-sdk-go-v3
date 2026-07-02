package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchGaussMySqlProxyAltRequest Request Object
type SwitchGaussMySqlProxyAltRequest struct {

	// **参数解释**：  内容类型。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  数据库代理ID，此参数是数据库代理的唯一标识。  获取方法请参见[查询数据库代理信息列表](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlProxyList.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为po01，长度为36个字符。  **默认取值**：  不涉及。
	ProxyId string `json:"proxy_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage *SwitchGaussMySqlProxyAltRequestXLanguage `json:"X-Language,omitempty"`

	Body *SwitchGaussMySqlProxyAltRequestBody `json:"body,omitempty"`
}

func (o SwitchGaussMySqlProxyAltRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyAltRequest struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyAltRequest", string(data)}, " ")
}

type SwitchGaussMySqlProxyAltRequestXLanguage struct {
	value string
}

type SwitchGaussMySqlProxyAltRequestXLanguageEnum struct {
	ZH_CN SwitchGaussMySqlProxyAltRequestXLanguage
	EN_US SwitchGaussMySqlProxyAltRequestXLanguage
}

func GetSwitchGaussMySqlProxyAltRequestXLanguageEnum() SwitchGaussMySqlProxyAltRequestXLanguageEnum {
	return SwitchGaussMySqlProxyAltRequestXLanguageEnum{
		ZH_CN: SwitchGaussMySqlProxyAltRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchGaussMySqlProxyAltRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchGaussMySqlProxyAltRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchGaussMySqlProxyAltRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchGaussMySqlProxyAltRequestXLanguage) UnmarshalJSON(b []byte) error {
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
