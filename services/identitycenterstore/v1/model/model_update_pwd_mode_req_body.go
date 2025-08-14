package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePwdModeReqBody The request body of UpdatePassword.
type UpdatePwdModeReqBody struct {

	// 重置密码方式：一次性密码/邮箱
	Mode UpdatePwdModeReqBodyMode `json:"mode"`
}

func (o UpdatePwdModeReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePwdModeReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePwdModeReqBody", string(data)}, " ")
}

type UpdatePwdModeReqBodyMode struct {
	value string
}

type UpdatePwdModeReqBodyModeEnum struct {
	OTP   UpdatePwdModeReqBodyMode
	EMAIL UpdatePwdModeReqBodyMode
}

func GetUpdatePwdModeReqBodyModeEnum() UpdatePwdModeReqBodyModeEnum {
	return UpdatePwdModeReqBodyModeEnum{
		OTP: UpdatePwdModeReqBodyMode{
			value: "OTP",
		},
		EMAIL: UpdatePwdModeReqBodyMode{
			value: "EMAIL",
		},
	}
}

func (c UpdatePwdModeReqBodyMode) Value() string {
	return c.value
}

func (c UpdatePwdModeReqBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePwdModeReqBodyMode) UnmarshalJSON(b []byte) error {
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
