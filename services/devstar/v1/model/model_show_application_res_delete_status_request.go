package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApplicationResDeleteStatusRequest Request Object
type ShowApplicationResDeleteStatusRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *ShowApplicationResDeleteStatusRequestXLanguage `json:"X-Language,omitempty"`

	// 应用id
	ApplicationId string `json:"application_id"`
}

func (o ShowApplicationResDeleteStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationResDeleteStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationResDeleteStatusRequest", string(data)}, " ")
}

type ShowApplicationResDeleteStatusRequestXLanguage struct {
	value string
}

type ShowApplicationResDeleteStatusRequestXLanguageEnum struct {
	ZH_CN ShowApplicationResDeleteStatusRequestXLanguage
	EN_US ShowApplicationResDeleteStatusRequestXLanguage
}

func GetShowApplicationResDeleteStatusRequestXLanguageEnum() ShowApplicationResDeleteStatusRequestXLanguageEnum {
	return ShowApplicationResDeleteStatusRequestXLanguageEnum{
		ZH_CN: ShowApplicationResDeleteStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowApplicationResDeleteStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowApplicationResDeleteStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowApplicationResDeleteStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplicationResDeleteStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
