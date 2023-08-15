package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Target struct {

	// 对象ID。
	TargetId string `json:"target_id"`

	// 对象名称。
	TargetName string `json:"target_name"`

	// 对象类型。 - USER：表示用户。   target_id：为用户ID。   target_name：为用户name。 - USERGROUP：表示用户组。   target_id：为用户组ID。   target_name：为用户组name。 - APPGROUP：应用组。   target_id：应用组id。   target_name：应用组名称。 - OU：组织单元。   target_id：OUID。   target_name：OU名称。 - ALL：表示所有。   target_id：default-apply-all-targets。   target_name：All-Targets。
	TargetType TargetTargetType `json:"target_type"`
}

func (o Target) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Target struct{}"
	}

	return strings.Join([]string{"Target", string(data)}, " ")
}

type TargetTargetType struct {
	value string
}

type TargetTargetTypeEnum struct {
	USER      TargetTargetType
	USERGROUP TargetTargetType
	APPGROUP  TargetTargetType
	OU        TargetTargetType
	ALL       TargetTargetType
}

func GetTargetTargetTypeEnum() TargetTargetTypeEnum {
	return TargetTargetTypeEnum{
		USER: TargetTargetType{
			value: "USER",
		},
		USERGROUP: TargetTargetType{
			value: "USERGROUP",
		},
		APPGROUP: TargetTargetType{
			value: "APPGROUP",
		},
		OU: TargetTargetType{
			value: "OU",
		},
		ALL: TargetTargetType{
			value: "ALL",
		},
	}
}

func (c TargetTargetType) Value() string {
	return c.value
}

func (c TargetTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetTargetType) UnmarshalJSON(b []byte) error {
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
