package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeSqlLimitSwitchStatusBody 设置SQL限流开关状态请求体
type ChangeSqlLimitSwitchStatusBody struct {

	// 开关状态
	SwitchStatus ChangeSqlLimitSwitchStatusBodySwitchStatus `json:"switch_status"`

	// 数据库类型
	DatastoreType ChangeSqlLimitSwitchStatusBodyDatastoreType `json:"datastore_type"`
}

func (o ChangeSqlLimitSwitchStatusBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSqlLimitSwitchStatusBody struct{}"
	}

	return strings.Join([]string{"ChangeSqlLimitSwitchStatusBody", string(data)}, " ")
}

type ChangeSqlLimitSwitchStatusBodySwitchStatus struct {
	value string
}

type ChangeSqlLimitSwitchStatusBodySwitchStatusEnum struct {
	ON  ChangeSqlLimitSwitchStatusBodySwitchStatus
	OFF ChangeSqlLimitSwitchStatusBodySwitchStatus
}

func GetChangeSqlLimitSwitchStatusBodySwitchStatusEnum() ChangeSqlLimitSwitchStatusBodySwitchStatusEnum {
	return ChangeSqlLimitSwitchStatusBodySwitchStatusEnum{
		ON: ChangeSqlLimitSwitchStatusBodySwitchStatus{
			value: "ON",
		},
		OFF: ChangeSqlLimitSwitchStatusBodySwitchStatus{
			value: "OFF",
		},
	}
}

func (c ChangeSqlLimitSwitchStatusBodySwitchStatus) Value() string {
	return c.value
}

func (c ChangeSqlLimitSwitchStatusBodySwitchStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeSqlLimitSwitchStatusBodySwitchStatus) UnmarshalJSON(b []byte) error {
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

type ChangeSqlLimitSwitchStatusBodyDatastoreType struct {
	value string
}

type ChangeSqlLimitSwitchStatusBodyDatastoreTypeEnum struct {
	MY_SQL      ChangeSqlLimitSwitchStatusBodyDatastoreType
	POSTGRE_SQL ChangeSqlLimitSwitchStatusBodyDatastoreType
}

func GetChangeSqlLimitSwitchStatusBodyDatastoreTypeEnum() ChangeSqlLimitSwitchStatusBodyDatastoreTypeEnum {
	return ChangeSqlLimitSwitchStatusBodyDatastoreTypeEnum{
		MY_SQL: ChangeSqlLimitSwitchStatusBodyDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: ChangeSqlLimitSwitchStatusBodyDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c ChangeSqlLimitSwitchStatusBodyDatastoreType) Value() string {
	return c.value
}

func (c ChangeSqlLimitSwitchStatusBodyDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeSqlLimitSwitchStatusBodyDatastoreType) UnmarshalJSON(b []byte) error {
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
