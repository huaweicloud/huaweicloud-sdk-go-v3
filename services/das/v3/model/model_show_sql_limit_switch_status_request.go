package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlLimitSwitchStatusRequest Request Object
type ShowSqlLimitSwitchStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ShowSqlLimitSwitchStatusRequestXLanguage `json:"X-Language,omitempty"`

	// 数据库类型
	DatastoreType ShowSqlLimitSwitchStatusRequestDatastoreType `json:"datastore_type"`
}

func (o ShowSqlLimitSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitSwitchStatusRequest", string(data)}, " ")
}

type ShowSqlLimitSwitchStatusRequestXLanguage struct {
	value string
}

type ShowSqlLimitSwitchStatusRequestXLanguageEnum struct {
	ZH_CN ShowSqlLimitSwitchStatusRequestXLanguage
	EN_US ShowSqlLimitSwitchStatusRequestXLanguage
}

func GetShowSqlLimitSwitchStatusRequestXLanguageEnum() ShowSqlLimitSwitchStatusRequestXLanguageEnum {
	return ShowSqlLimitSwitchStatusRequestXLanguageEnum{
		ZH_CN: ShowSqlLimitSwitchStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSqlLimitSwitchStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSqlLimitSwitchStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSqlLimitSwitchStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlLimitSwitchStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowSqlLimitSwitchStatusRequestDatastoreType struct {
	value string
}

type ShowSqlLimitSwitchStatusRequestDatastoreTypeEnum struct {
	MY_SQL      ShowSqlLimitSwitchStatusRequestDatastoreType
	POSTGRE_SQL ShowSqlLimitSwitchStatusRequestDatastoreType
}

func GetShowSqlLimitSwitchStatusRequestDatastoreTypeEnum() ShowSqlLimitSwitchStatusRequestDatastoreTypeEnum {
	return ShowSqlLimitSwitchStatusRequestDatastoreTypeEnum{
		MY_SQL: ShowSqlLimitSwitchStatusRequestDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: ShowSqlLimitSwitchStatusRequestDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c ShowSqlLimitSwitchStatusRequestDatastoreType) Value() string {
	return c.value
}

func (c ShowSqlLimitSwitchStatusRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlLimitSwitchStatusRequestDatastoreType) UnmarshalJSON(b []byte) error {
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
