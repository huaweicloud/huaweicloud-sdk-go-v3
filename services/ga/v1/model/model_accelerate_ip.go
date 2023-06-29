package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccelerateIp accelerate Ip
type AccelerateIp struct {

	// IP地址类型。
	IpType AccelerateIpIpType `json:"ip_type"`

	// IP地址。
	IpAddress *string `json:"ip_address,omitempty"`

	Area *Area `json:"area,omitempty"`
}

func (o AccelerateIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccelerateIp struct{}"
	}

	return strings.Join([]string{"AccelerateIp", string(data)}, " ")
}

type AccelerateIpIpType struct {
	value string
}

type AccelerateIpIpTypeEnum struct {
	IPV4 AccelerateIpIpType
}

func GetAccelerateIpIpTypeEnum() AccelerateIpIpTypeEnum {
	return AccelerateIpIpTypeEnum{
		IPV4: AccelerateIpIpType{
			value: "IPV4",
		},
	}
}

func (c AccelerateIpIpType) Value() string {
	return c.value
}

func (c AccelerateIpIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccelerateIpIpType) UnmarshalJSON(b []byte) error {
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
