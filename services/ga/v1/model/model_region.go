package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Region 区域。
type Region struct {

	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`

	// 区域所属地区，取值： - OUTOFCM： 中国大陆以外 - CM：中国大陆
	Area *string `json:"area,omitempty"`

	// 区域支持的终端节点类型。取值： EIP：弹性公网IP
	SupportedEndpointTypes *[]RegionSupportedEndpointTypes `json:"supported_endpoint_types,omitempty"`
}

func (o Region) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Region struct{}"
	}

	return strings.Join([]string{"Region", string(data)}, " ")
}

type RegionSupportedEndpointTypes struct {
	value string
}

type RegionSupportedEndpointTypesEnum struct {
	EIP RegionSupportedEndpointTypes
}

func GetRegionSupportedEndpointTypesEnum() RegionSupportedEndpointTypesEnum {
	return RegionSupportedEndpointTypesEnum{
		EIP: RegionSupportedEndpointTypes{
			value: "EIP",
		},
	}
}

func (c RegionSupportedEndpointTypes) Value() string {
	return c.value
}

func (c RegionSupportedEndpointTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegionSupportedEndpointTypes) UnmarshalJSON(b []byte) error {
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
