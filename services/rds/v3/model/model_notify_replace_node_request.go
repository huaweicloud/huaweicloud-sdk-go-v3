package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NotifyReplaceNodeRequest Request Object
type NotifyReplaceNodeRequest struct {

	// 语言。
	XLanguage *NotifyReplaceNodeRequestXLanguage `json:"X-Language,omitempty"`

	// 只读实例ID。
	InstanceId string `json:"instance_id"`

	Body *ReplaceNodeRequest `json:"body,omitempty"`
}

func (o NotifyReplaceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyReplaceNodeRequest struct{}"
	}

	return strings.Join([]string{"NotifyReplaceNodeRequest", string(data)}, " ")
}

type NotifyReplaceNodeRequestXLanguage struct {
	value string
}

type NotifyReplaceNodeRequestXLanguageEnum struct {
	ZH_CN NotifyReplaceNodeRequestXLanguage
	EN_US NotifyReplaceNodeRequestXLanguage
}

func GetNotifyReplaceNodeRequestXLanguageEnum() NotifyReplaceNodeRequestXLanguageEnum {
	return NotifyReplaceNodeRequestXLanguageEnum{
		ZH_CN: NotifyReplaceNodeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: NotifyReplaceNodeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c NotifyReplaceNodeRequestXLanguage) Value() string {
	return c.value
}

func (c NotifyReplaceNodeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotifyReplaceNodeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
