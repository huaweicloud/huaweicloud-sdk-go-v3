package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterStartSimulationRequest Request Object
type ExecuteCrossCloudDisasterStartSimulationRequest struct {

	// 语言。
	XLanguage *ExecuteCrossCloudDisasterStartSimulationRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *DisasterRecoverStartSimulationRequestBody `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterStartSimulationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterStartSimulationRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterStartSimulationRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterStartSimulationRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterStartSimulationRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterStartSimulationRequestXLanguage
	EN_US ExecuteCrossCloudDisasterStartSimulationRequestXLanguage
}

func GetExecuteCrossCloudDisasterStartSimulationRequestXLanguageEnum() ExecuteCrossCloudDisasterStartSimulationRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterStartSimulationRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterStartSimulationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterStartSimulationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterStartSimulationRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterStartSimulationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterStartSimulationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
