package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UnlockRepositoryResponse Response Object
type UnlockRepositoryResponse struct {

	// 锁定状态 - true 已锁定 - false 未锁定
	Locked         *UnlockRepositoryResponseLocked `json:"locked,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o UnlockRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockRepositoryResponse struct{}"
	}

	return strings.Join([]string{"UnlockRepositoryResponse", string(data)}, " ")
}

type UnlockRepositoryResponseLocked struct {
	value string
}

type UnlockRepositoryResponseLockedEnum struct {
	TRUE  UnlockRepositoryResponseLocked
	FALSE UnlockRepositoryResponseLocked
}

func GetUnlockRepositoryResponseLockedEnum() UnlockRepositoryResponseLockedEnum {
	return UnlockRepositoryResponseLockedEnum{
		TRUE: UnlockRepositoryResponseLocked{
			value: "true",
		},
		FALSE: UnlockRepositoryResponseLocked{
			value: "false",
		},
	}
}

func (c UnlockRepositoryResponseLocked) Value() string {
	return c.value
}

func (c UnlockRepositoryResponseLocked) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnlockRepositoryResponseLocked) UnmarshalJSON(b []byte) error {
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
