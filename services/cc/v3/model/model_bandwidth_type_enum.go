package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BandwidthTypeEnum 带宽类型包括： - BandwidthPackage (按带宽计费，需要绑定全域互联带宽，并指定分配带宽大小) - TestBandwidth (不收费的测试带宽，仅保留最小带宽，用于测试跨地域连通性）
type BandwidthTypeEnum struct {
	value string
}

type BandwidthTypeEnumEnum struct {
	BANDWIDTH_PACKAGE BandwidthTypeEnum
	TEST_BANDWIDTH    BandwidthTypeEnum
}

func GetBandwidthTypeEnumEnum() BandwidthTypeEnumEnum {
	return BandwidthTypeEnumEnum{
		BANDWIDTH_PACKAGE: BandwidthTypeEnum{
			value: "BandwidthPackage",
		},
		TEST_BANDWIDTH: BandwidthTypeEnum{
			value: "TestBandwidth",
		},
	}
}

func (c BandwidthTypeEnum) Value() string {
	return c.value
}

func (c BandwidthTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthTypeEnum) UnmarshalJSON(b []byte) error {
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
