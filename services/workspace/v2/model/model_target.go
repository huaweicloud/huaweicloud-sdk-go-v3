package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Target struct {

	// 对象ID。
	TargetId *string `json:"target_id,omitempty"`

	// 对象名称，长度不能超过55个字符。
	TargetName *string `json:"target_name,omitempty"`

	// 对象类型。 - INSTANCE：表示桌面。   target_id：为桌面的SID。   target_name：为桌面name。 - USER：表示用户。   target_id：为用户ID。   target_name：为用户name。 - USERGROUP：表示用户组。   target_id：为用户组ID。   target_name：为用户组name。 - CLIENTIP：终端IP地址。   target_id：终端IP地址。   target_name：终端IP地址。 - OU：组织单元。   target_id：OUID。   target_name：OU名称。 - DESKTOPSPOOL：表示桌面池。   target_id：为桌面池的ID。   target_name：为桌面池name。 - ALL：表示所有桌面。   target_id：default-apply-all-targets。   target_name：All-Targets。 - DESKTOP_TAG：表示桌面标签。   target_id：标签的key|标签的value。   target_name：标签的key|标签的value。
	TargetType *TargetTargetType `json:"target_type,omitempty"`
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
	INSTANCE     TargetTargetType
	DESKTOPSPOOL TargetTargetType
	USER         TargetTargetType
	CLIENTIP     TargetTargetType
	OU           TargetTargetType
	USERGROUP    TargetTargetType
	ALL          TargetTargetType
	DESKTOP_TAG  TargetTargetType
}

func GetTargetTargetTypeEnum() TargetTargetTypeEnum {
	return TargetTargetTypeEnum{
		INSTANCE: TargetTargetType{
			value: "INSTANCE",
		},
		DESKTOPSPOOL: TargetTargetType{
			value: "DESKTOPSPOOL",
		},
		USER: TargetTargetType{
			value: "USER",
		},
		CLIENTIP: TargetTargetType{
			value: "CLIENTIP",
		},
		OU: TargetTargetType{
			value: "OU",
		},
		USERGROUP: TargetTargetType{
			value: "USERGROUP",
		},
		ALL: TargetTargetType{
			value: "ALL",
		},
		DESKTOP_TAG: TargetTargetType{
			value: "DESKTOP_TAG",
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
