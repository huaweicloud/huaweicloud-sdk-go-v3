package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 可用分区结构体
type AvailableZones struct {

	// 可用区编码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 可用区端口号。
	Port *string `json:"port,omitempty" xml:"port"`

	// 可用区名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 可用区ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 分区上是否还有可用资源。 - true：还有资源。 - false：资源已售罄。
	ResourceAvailability *AvailableZonesResourceAvailability `json:"resource_availability,omitempty" xml:"resource_availability"`
}

func (o AvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZones struct{}"
	}

	return strings.Join([]string{"AvailableZones", string(data)}, " ")
}

type AvailableZonesResourceAvailability struct {
	value string
}

type AvailableZonesResourceAvailabilityEnum struct {
	TRUE  AvailableZonesResourceAvailability
	FALSE AvailableZonesResourceAvailability
}

func GetAvailableZonesResourceAvailabilityEnum() AvailableZonesResourceAvailabilityEnum {
	return AvailableZonesResourceAvailabilityEnum{
		TRUE: AvailableZonesResourceAvailability{
			value: "true",
		},
		FALSE: AvailableZonesResourceAvailability{
			value: "false",
		},
	}
}

func (c AvailableZonesResourceAvailability) Value() string {
	return c.value
}

func (c AvailableZonesResourceAvailability) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AvailableZonesResourceAvailability) UnmarshalJSON(b []byte) error {
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
