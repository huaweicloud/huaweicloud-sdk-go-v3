package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowActionsRequest Request Object
type ShowActionsRequest struct {

	// 请求语言类型。
	XLanguage *ShowActionsRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActionsRequest struct{}"
	}

	return strings.Join([]string{"ShowActionsRequest", string(data)}, " ")
}

type ShowActionsRequestXLanguage struct {
	value string
}

type ShowActionsRequestXLanguageEnum struct {
	EN_US ShowActionsRequestXLanguage
	ZH_CN ShowActionsRequestXLanguage
}

func GetShowActionsRequestXLanguageEnum() ShowActionsRequestXLanguageEnum {
	return ShowActionsRequestXLanguageEnum{
		EN_US: ShowActionsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowActionsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowActionsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowActionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowActionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
