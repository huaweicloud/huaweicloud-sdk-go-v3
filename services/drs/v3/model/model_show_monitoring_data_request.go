package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowMonitoringDataRequest struct {
	// 请求语言类型

	XLanguage ShowMonitoringDataRequestXLanguage `json:"X-Language"`

	Body *BatchQueryJobReq `json:"body,omitempty"`
}

func (o ShowMonitoringDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitoringDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitoringDataRequest", string(data)}, " ")
}

type ShowMonitoringDataRequestXLanguage struct {
	value string
}

type ShowMonitoringDataRequestXLanguageEnum struct {
	EN_US ShowMonitoringDataRequestXLanguage
	ZH_CN ShowMonitoringDataRequestXLanguage
}

func GetShowMonitoringDataRequestXLanguageEnum() ShowMonitoringDataRequestXLanguageEnum {
	return ShowMonitoringDataRequestXLanguageEnum{
		EN_US: ShowMonitoringDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowMonitoringDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowMonitoringDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitoringDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
