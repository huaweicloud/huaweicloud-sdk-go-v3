package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteScalingPolicyOption 执行或启用或停止伸缩策略
type ExecuteScalingPolicyOption struct {

	// 执行或启用或停止伸缩策略操作的标识。执行：execute。启用：resume。停止：pause。
	Action ExecuteScalingPolicyOptionAction `json:"action"`
}

func (o ExecuteScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyOption", string(data)}, " ")
}

type ExecuteScalingPolicyOptionAction struct {
	value string
}

type ExecuteScalingPolicyOptionActionEnum struct {
	EXECUTE ExecuteScalingPolicyOptionAction
}

func GetExecuteScalingPolicyOptionActionEnum() ExecuteScalingPolicyOptionActionEnum {
	return ExecuteScalingPolicyOptionActionEnum{
		EXECUTE: ExecuteScalingPolicyOptionAction{
			value: "execute",
		},
	}
}

func (c ExecuteScalingPolicyOptionAction) Value() string {
	return c.value
}

func (c ExecuteScalingPolicyOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteScalingPolicyOptionAction) UnmarshalJSON(b []byte) error {
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
