package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableLockStatus **参数解释**: 表锁状态 - LOCKED 已锁 - UNLOCKED 未锁  **约束限制** 不涉及 **取值范围**: - LOCKED - UNLOCKED  **默认值** 不涉及
type TableLockStatus struct {
	value string
}

type TableLockStatusEnum struct {
	LOCKED   TableLockStatus
	UNLOCKED TableLockStatus
}

func GetTableLockStatusEnum() TableLockStatusEnum {
	return TableLockStatusEnum{
		LOCKED: TableLockStatus{
			value: "LOCKED",
		},
		UNLOCKED: TableLockStatus{
			value: "UNLOCKED",
		},
	}
}

func (c TableLockStatus) Value() string {
	return c.value
}

func (c TableLockStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableLockStatus) UnmarshalJSON(b []byte) error {
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
