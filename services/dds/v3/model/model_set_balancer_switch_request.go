package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetBalancerSwitchRequest Request Object
type SetBalancerSwitchRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 开启或关闭集群均衡。
	Action SetBalancerSwitchRequestAction `json:"action"`
}

func (o SetBalancerSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBalancerSwitchRequest struct{}"
	}

	return strings.Join([]string{"SetBalancerSwitchRequest", string(data)}, " ")
}

type SetBalancerSwitchRequestAction struct {
	value string
}

type SetBalancerSwitchRequestActionEnum struct {
	START SetBalancerSwitchRequestAction
	STOP  SetBalancerSwitchRequestAction
}

func GetSetBalancerSwitchRequestActionEnum() SetBalancerSwitchRequestActionEnum {
	return SetBalancerSwitchRequestActionEnum{
		START: SetBalancerSwitchRequestAction{
			value: "start",
		},
		STOP: SetBalancerSwitchRequestAction{
			value: "stop",
		},
	}
}

func (c SetBalancerSwitchRequestAction) Value() string {
	return c.value
}

func (c SetBalancerSwitchRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetBalancerSwitchRequestAction) UnmarshalJSON(b []byte) error {
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
