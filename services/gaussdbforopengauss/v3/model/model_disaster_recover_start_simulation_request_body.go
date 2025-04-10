package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DisasterRecoverStartSimulationRequestBody struct {

	// 主实例下发容灾演练时必填，备实例下发容灾演练是不必填   - 日志保留空间占可使用剩余磁盘容量的比例, 可设置范围为1-100。
	XlogKeepRatio *int32 `json:"xlog_keep_ratio,omitempty"`

	// 容灾类型。
	DisasterType DisasterRecoverStartSimulationRequestBodyDisasterType `json:"disaster_type"`
}

func (o DisasterRecoverStartSimulationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoverStartSimulationRequestBody struct{}"
	}

	return strings.Join([]string{"DisasterRecoverStartSimulationRequestBody", string(data)}, " ")
}

type DisasterRecoverStartSimulationRequestBodyDisasterType struct {
	value string
}

type DisasterRecoverStartSimulationRequestBodyDisasterTypeEnum struct {
	STREAM DisasterRecoverStartSimulationRequestBodyDisasterType
}

func GetDisasterRecoverStartSimulationRequestBodyDisasterTypeEnum() DisasterRecoverStartSimulationRequestBodyDisasterTypeEnum {
	return DisasterRecoverStartSimulationRequestBodyDisasterTypeEnum{
		STREAM: DisasterRecoverStartSimulationRequestBodyDisasterType{
			value: "stream",
		},
	}
}

func (c DisasterRecoverStartSimulationRequestBodyDisasterType) Value() string {
	return c.value
}

func (c DisasterRecoverStartSimulationRequestBodyDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisasterRecoverStartSimulationRequestBodyDisasterType) UnmarshalJSON(b []byte) error {
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
