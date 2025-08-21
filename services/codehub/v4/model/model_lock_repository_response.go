package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LockRepositoryResponse Response Object
type LockRepositoryResponse struct {

	// 锁定状态 - true 已锁定 - false 未锁定
	Locked         *LockRepositoryResponseLocked `json:"locked,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o LockRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockRepositoryResponse struct{}"
	}

	return strings.Join([]string{"LockRepositoryResponse", string(data)}, " ")
}

type LockRepositoryResponseLocked struct {
	value string
}

type LockRepositoryResponseLockedEnum struct {
	TRUE  LockRepositoryResponseLocked
	FALSE LockRepositoryResponseLocked
}

func GetLockRepositoryResponseLockedEnum() LockRepositoryResponseLockedEnum {
	return LockRepositoryResponseLockedEnum{
		TRUE: LockRepositoryResponseLocked{
			value: "true",
		},
		FALSE: LockRepositoryResponseLocked{
			value: "false",
		},
	}
}

func (c LockRepositoryResponseLocked) Value() string {
	return c.value
}

func (c LockRepositoryResponseLocked) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LockRepositoryResponseLocked) UnmarshalJSON(b []byte) error {
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
