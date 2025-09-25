package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchAspStatusRequestBody asp开关切换请求体。
type SwitchAspStatusRequestBody struct {

	// **参数解释**: 开关状态，不区分大小写。 **约束限制**: 不涉及。 **取值范围**: - ON：开启ASP。 - OFF：关闭ASP。  **默认取值**: 不涉及。
	Status SwitchAspStatusRequestBodyStatus `json:"status"`
}

func (o SwitchAspStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAspStatusRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchAspStatusRequestBody", string(data)}, " ")
}

type SwitchAspStatusRequestBodyStatus struct {
	value string
}

type SwitchAspStatusRequestBodyStatusEnum struct {
	ON  SwitchAspStatusRequestBodyStatus
	OFF SwitchAspStatusRequestBodyStatus
}

func GetSwitchAspStatusRequestBodyStatusEnum() SwitchAspStatusRequestBodyStatusEnum {
	return SwitchAspStatusRequestBodyStatusEnum{
		ON: SwitchAspStatusRequestBodyStatus{
			value: "ON",
		},
		OFF: SwitchAspStatusRequestBodyStatus{
			value: "OFF",
		},
	}
}

func (c SwitchAspStatusRequestBodyStatus) Value() string {
	return c.value
}

func (c SwitchAspStatusRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchAspStatusRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
