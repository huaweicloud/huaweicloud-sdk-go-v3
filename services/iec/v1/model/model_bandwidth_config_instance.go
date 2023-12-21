package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BandwidthConfigInstance 带宽配置
type BandwidthConfigInstance struct {

	// 带宽类型，现支持WHOLE类型，即共享带宽，其他类型不支持。
	Sharetype BandwidthConfigInstanceSharetype `json:"sharetype"`

	// 带宽（Mbit/s）。
	Size *int32 `json:"size,omitempty"`

	Ids *[]string `json:"ids,omitempty"`

	BandwidthTypes *[]string `json:"bandwidth_types,omitempty"`
}

func (o BandwidthConfigInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthConfigInstance struct{}"
	}

	return strings.Join([]string{"BandwidthConfigInstance", string(data)}, " ")
}

type BandwidthConfigInstanceSharetype struct {
	value string
}

type BandwidthConfigInstanceSharetypeEnum struct {
	WHOLE BandwidthConfigInstanceSharetype
}

func GetBandwidthConfigInstanceSharetypeEnum() BandwidthConfigInstanceSharetypeEnum {
	return BandwidthConfigInstanceSharetypeEnum{
		WHOLE: BandwidthConfigInstanceSharetype{
			value: "WHOLE",
		},
	}
}

func (c BandwidthConfigInstanceSharetype) Value() string {
	return c.value
}

func (c BandwidthConfigInstanceSharetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthConfigInstanceSharetype) UnmarshalJSON(b []byte) error {
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
