package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpgradePlanStatus struct {

	// 自动升级计划状态
	Phase UpgradePlanStatusPhase `json:"phase"`
}

func (o UpgradePlanStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradePlanStatus struct{}"
	}

	return strings.Join([]string{"UpgradePlanStatus", string(data)}, " ")
}

type UpgradePlanStatusPhase struct {
	value string
}

type UpgradePlanStatusPhaseEnum struct {
	PENDING   UpgradePlanStatusPhase
	ENQUEUED  UpgradePlanStatusPhase
	UPGRADING UpgradePlanStatusPhase
	SUCCESS   UpgradePlanStatusPhase
	FAILED    UpgradePlanStatusPhase
	SKIPPED   UpgradePlanStatusPhase
}

func GetUpgradePlanStatusPhaseEnum() UpgradePlanStatusPhaseEnum {
	return UpgradePlanStatusPhaseEnum{
		PENDING: UpgradePlanStatusPhase{
			value: "Pending",
		},
		ENQUEUED: UpgradePlanStatusPhase{
			value: "Enqueued",
		},
		UPGRADING: UpgradePlanStatusPhase{
			value: "Upgrading",
		},
		SUCCESS: UpgradePlanStatusPhase{
			value: "Success",
		},
		FAILED: UpgradePlanStatusPhase{
			value: "Failed",
		},
		SKIPPED: UpgradePlanStatusPhase{
			value: "Skipped",
		},
	}
}

func (c UpgradePlanStatusPhase) Value() string {
	return c.value
}

func (c UpgradePlanStatusPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradePlanStatusPhase) UnmarshalJSON(b []byte) error {
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
