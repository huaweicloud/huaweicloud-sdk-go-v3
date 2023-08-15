package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowQosThresholdRequest Request Object
type ShowQosThresholdRequest struct {

	// 阈值类型。 * AUDIO：音频告警阈值 * VIDEO：视频告警阈值 * SCREEN：屏幕共享告警阈值 * CPU：CPU告警阈值
	ThresholdType ShowQosThresholdRequestThresholdType `json:"thresholdType"`
}

func (o ShowQosThresholdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQosThresholdRequest struct{}"
	}

	return strings.Join([]string{"ShowQosThresholdRequest", string(data)}, " ")
}

type ShowQosThresholdRequestThresholdType struct {
	value string
}

type ShowQosThresholdRequestThresholdTypeEnum struct {
	AUDIO  ShowQosThresholdRequestThresholdType
	VIDEO  ShowQosThresholdRequestThresholdType
	SCREEN ShowQosThresholdRequestThresholdType
	CPU    ShowQosThresholdRequestThresholdType
}

func GetShowQosThresholdRequestThresholdTypeEnum() ShowQosThresholdRequestThresholdTypeEnum {
	return ShowQosThresholdRequestThresholdTypeEnum{
		AUDIO: ShowQosThresholdRequestThresholdType{
			value: "AUDIO",
		},
		VIDEO: ShowQosThresholdRequestThresholdType{
			value: "VIDEO",
		},
		SCREEN: ShowQosThresholdRequestThresholdType{
			value: "SCREEN",
		},
		CPU: ShowQosThresholdRequestThresholdType{
			value: "CPU",
		},
	}
}

func (c ShowQosThresholdRequestThresholdType) Value() string {
	return c.value
}

func (c ShowQosThresholdRequestThresholdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowQosThresholdRequestThresholdType) UnmarshalJSON(b []byte) error {
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
