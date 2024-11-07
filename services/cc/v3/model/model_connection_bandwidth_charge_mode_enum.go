package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConnectionBandwidthChargeModeEnum - bwd (按需按带宽计费) - trf (按流量计费) - 95 (传统95计费) - 95avr (日95计费)
type ConnectionBandwidthChargeModeEnum struct {
	value string
}

type ConnectionBandwidthChargeModeEnumEnum struct {
	BWD     ConnectionBandwidthChargeModeEnum
	TRF     ConnectionBandwidthChargeModeEnum
	E_95    ConnectionBandwidthChargeModeEnum
	E_95AVR ConnectionBandwidthChargeModeEnum
}

func GetConnectionBandwidthChargeModeEnumEnum() ConnectionBandwidthChargeModeEnumEnum {
	return ConnectionBandwidthChargeModeEnumEnum{
		BWD: ConnectionBandwidthChargeModeEnum{
			value: "bwd",
		},
		TRF: ConnectionBandwidthChargeModeEnum{
			value: "trf",
		},
		E_95: ConnectionBandwidthChargeModeEnum{
			value: "95",
		},
		E_95AVR: ConnectionBandwidthChargeModeEnum{
			value: "95avr",
		},
	}
}

func (c ConnectionBandwidthChargeModeEnum) Value() string {
	return c.value
}

func (c ConnectionBandwidthChargeModeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionBandwidthChargeModeEnum) UnmarshalJSON(b []byte) error {
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
