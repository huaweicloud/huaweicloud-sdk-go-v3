package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateRecycleBinOption struct {

	// 回收站开启开关
	Switch UpdateRecycleBinOptionSwitch `json:"switch"`
}

func (o UpdateRecycleBinOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinOption struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinOption", string(data)}, " ")
}

type UpdateRecycleBinOptionSwitch struct {
	value string
}

type UpdateRecycleBinOptionSwitchEnum struct {
	ON  UpdateRecycleBinOptionSwitch
	OFF UpdateRecycleBinOptionSwitch
}

func GetUpdateRecycleBinOptionSwitchEnum() UpdateRecycleBinOptionSwitchEnum {
	return UpdateRecycleBinOptionSwitchEnum{
		ON: UpdateRecycleBinOptionSwitch{
			value: "on",
		},
		OFF: UpdateRecycleBinOptionSwitch{
			value: "off",
		},
	}
}

func (c UpdateRecycleBinOptionSwitch) Value() string {
	return c.value
}

func (c UpdateRecycleBinOptionSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRecycleBinOptionSwitch) UnmarshalJSON(b []byte) error {
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
