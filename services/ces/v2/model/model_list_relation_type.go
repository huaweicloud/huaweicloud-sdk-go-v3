package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRelationType **参数解释**： 屏蔽告警通知的实现方式。 **约束限制**： 不涉及。 **取值范围**： 枚举值，长度为[1,32]个字符，取值为: - ALARM_RULE：通过告警规则屏蔽告警通知。 - RESOURCE：通过资源屏蔽告警通知。使用方式：《告警屏蔽》页面点击《创建告警屏蔽》按钮，选择资源屏蔽。 - RESOURCE_POLICY_NOTIFICATION：通过告警策略屏蔽告警通知。使用方式：《告警屏蔽》页面点击《创建告警屏蔽》按钮，选择策略屏蔽。 - EVENT.SYS：屏蔽事件类告警通知。使用方式：《告警屏蔽》页面点击《创建告警屏蔽》按钮，选择事件屏蔽。 - RESOURCE_POLICY_ALARM：（已废弃，不推荐使用）通过屏蔽告警计算来屏蔽告警通知。 - DEFAULT：（已废弃，不推荐使用）默认包含RESOURCE、RESOURCE_POLICY_NOTIFICATION、EVENT.SYS **默认取值**： 不涉及。
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
