package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScaleConfigurationDataTriggerMetadata struct {

	// 度量类型，固定值\"Utilization\"，表示使用率。
	Type *ScaleConfigurationDataTriggerMetadataType `json:"type,omitempty"`

	// 期望值。
	Value *string `json:"value,omitempty"`
}

func (o ScaleConfigurationDataTriggerMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleConfigurationDataTriggerMetadata struct{}"
	}

	return strings.Join([]string{"ScaleConfigurationDataTriggerMetadata", string(data)}, " ")
}

type ScaleConfigurationDataTriggerMetadataType struct {
	value string
}

type ScaleConfigurationDataTriggerMetadataTypeEnum struct {
	UTILIZATION ScaleConfigurationDataTriggerMetadataType
}

func GetScaleConfigurationDataTriggerMetadataTypeEnum() ScaleConfigurationDataTriggerMetadataTypeEnum {
	return ScaleConfigurationDataTriggerMetadataTypeEnum{
		UTILIZATION: ScaleConfigurationDataTriggerMetadataType{
			value: "Utilization",
		},
	}
}

func (c ScaleConfigurationDataTriggerMetadataType) Value() string {
	return c.value
}

func (c ScaleConfigurationDataTriggerMetadataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScaleConfigurationDataTriggerMetadataType) UnmarshalJSON(b []byte) error {
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
