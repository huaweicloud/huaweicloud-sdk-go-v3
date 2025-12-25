package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstallStatus **参数解释**: 采集通道下发 - READY 等待下发 - ALL_SUCCESS 全部成功 - PARTIAL_SUCCESS 部分成功 - EXCEPTION 异常  **约束限制** 不涉及 **取值范围**: - READY - ALL_SUCCESS - PARTIAL_SUCCESS - EXCEPTION  **默认值** 不涉及
type InstallStatus struct {
	value string
}

type InstallStatusEnum struct {
	READY           InstallStatus
	ALL_SUCCESS     InstallStatus
	PARTIAL_SUCCESS InstallStatus
	EXCEPTION       InstallStatus
}

func GetInstallStatusEnum() InstallStatusEnum {
	return InstallStatusEnum{
		READY: InstallStatus{
			value: "READY",
		},
		ALL_SUCCESS: InstallStatus{
			value: "ALL_SUCCESS",
		},
		PARTIAL_SUCCESS: InstallStatus{
			value: "PARTIAL_SUCCESS",
		},
		EXCEPTION: InstallStatus{
			value: "EXCEPTION",
		},
	}
}

func (c InstallStatus) Value() string {
	return c.value
}

func (c InstallStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstallStatus) UnmarshalJSON(b []byte) error {
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
