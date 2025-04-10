package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DisasterRecoverStopXlogKeepRequestBody struct {

	// 容灾类型。
	DisasterType DisasterRecoverStopXlogKeepRequestBodyDisasterType `json:"disaster_type"`
}

func (o DisasterRecoverStopXlogKeepRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoverStopXlogKeepRequestBody struct{}"
	}

	return strings.Join([]string{"DisasterRecoverStopXlogKeepRequestBody", string(data)}, " ")
}

type DisasterRecoverStopXlogKeepRequestBodyDisasterType struct {
	value string
}

type DisasterRecoverStopXlogKeepRequestBodyDisasterTypeEnum struct {
	STREAM DisasterRecoverStopXlogKeepRequestBodyDisasterType
}

func GetDisasterRecoverStopXlogKeepRequestBodyDisasterTypeEnum() DisasterRecoverStopXlogKeepRequestBodyDisasterTypeEnum {
	return DisasterRecoverStopXlogKeepRequestBodyDisasterTypeEnum{
		STREAM: DisasterRecoverStopXlogKeepRequestBodyDisasterType{
			value: "stream",
		},
	}
}

func (c DisasterRecoverStopXlogKeepRequestBodyDisasterType) Value() string {
	return c.value
}

func (c DisasterRecoverStopXlogKeepRequestBodyDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisasterRecoverStopXlogKeepRequestBodyDisasterType) UnmarshalJSON(b []byte) error {
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
