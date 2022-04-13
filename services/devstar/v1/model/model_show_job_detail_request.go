package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowJobDetailRequest struct {
	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文

	XLanguage *ShowJobDetailRequestXLanguage `json:"X-Language,omitempty"`
	// 任务 ID: 可以通过调用 RunDevstarTemplateJob 返回结果获取

	JobId string `json:"job_id"`
}

func (o ShowJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}

type ShowJobDetailRequestXLanguage struct {
	value string
}

type ShowJobDetailRequestXLanguageEnum struct {
	ZH_CN ShowJobDetailRequestXLanguage
	EN_US ShowJobDetailRequestXLanguage
}

func GetShowJobDetailRequestXLanguageEnum() ShowJobDetailRequestXLanguageEnum {
	return ShowJobDetailRequestXLanguageEnum{
		ZH_CN: ShowJobDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowJobDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowJobDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
