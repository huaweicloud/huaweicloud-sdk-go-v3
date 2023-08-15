package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetQosThresholdRequest Request Object
type SetQosThresholdRequest struct {

	// 阈值类型： * AUDIO：音频相关阈值 * VIDEO：视频相关阈值 * SCREEN：屏幕共享相关阈值 * CPU：CPU相关阈值
	ThresholdType SetQosThresholdRequestThresholdType `json:"thresholdType"`

	Body *SetQosThresholdReq `json:"body,omitempty"`
}

func (o SetQosThresholdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetQosThresholdRequest struct{}"
	}

	return strings.Join([]string{"SetQosThresholdRequest", string(data)}, " ")
}

type SetQosThresholdRequestThresholdType struct {
	value string
}

type SetQosThresholdRequestThresholdTypeEnum struct {
	AUDIO  SetQosThresholdRequestThresholdType
	VIDEO  SetQosThresholdRequestThresholdType
	SCREEN SetQosThresholdRequestThresholdType
	CPU    SetQosThresholdRequestThresholdType
}

func GetSetQosThresholdRequestThresholdTypeEnum() SetQosThresholdRequestThresholdTypeEnum {
	return SetQosThresholdRequestThresholdTypeEnum{
		AUDIO: SetQosThresholdRequestThresholdType{
			value: "AUDIO",
		},
		VIDEO: SetQosThresholdRequestThresholdType{
			value: "VIDEO",
		},
		SCREEN: SetQosThresholdRequestThresholdType{
			value: "SCREEN",
		},
		CPU: SetQosThresholdRequestThresholdType{
			value: "CPU",
		},
	}
}

func (c SetQosThresholdRequestThresholdType) Value() string {
	return c.value
}

func (c SetQosThresholdRequestThresholdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetQosThresholdRequestThresholdType) UnmarshalJSON(b []byte) error {
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
