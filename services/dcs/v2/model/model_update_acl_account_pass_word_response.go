package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAclAccountPassWordResponse Response Object
type UpdateAclAccountPassWordResponse struct {

	// 锁定时间。验证失败时和锁定时该参数返回不为null
	LockTime *string `json:"lock_time,omitempty"`

	// 密码修改结果： - 成功：success； - 密码验证失败：passwordFailed； - 已锁定：locked； - 失败：failed。
	Result *UpdateAclAccountPassWordResponseResult `json:"result,omitempty"`

	// 锁定剩余时间。锁定时该参数返回不为null
	LockTimeLeft *string `json:"lock_time_left,omitempty"`

	// 密码验证剩余次数。验证失败时该参数返回不为null
	RetryTimesLeft *string `json:"retry_times_left,omitempty"`

	// 修改结果。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAclAccountPassWordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountPassWordResponse struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountPassWordResponse", string(data)}, " ")
}

type UpdateAclAccountPassWordResponseResult struct {
	value string
}

type UpdateAclAccountPassWordResponseResultEnum struct {
	SUCCESS         UpdateAclAccountPassWordResponseResult
	PASSWORD_FAILED UpdateAclAccountPassWordResponseResult
	LOCKED          UpdateAclAccountPassWordResponseResult
	FAILED          UpdateAclAccountPassWordResponseResult
}

func GetUpdateAclAccountPassWordResponseResultEnum() UpdateAclAccountPassWordResponseResultEnum {
	return UpdateAclAccountPassWordResponseResultEnum{
		SUCCESS: UpdateAclAccountPassWordResponseResult{
			value: "success",
		},
		PASSWORD_FAILED: UpdateAclAccountPassWordResponseResult{
			value: "passwordFailed",
		},
		LOCKED: UpdateAclAccountPassWordResponseResult{
			value: "locked",
		},
		FAILED: UpdateAclAccountPassWordResponseResult{
			value: "failed",
		},
	}
}

func (c UpdateAclAccountPassWordResponseResult) Value() string {
	return c.value
}

func (c UpdateAclAccountPassWordResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAclAccountPassWordResponseResult) UnmarshalJSON(b []byte) error {
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
