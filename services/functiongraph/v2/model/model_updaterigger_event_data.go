package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateriggerEventData struct {

	// 串行处理数据
	IsSerial *bool `json:"is_serial,omitempty"`

	// 最大字节数
	MaxFetchBytes *int32 `json:"max_fetch_bytes,omitempty"`

	// 拉取周期
	PollingInterval *int32 `json:"polling_interval,omitempty"`

	// 拉取周期单位
	PollingUnit *UpdateriggerEventDataPollingUnit `json:"polling_unit,omitempty"`
}

func (o UpdateriggerEventData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateriggerEventData struct{}"
	}

	return strings.Join([]string{"UpdateriggerEventData", string(data)}, " ")
}

type UpdateriggerEventDataPollingUnit struct {
	value string
}

type UpdateriggerEventDataPollingUnitEnum struct {
	MS UpdateriggerEventDataPollingUnit
	S  UpdateriggerEventDataPollingUnit
}

func GetUpdateriggerEventDataPollingUnitEnum() UpdateriggerEventDataPollingUnitEnum {
	return UpdateriggerEventDataPollingUnitEnum{
		MS: UpdateriggerEventDataPollingUnit{
			value: "ms",
		},
		S: UpdateriggerEventDataPollingUnit{
			value: "s",
		},
	}
}

func (c UpdateriggerEventDataPollingUnit) Value() string {
	return c.value
}

func (c UpdateriggerEventDataPollingUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateriggerEventDataPollingUnit) UnmarshalJSON(b []byte) error {
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
