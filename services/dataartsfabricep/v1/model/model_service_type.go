package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServiceType Service的类型，可选值： - PGSQL_SERVICE：DWS Pay-By-Query - LLM_MODEL：大语言模型
type ServiceType struct {
	value string
}

type ServiceTypeEnum struct {
	PGSQL_SERVICE ServiceType
	LLM_MODEL     ServiceType
}

func GetServiceTypeEnum() ServiceTypeEnum {
	return ServiceTypeEnum{
		PGSQL_SERVICE: ServiceType{
			value: "PGSQL_SERVICE",
		},
		LLM_MODEL: ServiceType{
			value: "LLM_MODEL",
		},
	}
}

func (c ServiceType) Value() string {
	return c.value
}

func (c ServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServiceType) UnmarshalJSON(b []byte) error {
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
