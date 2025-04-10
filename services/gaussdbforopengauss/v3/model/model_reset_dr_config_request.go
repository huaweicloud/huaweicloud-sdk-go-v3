package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResetDrConfigRequest Request Object
type ResetDrConfigRequest struct {

	// 语言。
	XLanguage *ResetDrConfigRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *ResetDrConfigRequestBody `json:"body,omitempty"`
}

func (o ResetDrConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDrConfigRequest struct{}"
	}

	return strings.Join([]string{"ResetDrConfigRequest", string(data)}, " ")
}

type ResetDrConfigRequestXLanguage struct {
	value string
}

type ResetDrConfigRequestXLanguageEnum struct {
	ZH_CN ResetDrConfigRequestXLanguage
	EN_US ResetDrConfigRequestXLanguage
}

func GetResetDrConfigRequestXLanguageEnum() ResetDrConfigRequestXLanguageEnum {
	return ResetDrConfigRequestXLanguageEnum{
		ZH_CN: ResetDrConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ResetDrConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ResetDrConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ResetDrConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetDrConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
