package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobListRequest Request Object
type ShowJobListRequest struct {

	// 请求语言类型
	XLanguage *ShowJobListRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryJobsReq `json:"body,omitempty"`
}

func (o ShowJobListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobListRequest struct{}"
	}

	return strings.Join([]string{"ShowJobListRequest", string(data)}, " ")
}

type ShowJobListRequestXLanguage struct {
	value string
}

type ShowJobListRequestXLanguageEnum struct {
	EN_US ShowJobListRequestXLanguage
	ZH_CN ShowJobListRequestXLanguage
}

func GetShowJobListRequestXLanguageEnum() ShowJobListRequestXLanguageEnum {
	return ShowJobListRequestXLanguageEnum{
		EN_US: ShowJobListRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowJobListRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowJobListRequestXLanguage) Value() string {
	return c.value
}

func (c ShowJobListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobListRequestXLanguage) UnmarshalJSON(b []byte) error {
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
