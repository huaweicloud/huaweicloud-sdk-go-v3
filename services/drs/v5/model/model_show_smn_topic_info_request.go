package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSmnTopicInfoRequest Request Object
type ShowSmnTopicInfoRequest struct {

	// 请求语言类型。
	XLanguage *ShowSmnTopicInfoRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowSmnTopicInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmnTopicInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSmnTopicInfoRequest", string(data)}, " ")
}

type ShowSmnTopicInfoRequestXLanguage struct {
	value string
}

type ShowSmnTopicInfoRequestXLanguageEnum struct {
	EN_US ShowSmnTopicInfoRequestXLanguage
	ZH_CN ShowSmnTopicInfoRequestXLanguage
}

func GetShowSmnTopicInfoRequestXLanguageEnum() ShowSmnTopicInfoRequestXLanguageEnum {
	return ShowSmnTopicInfoRequestXLanguageEnum{
		EN_US: ShowSmnTopicInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSmnTopicInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSmnTopicInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSmnTopicInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmnTopicInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
