package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRelationType 屏蔽告警通知或者屏蔽告警计算的实现方式。ALARM_RULE：通过告警规则屏蔽告警通知，RESOURCE：通过资源屏蔽告警通知，RESOURCE_POLICY_NOTIFICATION：通过告警策略和告警资源屏蔽告警通知，RESOURCE_POLICY_ALARM：通过告警策略和告警资源屏蔽计算告警，DEFAULT：默认包含RESOURCE、RESOURCE_POLICY_NOTIFICATION（查询告警屏蔽列表时使用）。
type ListRelationType struct {
	value string
}

type ListRelationTypeEnum struct {
	ALARM_RULE                   ListRelationType
	RESOURCE                     ListRelationType
	RESOURCE_POLICY_NOTIFICATION ListRelationType
	RESOURCE_POLICY_ALARM        ListRelationType
	DEFAULT                      ListRelationType
}

func GetListRelationTypeEnum() ListRelationTypeEnum {
	return ListRelationTypeEnum{
		ALARM_RULE: ListRelationType{
			value: "ALARM_RULE",
		},
		RESOURCE: ListRelationType{
			value: "RESOURCE",
		},
		RESOURCE_POLICY_NOTIFICATION: ListRelationType{
			value: "RESOURCE_POLICY_NOTIFICATION",
		},
		RESOURCE_POLICY_ALARM: ListRelationType{
			value: "RESOURCE_POLICY_ALARM",
		},
		DEFAULT: ListRelationType{
			value: "DEFAULT",
		},
	}
}

func (c ListRelationType) Value() string {
	return c.value
}

func (c ListRelationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRelationType) UnmarshalJSON(b []byte) error {
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
