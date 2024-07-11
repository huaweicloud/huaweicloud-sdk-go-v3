package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeTransactionSwitchStatusRequest Request Object
type ChangeTransactionSwitchStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ChangeTransactionSwitchStatusRequestXLanguage `json:"X-Language,omitempty"`

	Body *TransactionSwitchReq `json:"body,omitempty"`
}

func (o ChangeTransactionSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTransactionSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeTransactionSwitchStatusRequest", string(data)}, " ")
}

type ChangeTransactionSwitchStatusRequestXLanguage struct {
	value string
}

type ChangeTransactionSwitchStatusRequestXLanguageEnum struct {
	ZH_CN ChangeTransactionSwitchStatusRequestXLanguage
	EN_US ChangeTransactionSwitchStatusRequestXLanguage
}

func GetChangeTransactionSwitchStatusRequestXLanguageEnum() ChangeTransactionSwitchStatusRequestXLanguageEnum {
	return ChangeTransactionSwitchStatusRequestXLanguageEnum{
		ZH_CN: ChangeTransactionSwitchStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeTransactionSwitchStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeTransactionSwitchStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeTransactionSwitchStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeTransactionSwitchStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
