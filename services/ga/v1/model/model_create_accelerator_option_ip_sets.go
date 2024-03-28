package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAcceleratorOptionIpSets struct {

	// IP地址类型，取值：IPV4，IPV6
	IpType CreateAcceleratorOptionIpSetsIpType `json:"ip_type"`

	Area *Area `json:"area"`
}

func (o CreateAcceleratorOptionIpSets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceleratorOptionIpSets struct{}"
	}

	return strings.Join([]string{"CreateAcceleratorOptionIpSets", string(data)}, " ")
}

type CreateAcceleratorOptionIpSetsIpType struct {
	value string
}

type CreateAcceleratorOptionIpSetsIpTypeEnum struct {
	IPV4 CreateAcceleratorOptionIpSetsIpType
	IPV6 CreateAcceleratorOptionIpSetsIpType
}

func GetCreateAcceleratorOptionIpSetsIpTypeEnum() CreateAcceleratorOptionIpSetsIpTypeEnum {
	return CreateAcceleratorOptionIpSetsIpTypeEnum{
		IPV4: CreateAcceleratorOptionIpSetsIpType{
			value: "IPV4",
		},
		IPV6: CreateAcceleratorOptionIpSetsIpType{
			value: "IPV6",
		},
	}
}

func (c CreateAcceleratorOptionIpSetsIpType) Value() string {
	return c.value
}

func (c CreateAcceleratorOptionIpSetsIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAcceleratorOptionIpSetsIpType) UnmarshalJSON(b []byte) error {
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
