package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExchangeInstanceVipProxyIpRequest Request Object
type ExchangeInstanceVipProxyIpRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *ExchangeInstanceVipProxyIpRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  数据库代理ID，此参数是数据库代理的唯一标识。  获取方法请参见[查询数据库代理](https://support.huaweicloud.com/api-taurusdb/taurusdb_11_0016.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为po01，长度为36个字符。  **默认取值**：  不涉及。
	ProxyId string `json:"proxy_id"`
}

func (o ExchangeInstanceVipProxyIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExchangeInstanceVipProxyIpRequest struct{}"
	}

	return strings.Join([]string{"ExchangeInstanceVipProxyIpRequest", string(data)}, " ")
}

type ExchangeInstanceVipProxyIpRequestXLanguage struct {
	value string
}

type ExchangeInstanceVipProxyIpRequestXLanguageEnum struct {
	EN_US ExchangeInstanceVipProxyIpRequestXLanguage
	ZH_CN ExchangeInstanceVipProxyIpRequestXLanguage
}

func GetExchangeInstanceVipProxyIpRequestXLanguageEnum() ExchangeInstanceVipProxyIpRequestXLanguageEnum {
	return ExchangeInstanceVipProxyIpRequestXLanguageEnum{
		EN_US: ExchangeInstanceVipProxyIpRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExchangeInstanceVipProxyIpRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExchangeInstanceVipProxyIpRequestXLanguage) Value() string {
	return c.value
}

func (c ExchangeInstanceVipProxyIpRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExchangeInstanceVipProxyIpRequestXLanguage) UnmarshalJSON(b []byte) error {
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
