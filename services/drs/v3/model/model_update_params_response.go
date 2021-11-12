package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateParamsResponse struct {
	// 修改参数是否成功

	Success *bool `json:"success,omitempty"`
	// 是否需要重启

	ShouldRestart *UpdateParamsResponseShouldRestart `json:"should_restart,omitempty"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParamsResponse struct{}"
	}

	return strings.Join([]string{"UpdateParamsResponse", string(data)}, " ")
}

type UpdateParamsResponseShouldRestart struct {
	value string
}

type UpdateParamsResponseShouldRestartEnum struct {
	TRUE  UpdateParamsResponseShouldRestart
	FALSE UpdateParamsResponseShouldRestart
}

func GetUpdateParamsResponseShouldRestartEnum() UpdateParamsResponseShouldRestartEnum {
	return UpdateParamsResponseShouldRestartEnum{
		TRUE: UpdateParamsResponseShouldRestart{
			value: "true",
		},
		FALSE: UpdateParamsResponseShouldRestart{
			value: "false",
		},
	}
}

func (c UpdateParamsResponseShouldRestart) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParamsResponseShouldRestart) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
