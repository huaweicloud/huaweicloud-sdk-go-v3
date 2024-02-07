package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NexthopDict 下一跳对象
type NexthopDict struct {

	// 功能说明：下一跳的实际地址  约束：同一virtualnexthop中的nexthops列表，不同下一跳之间的地址不允许重复
	IpAddress *string `json:"ip_address,omitempty"`

	// 功能说明：标识主备模式，与virtualnexthop的forward_mode配合使用  取值范围：'ACTIVE'主、'STANDBY'备
	Mode *NexthopDictMode `json:"mode,omitempty"`
}

func (o NexthopDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NexthopDict struct{}"
	}

	return strings.Join([]string{"NexthopDict", string(data)}, " ")
}

type NexthopDictMode struct {
	value string
}

type NexthopDictModeEnum struct {
	ACTIVE  NexthopDictMode
	STANDBY NexthopDictMode
}

func GetNexthopDictModeEnum() NexthopDictModeEnum {
	return NexthopDictModeEnum{
		ACTIVE: NexthopDictMode{
			value: "ACTIVE",
		},
		STANDBY: NexthopDictMode{
			value: "STANDBY",
		},
	}
}

func (c NexthopDictMode) Value() string {
	return c.value
}

func (c NexthopDictMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NexthopDictMode) UnmarshalJSON(b []byte) error {
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
