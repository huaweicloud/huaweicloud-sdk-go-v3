package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTransactionSwitchStatusRequest Request Object
type ShowTransactionSwitchStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库类型。仅支持MySQL
	DatastoreType ShowTransactionSwitchStatusRequestDatastoreType `json:"datastore_type"`

	// 语言
	XLanguage *ShowTransactionSwitchStatusRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowTransactionSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransactionSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowTransactionSwitchStatusRequest", string(data)}, " ")
}

type ShowTransactionSwitchStatusRequestDatastoreType struct {
	value string
}

type ShowTransactionSwitchStatusRequestDatastoreTypeEnum struct {
	MY_SQL ShowTransactionSwitchStatusRequestDatastoreType
}

func GetShowTransactionSwitchStatusRequestDatastoreTypeEnum() ShowTransactionSwitchStatusRequestDatastoreTypeEnum {
	return ShowTransactionSwitchStatusRequestDatastoreTypeEnum{
		MY_SQL: ShowTransactionSwitchStatusRequestDatastoreType{
			value: "MySQL",
		},
	}
}

func (c ShowTransactionSwitchStatusRequestDatastoreType) Value() string {
	return c.value
}

func (c ShowTransactionSwitchStatusRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTransactionSwitchStatusRequestDatastoreType) UnmarshalJSON(b []byte) error {
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

type ShowTransactionSwitchStatusRequestXLanguage struct {
	value string
}

type ShowTransactionSwitchStatusRequestXLanguageEnum struct {
	ZH_CN ShowTransactionSwitchStatusRequestXLanguage
	EN_US ShowTransactionSwitchStatusRequestXLanguage
}

func GetShowTransactionSwitchStatusRequestXLanguageEnum() ShowTransactionSwitchStatusRequestXLanguageEnum {
	return ShowTransactionSwitchStatusRequestXLanguageEnum{
		ZH_CN: ShowTransactionSwitchStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowTransactionSwitchStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowTransactionSwitchStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowTransactionSwitchStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTransactionSwitchStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
