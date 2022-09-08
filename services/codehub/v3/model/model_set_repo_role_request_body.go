package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SetRepoRoleRequestBody struct {

	// 设置仓库的成员权限，取值范围：20 -> 只读成员 30->普通成员，40->管理员
	Role SetRepoRoleRequestBodyRole `json:"role"`
}

func (o SetRepoRoleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRepoRoleRequestBody struct{}"
	}

	return strings.Join([]string{"SetRepoRoleRequestBody", string(data)}, " ")
}

type SetRepoRoleRequestBodyRole struct {
	value int32
}

type SetRepoRoleRequestBodyRoleEnum struct {
	E_20 SetRepoRoleRequestBodyRole
	E_30 SetRepoRoleRequestBodyRole
	E_40 SetRepoRoleRequestBodyRole
}

func GetSetRepoRoleRequestBodyRoleEnum() SetRepoRoleRequestBodyRoleEnum {
	return SetRepoRoleRequestBodyRoleEnum{
		E_20: SetRepoRoleRequestBodyRole{
			value: 20,
		}, E_30: SetRepoRoleRequestBodyRole{
			value: 30,
		}, E_40: SetRepoRoleRequestBodyRole{
			value: 40,
		},
	}
}

func (c SetRepoRoleRequestBodyRole) Value() int32 {
	return c.value
}

func (c SetRepoRoleRequestBodyRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetRepoRoleRequestBodyRole) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
