package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FeatureStatus 功能状态。
type FeatureStatus struct {
	value string
}

type FeatureStatusEnum struct {
	ON      FeatureStatus
	OFF     FeatureStatus
	PENDING FeatureStatus
}

func GetFeatureStatusEnum() FeatureStatusEnum {
	return FeatureStatusEnum{
		ON: FeatureStatus{
			value: "on",
		},
		OFF: FeatureStatus{
			value: "off",
		},
		PENDING: FeatureStatus{
			value: "pending",
		},
	}
}

func (c FeatureStatus) Value() string {
	return c.value
}

func (c FeatureStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FeatureStatus) UnmarshalJSON(b []byte) error {
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
