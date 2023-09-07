package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataFilteringResultRequest Request Object
type ShowDataFilteringResultRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDataFilteringResultRequestXLanguage `json:"X-Language,omitempty"`

	// 数据过滤校验请求ID
	QueryId string `json:"query_id"`
}

func (o ShowDataFilteringResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataFilteringResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDataFilteringResultRequest", string(data)}, " ")
}

type ShowDataFilteringResultRequestXLanguage struct {
	value string
}

type ShowDataFilteringResultRequestXLanguageEnum struct {
	EN_US ShowDataFilteringResultRequestXLanguage
	ZH_CN ShowDataFilteringResultRequestXLanguage
}

func GetShowDataFilteringResultRequestXLanguageEnum() ShowDataFilteringResultRequestXLanguageEnum {
	return ShowDataFilteringResultRequestXLanguageEnum{
		EN_US: ShowDataFilteringResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDataFilteringResultRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDataFilteringResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDataFilteringResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataFilteringResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
