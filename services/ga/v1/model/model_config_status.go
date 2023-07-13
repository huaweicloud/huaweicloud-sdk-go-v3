package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigStatus 配置状态，取值： - ACTIVE：运行中 - PENDING：待定 - ERROR：错误 - DELETING：正在删除
type ConfigStatus struct {
	value string
}

type ConfigStatusEnum struct {
	ACTIVE   ConfigStatus
	PENDING  ConfigStatus
	ERROR    ConfigStatus
	DELETING ConfigStatus
}

func GetConfigStatusEnum() ConfigStatusEnum {
	return ConfigStatusEnum{
		ACTIVE: ConfigStatus{
			value: "ACTIVE",
		},
		PENDING: ConfigStatus{
			value: "PENDING",
		},
		ERROR: ConfigStatus{
			value: "ERROR",
		},
		DELETING: ConfigStatus{
			value: "DELETING",
		},
	}
}

func (c ConfigStatus) Value() string {
	return c.value
}

func (c ConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigStatus) UnmarshalJSON(b []byte) error {
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
