package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlavorId 资源规格。
type FlavorId struct {
	value string
}

type FlavorIdEnum struct {
	MICRO_5_G0_5_C1_G     FlavorId
	SMALL_10_G1_0_C2_G    FlavorId
	STANDARD_30_G2_0_C4_G FlavorId
	LARGE_50_G4_0_C8_G    FlavorId
	XLARGE_100_G4_0_C16_G FlavorId
	CUSTOM_XGX_XX_X       FlavorId
}

func GetFlavorIdEnum() FlavorIdEnum {
	return FlavorIdEnum{
		MICRO_5_G0_5_C1_G: FlavorId{
			value: "MICRO-5G:0.5C:1G",
		},
		SMALL_10_G1_0_C2_G: FlavorId{
			value: "SMALL-10G:1.0C:2G",
		},
		STANDARD_30_G2_0_C4_G: FlavorId{
			value: "STANDARD-30G:2.0C:4G",
		},
		LARGE_50_G4_0_C8_G: FlavorId{
			value: "LARGE-50G:4.0C:8G",
		},
		XLARGE_100_G4_0_C16_G: FlavorId{
			value: "XLARGE-100G:4.0C:16G",
		},
		CUSTOM_XGX_XX_X: FlavorId{
			value: "CUSTOM-XG:X-X:X-X",
		},
	}
}

func (c FlavorId) Value() string {
	return c.value
}

func (c FlavorId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlavorId) UnmarshalJSON(b []byte) error {
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
