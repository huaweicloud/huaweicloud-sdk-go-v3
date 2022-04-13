package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ChangeSqlSwitchRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 请求语言类型。

	XLanguage *ChangeSqlSwitchRequestXLanguage `json:"X-Language,omitempty"`

	Body *ChangeSqlSwitchBody `json:"body,omitempty"`
}

func (o ChangeSqlSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSqlSwitchRequest struct{}"
	}

	return strings.Join([]string{"ChangeSqlSwitchRequest", string(data)}, " ")
}

type ChangeSqlSwitchRequestXLanguage struct {
	value string
}

type ChangeSqlSwitchRequestXLanguageEnum struct {
	EN_US ChangeSqlSwitchRequestXLanguage
	ZH_CN ChangeSqlSwitchRequestXLanguage
}

func GetChangeSqlSwitchRequestXLanguageEnum() ChangeSqlSwitchRequestXLanguageEnum {
	return ChangeSqlSwitchRequestXLanguageEnum{
		EN_US: ChangeSqlSwitchRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ChangeSqlSwitchRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ChangeSqlSwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeSqlSwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
