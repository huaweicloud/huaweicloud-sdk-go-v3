package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PauseScalingGroupOption 启停伸缩组请求
type PauseScalingGroupOption struct {

	// 启用或停止伸缩组操作的标识。启用：resume 停止：pause
	Action PauseScalingGroupOptionAction `json:"action"`
}

func (o PauseScalingGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingGroupOption struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupOption", string(data)}, " ")
}

type PauseScalingGroupOptionAction struct {
	value string
}

type PauseScalingGroupOptionActionEnum struct {
	PAUSE PauseScalingGroupOptionAction
}

func GetPauseScalingGroupOptionActionEnum() PauseScalingGroupOptionActionEnum {
	return PauseScalingGroupOptionActionEnum{
		PAUSE: PauseScalingGroupOptionAction{
			value: "pause",
		},
	}
}

func (c PauseScalingGroupOptionAction) Value() string {
	return c.value
}

func (c PauseScalingGroupOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PauseScalingGroupOptionAction) UnmarshalJSON(b []byte) error {
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
