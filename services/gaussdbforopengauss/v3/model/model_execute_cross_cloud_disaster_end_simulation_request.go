package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterEndSimulationRequest Request Object
type ExecuteCrossCloudDisasterEndSimulationRequest struct {

	// 语言。
	XLanguage *ExecuteCrossCloudDisasterEndSimulationRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *DisasterRecoverStopSimulationRequestBody `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterEndSimulationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterEndSimulationRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterEndSimulationRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterEndSimulationRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterEndSimulationRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterEndSimulationRequestXLanguage
	EN_US ExecuteCrossCloudDisasterEndSimulationRequestXLanguage
}

func GetExecuteCrossCloudDisasterEndSimulationRequestXLanguageEnum() ExecuteCrossCloudDisasterEndSimulationRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterEndSimulationRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterEndSimulationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterEndSimulationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterEndSimulationRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterEndSimulationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterEndSimulationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
