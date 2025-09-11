package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwitchWdrSnapshotRequestBody struct {

	// **参数解释**： 开关状态。 **约束限制**： 不涉及。 **取值范围**： - ON：WDR快照开启。 - OFF：WDR快照关闭。  **默认取值**： 不涉及。
	Status SwitchWdrSnapshotRequestBodyStatus `json:"status"`
}

func (o SwitchWdrSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchWdrSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchWdrSnapshotRequestBody", string(data)}, " ")
}

type SwitchWdrSnapshotRequestBodyStatus struct {
	value string
}

type SwitchWdrSnapshotRequestBodyStatusEnum struct {
	ON  SwitchWdrSnapshotRequestBodyStatus
	OFF SwitchWdrSnapshotRequestBodyStatus
}

func GetSwitchWdrSnapshotRequestBodyStatusEnum() SwitchWdrSnapshotRequestBodyStatusEnum {
	return SwitchWdrSnapshotRequestBodyStatusEnum{
		ON: SwitchWdrSnapshotRequestBodyStatus{
			value: "ON",
		},
		OFF: SwitchWdrSnapshotRequestBodyStatus{
			value: "OFF",
		},
	}
}

func (c SwitchWdrSnapshotRequestBodyStatus) Value() string {
	return c.value
}

func (c SwitchWdrSnapshotRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchWdrSnapshotRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
