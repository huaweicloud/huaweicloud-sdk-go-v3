package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CbcOperationLock struct {

	// 限制操作场景： - TO_PERIOD_LOCK：按需转包周期场景锁，不允许进行删除、规格变更、按需转包周期等 - SPEC_CHG_LOCK：包周期规格变更场景锁，不允许进行删除、规格变更等
	LockScene *CbcOperationLockLockScene `json:"lock_scene,omitempty"`

	// 发起限制操作对象的标志
	LockSourceId *string `json:"lock_source_id,omitempty"`
}

func (o CbcOperationLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcOperationLock struct{}"
	}

	return strings.Join([]string{"CbcOperationLock", string(data)}, " ")
}

type CbcOperationLockLockScene struct {
	value string
}

type CbcOperationLockLockSceneEnum struct {
	TO_PERIOD_LOCK CbcOperationLockLockScene
	PEC_CHG_LOCK   CbcOperationLockLockScene
}

func GetCbcOperationLockLockSceneEnum() CbcOperationLockLockSceneEnum {
	return CbcOperationLockLockSceneEnum{
		TO_PERIOD_LOCK: CbcOperationLockLockScene{
			value: "TO_PERIOD_LOCK",
		},
		PEC_CHG_LOCK: CbcOperationLockLockScene{
			value: "PEC_CHG_LOCK",
		},
	}
}

func (c CbcOperationLockLockScene) Value() string {
	return c.value
}

func (c CbcOperationLockLockScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CbcOperationLockLockScene) UnmarshalJSON(b []byte) error {
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
