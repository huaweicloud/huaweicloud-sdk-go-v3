package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PauseScalingPolicyOption 执行或启用或停止伸缩策略
type PauseScalingPolicyOption struct {

	// 执行或启用或停止伸缩策略操作的标识。执行：execute。启用：resume。停止：pause。
	Action PauseScalingPolicyOptionAction `json:"action"`
}

func (o PauseScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"PauseScalingPolicyOption", string(data)}, " ")
}

type PauseScalingPolicyOptionAction struct {
	value string
}

type PauseScalingPolicyOptionActionEnum struct {
	PAUSE PauseScalingPolicyOptionAction
}

func GetPauseScalingPolicyOptionActionEnum() PauseScalingPolicyOptionActionEnum {
	return PauseScalingPolicyOptionActionEnum{
		PAUSE: PauseScalingPolicyOptionAction{
			value: "pause",
		},
	}
}

func (c PauseScalingPolicyOptionAction) Value() string {
	return c.value
}

func (c PauseScalingPolicyOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PauseScalingPolicyOptionAction) UnmarshalJSON(b []byte) error {
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
