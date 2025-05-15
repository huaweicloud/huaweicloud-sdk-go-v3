package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CrossRegionTypeEnum 跨地域类型。 - intra-region （同地域） - inter-region （跨地域）
type CrossRegionTypeEnum struct {
	value string
}

type CrossRegionTypeEnumEnum struct {
	INTRA_REGION CrossRegionTypeEnum
	INTER_REGION CrossRegionTypeEnum
}

func GetCrossRegionTypeEnumEnum() CrossRegionTypeEnumEnum {
	return CrossRegionTypeEnumEnum{
		INTRA_REGION: CrossRegionTypeEnum{
			value: "intra-region",
		},
		INTER_REGION: CrossRegionTypeEnum{
			value: "inter-region",
		},
	}
}

func (c CrossRegionTypeEnum) Value() string {
	return c.value
}

func (c CrossRegionTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CrossRegionTypeEnum) UnmarshalJSON(b []byte) error {
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
