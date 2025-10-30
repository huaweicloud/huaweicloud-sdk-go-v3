package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OfferingStatus 商品状态说明：   - TESTING - 测试   - ONSALE - 在售   - SUSPENDED - 已停售   - RETIREMENT - 产品退出
type OfferingStatus struct {
	value string
}

type OfferingStatusEnum struct {
	TESTING    OfferingStatus
	ONSALE     OfferingStatus
	SUSPENDED  OfferingStatus
	RETIREMENT OfferingStatus
}

func GetOfferingStatusEnum() OfferingStatusEnum {
	return OfferingStatusEnum{
		TESTING: OfferingStatus{
			value: "TESTING",
		},
		ONSALE: OfferingStatus{
			value: "ONSALE",
		},
		SUSPENDED: OfferingStatus{
			value: "SUSPENDED",
		},
		RETIREMENT: OfferingStatus{
			value: "RETIREMENT",
		},
	}
}

func (c OfferingStatus) Value() string {
	return c.value
}

func (c OfferingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OfferingStatus) UnmarshalJSON(b []byte) error {
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
