package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyEpsQuotaRequest Request Object
type ModifyEpsQuotaRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ModifyEpsQuotaRequestXLanguage `json:"X-Language,omitempty"`

	Body *ModifyEpsQuotaRequestBody `json:"body,omitempty"`
}

func (o ModifyEpsQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEpsQuotaRequest struct{}"
	}

	return strings.Join([]string{"ModifyEpsQuotaRequest", string(data)}, " ")
}

type ModifyEpsQuotaRequestXLanguage struct {
	value string
}

type ModifyEpsQuotaRequestXLanguageEnum struct {
	ZH_CN ModifyEpsQuotaRequestXLanguage
	EN_US ModifyEpsQuotaRequestXLanguage
}

func GetModifyEpsQuotaRequestXLanguageEnum() ModifyEpsQuotaRequestXLanguageEnum {
	return ModifyEpsQuotaRequestXLanguageEnum{
		ZH_CN: ModifyEpsQuotaRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifyEpsQuotaRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifyEpsQuotaRequestXLanguage) Value() string {
	return c.value
}

func (c ModifyEpsQuotaRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyEpsQuotaRequestXLanguage) UnmarshalJSON(b []byte) error {
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
