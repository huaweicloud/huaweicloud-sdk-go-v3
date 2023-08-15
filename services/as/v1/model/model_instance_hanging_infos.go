package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceHangingInfos 挂钩实例信息
type InstanceHangingInfos struct {

	// 生命周期挂钩名称。
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	// 生命周期操作令牌，用于指定生命周期回调对象。
	LifecycleActionKey *string `json:"lifecycle_action_key,omitempty"`

	// 伸缩实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 伸缩组ID。
	ScalingGroupId *string `json:"scaling_group_id,omitempty"`

	// 伸缩实例挂钩的挂起状态。HANGING：挂起。CONTINUE：继续。ABANDON：终止。
	LifecycleHookStatus *InstanceHangingInfosLifecycleHookStatus `json:"lifecycle_hook_status,omitempty"`

	// 超时时间，遵循UTC时间，格式为：YYYY-MM-DDThh:mm:ssZZ。
	Timeout *string `json:"timeout,omitempty"`

	// 生命周期挂钩默认回调操作。
	DefaultResult *string `json:"default_result,omitempty"`
}

func (o InstanceHangingInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceHangingInfos struct{}"
	}

	return strings.Join([]string{"InstanceHangingInfos", string(data)}, " ")
}

type InstanceHangingInfosLifecycleHookStatus struct {
	value string
}

type InstanceHangingInfosLifecycleHookStatusEnum struct {
	HANGING  InstanceHangingInfosLifecycleHookStatus
	CONTINUE InstanceHangingInfosLifecycleHookStatus
	ABANDON  InstanceHangingInfosLifecycleHookStatus
}

func GetInstanceHangingInfosLifecycleHookStatusEnum() InstanceHangingInfosLifecycleHookStatusEnum {
	return InstanceHangingInfosLifecycleHookStatusEnum{
		HANGING: InstanceHangingInfosLifecycleHookStatus{
			value: "HANGING",
		},
		CONTINUE: InstanceHangingInfosLifecycleHookStatus{
			value: "CONTINUE",
		},
		ABANDON: InstanceHangingInfosLifecycleHookStatus{
			value: "ABANDON",
		},
	}
}

func (c InstanceHangingInfosLifecycleHookStatus) Value() string {
	return c.value
}

func (c InstanceHangingInfosLifecycleHookStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceHangingInfosLifecycleHookStatus) UnmarshalJSON(b []byte) error {
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
