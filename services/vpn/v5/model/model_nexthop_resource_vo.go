package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NexthopResourceVo 下一跳资源信息
type NexthopResourceVo struct {

	// 下一跳资源ID
	Id *string `json:"id,omitempty"`

	// 下一跳资源类型
	Type *NexthopResourceVoType `json:"type,omitempty"`
}

func (o NexthopResourceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NexthopResourceVo struct{}"
	}

	return strings.Join([]string{"NexthopResourceVo", string(data)}, " ")
}

type NexthopResourceVoType struct {
	value string
}

type NexthopResourceVoTypeEnum struct {
	VPN_CONNECTION NexthopResourceVoType
	VPN_GATEWAY    NexthopResourceVoType
	ER             NexthopResourceVoType
}

func GetNexthopResourceVoTypeEnum() NexthopResourceVoTypeEnum {
	return NexthopResourceVoTypeEnum{
		VPN_CONNECTION: NexthopResourceVoType{
			value: "vpn_connection",
		},
		VPN_GATEWAY: NexthopResourceVoType{
			value: "vpn_gateway",
		},
		ER: NexthopResourceVoType{
			value: "er",
		},
	}
}

func (c NexthopResourceVoType) Value() string {
	return c.value
}

func (c NexthopResourceVoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NexthopResourceVoType) UnmarshalJSON(b []byte) error {
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
