package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetResourceGroupResources struct {

	// 资源健康状态，取值为health（已设置告警规则且无告警触发的资源）、unhealthy（已设置告警规则且有告警触发的资源）、no_alarm_rule（未关联告警规则）
	Status GetResourceGroupResourcesStatus `json:"status"`

	// 资源的维度信息
	Dimensions []Dimension2 `json:"dimensions"`
}

func (o GetResourceGroupResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceGroupResources struct{}"
	}

	return strings.Join([]string{"GetResourceGroupResources", string(data)}, " ")
}

type GetResourceGroupResourcesStatus struct {
	value string
}

type GetResourceGroupResourcesStatusEnum struct {
	HEALTH        GetResourceGroupResourcesStatus
	UNHEALTHY     GetResourceGroupResourcesStatus
	NO_ALARM_RULE GetResourceGroupResourcesStatus
}

func GetGetResourceGroupResourcesStatusEnum() GetResourceGroupResourcesStatusEnum {
	return GetResourceGroupResourcesStatusEnum{
		HEALTH: GetResourceGroupResourcesStatus{
			value: "health",
		},
		UNHEALTHY: GetResourceGroupResourcesStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: GetResourceGroupResourcesStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c GetResourceGroupResourcesStatus) Value() string {
	return c.value
}

func (c GetResourceGroupResourcesStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetResourceGroupResourcesStatus) UnmarshalJSON(b []byte) error {
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
