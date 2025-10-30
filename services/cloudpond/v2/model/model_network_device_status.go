package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NetworkDeviceStatus - PENDING_PAYMENT：待支付 - DELIVERING：交付中 - USING：使用中
type NetworkDeviceStatus struct {
	value string
}

type NetworkDeviceStatusEnum struct {
	PENDING_PAYMENT NetworkDeviceStatus
	DELIVERING      NetworkDeviceStatus
	USING           NetworkDeviceStatus
}

func GetNetworkDeviceStatusEnum() NetworkDeviceStatusEnum {
	return NetworkDeviceStatusEnum{
		PENDING_PAYMENT: NetworkDeviceStatus{
			value: "PENDING_PAYMENT",
		},
		DELIVERING: NetworkDeviceStatus{
			value: "DELIVERING",
		},
		USING: NetworkDeviceStatus{
			value: "USING",
		},
	}
}

func (c NetworkDeviceStatus) Value() string {
	return c.value
}

func (c NetworkDeviceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkDeviceStatus) UnmarshalJSON(b []byte) error {
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
