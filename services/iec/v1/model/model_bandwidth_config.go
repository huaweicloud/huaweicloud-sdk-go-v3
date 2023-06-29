package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BandwidthConfig 带宽配置
type BandwidthConfig struct {

	// 带宽类型，现支持WHOLE类型，即共享带宽，其他类型不支持。
	Sharetype BandwidthConfigSharetype `json:"sharetype"`

	// 带宽（Mbit/s）。
	Size *int32 `json:"size,omitempty"`
}

func (o BandwidthConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthConfig struct{}"
	}

	return strings.Join([]string{"BandwidthConfig", string(data)}, " ")
}

type BandwidthConfigSharetype struct {
	value string
}

type BandwidthConfigSharetypeEnum struct {
	WHOLE BandwidthConfigSharetype
}

func GetBandwidthConfigSharetypeEnum() BandwidthConfigSharetypeEnum {
	return BandwidthConfigSharetypeEnum{
		WHOLE: BandwidthConfigSharetype{
			value: "WHOLE",
		},
	}
}

func (c BandwidthConfigSharetype) Value() string {
	return c.value
}

func (c BandwidthConfigSharetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthConfigSharetype) UnmarshalJSON(b []byte) error {
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
