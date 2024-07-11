package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TransactionSwitchReq struct {

	// 开关状态
	SwitchStatus TransactionSwitchReqSwitchStatus `json:"switch_status"`

	// 数据库类型。仅支持MySQL
	DatastoreType TransactionSwitchReqDatastoreType `json:"datastore_type"`
}

func (o TransactionSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransactionSwitchReq struct{}"
	}

	return strings.Join([]string{"TransactionSwitchReq", string(data)}, " ")
}

type TransactionSwitchReqSwitchStatus struct {
	value string
}

type TransactionSwitchReqSwitchStatusEnum struct {
	ENABLED  TransactionSwitchReqSwitchStatus
	DISABLED TransactionSwitchReqSwitchStatus
}

func GetTransactionSwitchReqSwitchStatusEnum() TransactionSwitchReqSwitchStatusEnum {
	return TransactionSwitchReqSwitchStatusEnum{
		ENABLED: TransactionSwitchReqSwitchStatus{
			value: "Enabled",
		},
		DISABLED: TransactionSwitchReqSwitchStatus{
			value: "Disabled",
		},
	}
}

func (c TransactionSwitchReqSwitchStatus) Value() string {
	return c.value
}

func (c TransactionSwitchReqSwitchStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransactionSwitchReqSwitchStatus) UnmarshalJSON(b []byte) error {
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

type TransactionSwitchReqDatastoreType struct {
	value string
}

type TransactionSwitchReqDatastoreTypeEnum struct {
	MY_SQL TransactionSwitchReqDatastoreType
}

func GetTransactionSwitchReqDatastoreTypeEnum() TransactionSwitchReqDatastoreTypeEnum {
	return TransactionSwitchReqDatastoreTypeEnum{
		MY_SQL: TransactionSwitchReqDatastoreType{
			value: "MySQL",
		},
	}
}

func (c TransactionSwitchReqDatastoreType) Value() string {
	return c.value
}

func (c TransactionSwitchReqDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransactionSwitchReqDatastoreType) UnmarshalJSON(b []byte) error {
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
