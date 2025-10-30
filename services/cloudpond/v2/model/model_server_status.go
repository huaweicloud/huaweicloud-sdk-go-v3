package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerStatus - PENDING_PAYMENT：待支付 - DELIVERING：交付中 - USING：使用中
type ServerStatus struct {
	value string
}

type ServerStatusEnum struct {
	PENDING_PAYMENT ServerStatus
	DELIVERING      ServerStatus
	USING           ServerStatus
}

func GetServerStatusEnum() ServerStatusEnum {
	return ServerStatusEnum{
		PENDING_PAYMENT: ServerStatus{
			value: "PENDING_PAYMENT",
		},
		DELIVERING: ServerStatus{
			value: "DELIVERING",
		},
		USING: ServerStatus{
			value: "USING",
		},
	}
}

func (c ServerStatus) Value() string {
	return c.value
}

func (c ServerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerStatus) UnmarshalJSON(b []byte) error {
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
