package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPositionResultRequest Request Object
type ShowPositionResultRequest struct {

	// 请求语言类型。
	XLanguage *ShowPositionResultRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 位点信息采集的ID，由采集数据库位点信息接口返回的ID。
	QueryId string `json:"query_id"`
}

func (o ShowPositionResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPositionResultRequest struct{}"
	}

	return strings.Join([]string{"ShowPositionResultRequest", string(data)}, " ")
}

type ShowPositionResultRequestXLanguage struct {
	value string
}

type ShowPositionResultRequestXLanguageEnum struct {
	EN_US ShowPositionResultRequestXLanguage
	ZH_CN ShowPositionResultRequestXLanguage
}

func GetShowPositionResultRequestXLanguageEnum() ShowPositionResultRequestXLanguageEnum {
	return ShowPositionResultRequestXLanguageEnum{
		EN_US: ShowPositionResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowPositionResultRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowPositionResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowPositionResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPositionResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
