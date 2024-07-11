package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MoveAppGroupsRequestBody struct {

	// 分组id
	Id string `json:"id"`

	// 移动方向，1为上移，-1为下移
	Movement MoveAppGroupsRequestBodyMovement `json:"movement"`
}

func (o MoveAppGroupsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppGroupsRequestBody struct{}"
	}

	return strings.Join([]string{"MoveAppGroupsRequestBody", string(data)}, " ")
}

type MoveAppGroupsRequestBodyMovement struct {
	value int32
}

type MoveAppGroupsRequestBodyMovementEnum struct {
	E_1       MoveAppGroupsRequestBodyMovement
	E_MINUS_1 MoveAppGroupsRequestBodyMovement
}

func GetMoveAppGroupsRequestBodyMovementEnum() MoveAppGroupsRequestBodyMovementEnum {
	return MoveAppGroupsRequestBodyMovementEnum{
		E_1: MoveAppGroupsRequestBodyMovement{
			value: 1,
		}, E_MINUS_1: MoveAppGroupsRequestBodyMovement{
			value: -1,
		},
	}
}

func (c MoveAppGroupsRequestBodyMovement) Value() int32 {
	return c.value
}

func (c MoveAppGroupsRequestBodyMovement) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MoveAppGroupsRequestBodyMovement) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
