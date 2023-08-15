package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Eip eip字段数据结构说明
type Eip struct {

	// 弹性公网IP地址类型。类型枚举值：5_bgp、5_sbgp详情请参见《虚拟私有云API参考》申请弹性公网IP章节的publicip字段说明。
	Iptype EipIptype `json:"iptype"`

	Bandwidth *BandWidth `json:"bandwidth"`

	Extendparam *ExtendParamEip `json:"extendparam"`
}

func (o Eip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Eip struct{}"
	}

	return strings.Join([]string{"Eip", string(data)}, " ")
}

type EipIptype struct {
	value string
}

type EipIptypeEnum struct {
	E_5_BGP  EipIptype
	E_5_SBGP EipIptype
}

func GetEipIptypeEnum() EipIptypeEnum {
	return EipIptypeEnum{
		E_5_BGP: EipIptype{
			value: "5_bgp",
		},
		E_5_SBGP: EipIptype{
			value: "5_sbgp",
		},
	}
}

func (c EipIptype) Value() string {
	return c.value
}

func (c EipIptype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipIptype) UnmarshalJSON(b []byte) error {
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
