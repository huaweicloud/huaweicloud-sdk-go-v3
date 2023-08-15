package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeSqlLimitSwitchStatusRequest Request Object
type ChangeSqlLimitSwitchStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ChangeSqlLimitSwitchStatusRequestXLanguage `json:"X-Language,omitempty"`

	Body *ChangeSqlLimitSwitchStatusBody `json:"body,omitempty"`
}

func (o ChangeSqlLimitSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSqlLimitSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeSqlLimitSwitchStatusRequest", string(data)}, " ")
}

type ChangeSqlLimitSwitchStatusRequestXLanguage struct {
	value string
}

type ChangeSqlLimitSwitchStatusRequestXLanguageEnum struct {
	ZH_CN ChangeSqlLimitSwitchStatusRequestXLanguage
	EN_US ChangeSqlLimitSwitchStatusRequestXLanguage
}

func GetChangeSqlLimitSwitchStatusRequestXLanguageEnum() ChangeSqlLimitSwitchStatusRequestXLanguageEnum {
	return ChangeSqlLimitSwitchStatusRequestXLanguageEnum{
		ZH_CN: ChangeSqlLimitSwitchStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeSqlLimitSwitchStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeSqlLimitSwitchStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeSqlLimitSwitchStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeSqlLimitSwitchStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
