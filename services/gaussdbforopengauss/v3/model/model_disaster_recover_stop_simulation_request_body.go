package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DisasterRecoverStopSimulationRequestBody struct {

	// 容灾类型。
	DisasterType DisasterRecoverStopSimulationRequestBodyDisasterType `json:"disaster_type"`
}

func (o DisasterRecoverStopSimulationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoverStopSimulationRequestBody struct{}"
	}

	return strings.Join([]string{"DisasterRecoverStopSimulationRequestBody", string(data)}, " ")
}

type DisasterRecoverStopSimulationRequestBodyDisasterType struct {
	value string
}

type DisasterRecoverStopSimulationRequestBodyDisasterTypeEnum struct {
	STREAM DisasterRecoverStopSimulationRequestBodyDisasterType
}

func GetDisasterRecoverStopSimulationRequestBodyDisasterTypeEnum() DisasterRecoverStopSimulationRequestBodyDisasterTypeEnum {
	return DisasterRecoverStopSimulationRequestBodyDisasterTypeEnum{
		STREAM: DisasterRecoverStopSimulationRequestBodyDisasterType{
			value: "stream",
		},
	}
}

func (c DisasterRecoverStopSimulationRequestBodyDisasterType) Value() string {
	return c.value
}

func (c DisasterRecoverStopSimulationRequestBodyDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisasterRecoverStopSimulationRequestBodyDisasterType) UnmarshalJSON(b []byte) error {
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
