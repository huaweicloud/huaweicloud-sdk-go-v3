package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EipSpecBandwidth struct {

	// 带宽大小
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 带宽类型
	Sharetype *EipSpecBandwidthSharetype `json:"sharetype,omitempty" xml:"sharetype"`
}

func (o EipSpecBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipSpecBandwidth struct{}"
	}

	return strings.Join([]string{"EipSpecBandwidth", string(data)}, " ")
}

type EipSpecBandwidthSharetype struct {
	value string
}

type EipSpecBandwidthSharetypeEnum struct {
	PER   EipSpecBandwidthSharetype
	WHOLE EipSpecBandwidthSharetype
}

func GetEipSpecBandwidthSharetypeEnum() EipSpecBandwidthSharetypeEnum {
	return EipSpecBandwidthSharetypeEnum{
		PER: EipSpecBandwidthSharetype{
			value: "PER",
		},
		WHOLE: EipSpecBandwidthSharetype{
			value: "WHOLE",
		},
	}
}

func (c EipSpecBandwidthSharetype) Value() string {
	return c.value
}

func (c EipSpecBandwidthSharetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipSpecBandwidthSharetype) UnmarshalJSON(b []byte) error {
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
