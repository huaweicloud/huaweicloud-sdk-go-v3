package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePasswordResponse Response Object
type UpdatePasswordResponse struct {

	// 锁定时间。验证失败时和锁定时该参数返回不为null
	LockTime *string `json:"lock_time,omitempty"`

	// 密码修改结果： - 成功：success； - 密码验证失败：passwordFailed； - 已锁定：locked； - 失败：failed。
	Result *UpdatePasswordResponseResult `json:"result,omitempty"`

	// 锁定剩余时间。锁定时该参数返回不为null
	LockTimeLeft *string `json:"lock_time_left,omitempty"`

	// 密码验证剩余次数。验证失败时该参数返回不为null
	RetryTimesLeft *string `json:"retry_times_left,omitempty"`

	// 修改结果。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdatePasswordResponse", string(data)}, " ")
}

type UpdatePasswordResponseResult struct {
	value string
}

type UpdatePasswordResponseResultEnum struct {
	SUCCESS         UpdatePasswordResponseResult
	PASSWORD_FAILED UpdatePasswordResponseResult
	LOCKED          UpdatePasswordResponseResult
	FAILED          UpdatePasswordResponseResult
}

func GetUpdatePasswordResponseResultEnum() UpdatePasswordResponseResultEnum {
	return UpdatePasswordResponseResultEnum{
		SUCCESS: UpdatePasswordResponseResult{
			value: "success",
		},
		PASSWORD_FAILED: UpdatePasswordResponseResult{
			value: "passwordFailed",
		},
		LOCKED: UpdatePasswordResponseResult{
			value: "locked",
		},
		FAILED: UpdatePasswordResponseResult{
			value: "failed",
		},
	}
}

func (c UpdatePasswordResponseResult) Value() string {
	return c.value
}

func (c UpdatePasswordResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePasswordResponseResult) UnmarshalJSON(b []byte) error {
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
