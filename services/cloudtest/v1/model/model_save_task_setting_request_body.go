package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SaveTaskSettingRequestBody struct {

	// 自定义时长
	MaxRunningTime *int32 `json:"max_running_time,omitempty"`

	// 是否长期运行
	IsLongRun *SaveTaskSettingRequestBodyIsLongRun `json:"is_long_run,omitempty"`
}

func (o SaveTaskSettingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTaskSettingRequestBody struct{}"
	}

	return strings.Join([]string{"SaveTaskSettingRequestBody", string(data)}, " ")
}

type SaveTaskSettingRequestBodyIsLongRun struct {
	value string
}

type SaveTaskSettingRequestBodyIsLongRunEnum struct {
	E_1 SaveTaskSettingRequestBodyIsLongRun
	E_0 SaveTaskSettingRequestBodyIsLongRun
}

func GetSaveTaskSettingRequestBodyIsLongRunEnum() SaveTaskSettingRequestBodyIsLongRunEnum {
	return SaveTaskSettingRequestBodyIsLongRunEnum{
		E_1: SaveTaskSettingRequestBodyIsLongRun{
			value: "1",
		},
		E_0: SaveTaskSettingRequestBodyIsLongRun{
			value: "0",
		},
	}
}

func (c SaveTaskSettingRequestBodyIsLongRun) Value() string {
	return c.value
}

func (c SaveTaskSettingRequestBodyIsLongRun) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SaveTaskSettingRequestBodyIsLongRun) UnmarshalJSON(b []byte) error {
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
