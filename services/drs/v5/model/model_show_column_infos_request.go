package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowColumnInfosRequest Request Object
type ShowColumnInfosRequest struct {

	// 请求语言类型。
	XLanguage *ShowColumnInfosRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *ShowColumnInfoReq `json:"body,omitempty"`
}

func (o ShowColumnInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowColumnInfosRequest struct{}"
	}

	return strings.Join([]string{"ShowColumnInfosRequest", string(data)}, " ")
}

type ShowColumnInfosRequestXLanguage struct {
	value string
}

type ShowColumnInfosRequestXLanguageEnum struct {
	EN_US ShowColumnInfosRequestXLanguage
	ZH_CN ShowColumnInfosRequestXLanguage
}

func GetShowColumnInfosRequestXLanguageEnum() ShowColumnInfosRequestXLanguageEnum {
	return ShowColumnInfosRequestXLanguageEnum{
		EN_US: ShowColumnInfosRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowColumnInfosRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowColumnInfosRequestXLanguage) Value() string {
	return c.value
}

func (c ShowColumnInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowColumnInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
