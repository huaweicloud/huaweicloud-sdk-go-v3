package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserAuth struct {

	// 用户id，需要从IAM服务获取
	UserId string `json:"user_id"`

	// 用户名，需要从IAM服务获取
	UserName string `json:"user_name"`

	// 用户权限，7表示管理权限，3表示编辑权限，1表示读取权限
	Auth UserAuthAuth `json:"auth"`
}

func (o UserAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAuth struct{}"
	}

	return strings.Join([]string{"UserAuth", string(data)}, " ")
}

type UserAuthAuth struct {
	value int64
}

type UserAuthAuthEnum struct {
	E_7 UserAuthAuth
	E_3 UserAuthAuth
	E_1 UserAuthAuth
}

func GetUserAuthAuthEnum() UserAuthAuthEnum {
	return UserAuthAuthEnum{
		E_7: UserAuthAuth{
			value: 7,
		}, E_3: UserAuthAuth{
			value: 3,
		}, E_1: UserAuthAuth{
			value: 1,
		},
	}
}

func (c UserAuthAuth) Value() int64 {
	return c.value
}

func (c UserAuthAuth) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserAuthAuth) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int64")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int64); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int64 error")
	}
}
