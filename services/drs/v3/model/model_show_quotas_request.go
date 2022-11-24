package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowQuotasRequest struct {

	// 请求语言类型。
	XLanguage *ShowQuotasRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}

type ShowQuotasRequestXLanguage struct {
	value string
}

type ShowQuotasRequestXLanguageEnum struct {
	EN_US ShowQuotasRequestXLanguage
	ZH_CN ShowQuotasRequestXLanguage
}

func GetShowQuotasRequestXLanguageEnum() ShowQuotasRequestXLanguageEnum {
	return ShowQuotasRequestXLanguageEnum{
		EN_US: ShowQuotasRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowQuotasRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowQuotasRequestXLanguage) Value() string {
	return c.value
}

func (c ShowQuotasRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowQuotasRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
