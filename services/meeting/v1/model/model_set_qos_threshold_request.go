package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SetQosThresholdRequest struct {

	// 阈值类型： * AUDIO：音频相关阈值。 * VIDEO：视频相关阈值。 * SCREEN：屏幕共享相关阈值。 * CPU：CPU相关阈值。
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

func (c SetQosThresholdRequestThresholdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetQosThresholdRequestThresholdType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
