package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWelfareAreaInfoRequest Request Object
type ShowWelfareAreaInfoRequest struct {

	// language
	XLanguage *ShowWelfareAreaInfoRequestXLanguage `json:"x-language,omitempty"`
}

func (o ShowWelfareAreaInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWelfareAreaInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowWelfareAreaInfoRequest", string(data)}, " ")
}

type ShowWelfareAreaInfoRequestXLanguage struct {
	value string
}

type ShowWelfareAreaInfoRequestXLanguageEnum struct {
	ZH_CN ShowWelfareAreaInfoRequestXLanguage
	EN_US ShowWelfareAreaInfoRequestXLanguage
}

func GetShowWelfareAreaInfoRequestXLanguageEnum() ShowWelfareAreaInfoRequestXLanguageEnum {
	return ShowWelfareAreaInfoRequestXLanguageEnum{
		ZH_CN: ShowWelfareAreaInfoRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowWelfareAreaInfoRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowWelfareAreaInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ShowWelfareAreaInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWelfareAreaInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
