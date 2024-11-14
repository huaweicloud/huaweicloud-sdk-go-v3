package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TrainingAllocatedResource struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源计费类型。 * PERIODIC: 包周期 * ONE_TIME：一次性计费 > * 一次性计费包括：租户订购的一次性资源，SP管理员分配给租户的一次性资源。
	ChargeMode *TrainingAllocatedResourceChargeMode `json:"charge_mode,omitempty"`

	// 资源过期时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o TrainingAllocatedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingAllocatedResource struct{}"
	}

	return strings.Join([]string{"TrainingAllocatedResource", string(data)}, " ")
}

type TrainingAllocatedResourceChargeMode struct {
	value string
}

type TrainingAllocatedResourceChargeModeEnum struct {
	PERIODIC TrainingAllocatedResourceChargeMode
	ONE_TIME TrainingAllocatedResourceChargeMode
}

func GetTrainingAllocatedResourceChargeModeEnum() TrainingAllocatedResourceChargeModeEnum {
	return TrainingAllocatedResourceChargeModeEnum{
		PERIODIC: TrainingAllocatedResourceChargeMode{
			value: "PERIODIC",
		},
		ONE_TIME: TrainingAllocatedResourceChargeMode{
			value: "ONE_TIME",
		},
	}
}

func (c TrainingAllocatedResourceChargeMode) Value() string {
	return c.value
}

func (c TrainingAllocatedResourceChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrainingAllocatedResourceChargeMode) UnmarshalJSON(b []byte) error {
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
