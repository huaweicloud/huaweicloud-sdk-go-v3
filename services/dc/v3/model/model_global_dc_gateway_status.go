package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GlobalDcGatewayStatus GDGW资源状态，DOWN (未连接状态), PENDING_UPDATE（连接状态更新中），ACTIVE (已连接状态), ERROR (出错)。
type GlobalDcGatewayStatus struct {
	value string
}

type GlobalDcGatewayStatusEnum struct {
	DOWN           GlobalDcGatewayStatus
	PENDING_UPDATE GlobalDcGatewayStatus
	ACTIVE         GlobalDcGatewayStatus
	ERROR          GlobalDcGatewayStatus
}

func GetGlobalDcGatewayStatusEnum() GlobalDcGatewayStatusEnum {
	return GlobalDcGatewayStatusEnum{
		DOWN: GlobalDcGatewayStatus{
			value: "DOWN",
		},
		PENDING_UPDATE: GlobalDcGatewayStatus{
			value: "PENDING_UPDATE",
		},
		ACTIVE: GlobalDcGatewayStatus{
			value: "ACTIVE",
		},
		ERROR: GlobalDcGatewayStatus{
			value: "ERROR",
		},
	}
}

func (c GlobalDcGatewayStatus) Value() string {
	return c.value
}

func (c GlobalDcGatewayStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalDcGatewayStatus) UnmarshalJSON(b []byte) error {
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
