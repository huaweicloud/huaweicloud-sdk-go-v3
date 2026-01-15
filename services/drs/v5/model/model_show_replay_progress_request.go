package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReplayProgressRequest Request Object
type ShowReplayProgressRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowReplayProgressRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowReplayProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowReplayProgressRequest", string(data)}, " ")
}

type ShowReplayProgressRequestXLanguage struct {
	value string
}

type ShowReplayProgressRequestXLanguageEnum struct {
	EN_US ShowReplayProgressRequestXLanguage
	ZH_CN ShowReplayProgressRequestXLanguage
}

func GetShowReplayProgressRequestXLanguageEnum() ShowReplayProgressRequestXLanguageEnum {
	return ShowReplayProgressRequestXLanguageEnum{
		EN_US: ShowReplayProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReplayProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReplayProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReplayProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
