package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConnectionInfoFlavor struct {

	// 规格名称
	Name *string `json:"name,omitempty"`

	// 并发规格类型
	ConcurrencyType *ConnectionInfoFlavorConcurrencyType `json:"concurrency_type,omitempty"`

	// 并发数
	Concurrency *int32 `json:"concurrency,omitempty"`

	// 带宽类型
	BandwidthType *ConnectionInfoFlavorBandwidthType `json:"bandwidth_type,omitempty"`
}

func (o ConnectionInfoFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionInfoFlavor struct{}"
	}

	return strings.Join([]string{"ConnectionInfoFlavor", string(data)}, " ")
}

type ConnectionInfoFlavorConcurrencyType struct {
	value string
}

type ConnectionInfoFlavorConcurrencyTypeEnum struct {
	SHARED    ConnectionInfoFlavorConcurrencyType
	EXCLUSIVE ConnectionInfoFlavorConcurrencyType
}

func GetConnectionInfoFlavorConcurrencyTypeEnum() ConnectionInfoFlavorConcurrencyTypeEnum {
	return ConnectionInfoFlavorConcurrencyTypeEnum{
		SHARED: ConnectionInfoFlavorConcurrencyType{
			value: "shared",
		},
		EXCLUSIVE: ConnectionInfoFlavorConcurrencyType{
			value: "exclusive",
		},
	}
}

func (c ConnectionInfoFlavorConcurrencyType) Value() string {
	return c.value
}

func (c ConnectionInfoFlavorConcurrencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionInfoFlavorConcurrencyType) UnmarshalJSON(b []byte) error {
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

type ConnectionInfoFlavorBandwidthType struct {
	value string
}

type ConnectionInfoFlavorBandwidthTypeEnum struct {
	SHARED    ConnectionInfoFlavorBandwidthType
	EXCLUSIVE ConnectionInfoFlavorBandwidthType
}

func GetConnectionInfoFlavorBandwidthTypeEnum() ConnectionInfoFlavorBandwidthTypeEnum {
	return ConnectionInfoFlavorBandwidthTypeEnum{
		SHARED: ConnectionInfoFlavorBandwidthType{
			value: "shared",
		},
		EXCLUSIVE: ConnectionInfoFlavorBandwidthType{
			value: "exclusive",
		},
	}
}

func (c ConnectionInfoFlavorBandwidthType) Value() string {
	return c.value
}

func (c ConnectionInfoFlavorBandwidthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionInfoFlavorBandwidthType) UnmarshalJSON(b []byte) error {
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
