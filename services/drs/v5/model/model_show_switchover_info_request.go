package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSwitchoverInfoRequest Request Object
type ShowSwitchoverInfoRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowSwitchoverInfoRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowSwitchoverInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSwitchoverInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSwitchoverInfoRequest", string(data)}, " ")
}

type ShowSwitchoverInfoRequestXLanguage struct {
	value string
}

type ShowSwitchoverInfoRequestXLanguageEnum struct {
	EN_US ShowSwitchoverInfoRequestXLanguage
	ZH_CN ShowSwitchoverInfoRequestXLanguage
}

func GetShowSwitchoverInfoRequestXLanguageEnum() ShowSwitchoverInfoRequestXLanguageEnum {
	return ShowSwitchoverInfoRequestXLanguageEnum{
		EN_US: ShowSwitchoverInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSwitchoverInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSwitchoverInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSwitchoverInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSwitchoverInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
