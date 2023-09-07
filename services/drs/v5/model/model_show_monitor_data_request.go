package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMonitorDataRequest Request Object
type ShowMonitorDataRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowMonitorDataRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowMonitorDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitorDataRequest", string(data)}, " ")
}

type ShowMonitorDataRequestXLanguage struct {
	value string
}

type ShowMonitorDataRequestXLanguageEnum struct {
	EN_US ShowMonitorDataRequestXLanguage
	ZH_CN ShowMonitorDataRequestXLanguage
}

func GetShowMonitorDataRequestXLanguageEnum() ShowMonitorDataRequestXLanguageEnum {
	return ShowMonitorDataRequestXLanguageEnum{
		EN_US: ShowMonitorDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowMonitorDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowMonitorDataRequestXLanguage) Value() string {
	return c.value
}

func (c ShowMonitorDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitorDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
