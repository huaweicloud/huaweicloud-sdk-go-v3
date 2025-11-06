package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RelationType **参数解释**： 屏蔽告警通知的实现方式。 **约束限制**： 不涉及。 **取值范围**： - ALARM_RULE：通过告警规则屏蔽告警通知(不支持通过该接口修改，请通过接口“[批量设置告警通知屏蔽规则](BatchUpdateNotificationMasks.xml)”设置)。 - RESOURCE：通过资源屏蔽告警通知。 - RESOURCE_POLICY_NOTIFICATION：通过告警策略屏蔽告警通知。 - EVENT.SYS: 通过事件来蔽告警通知。 - RESOURCE_POLICY_ALARM：（已废弃）通过屏蔽告警计算来屏蔽告警通知。 **默认取值**： 不涉及。
type RelationType struct {
	value string
}

type RelationTypeEnum struct {
	ALARM_RULE                   RelationType
	RESOURCE                     RelationType
	RESOURCE_POLICY_NOTIFICATION RelationType
	RESOURCE_POLICY_ALARM        RelationType
	EVENT_SYS                    RelationType
}

func GetRelationTypeEnum() RelationTypeEnum {
	return RelationTypeEnum{
		ALARM_RULE: RelationType{
			value: "ALARM_RULE",
		},
		RESOURCE: RelationType{
			value: "RESOURCE",
		},
		RESOURCE_POLICY_NOTIFICATION: RelationType{
			value: "RESOURCE_POLICY_NOTIFICATION",
		},
		RESOURCE_POLICY_ALARM: RelationType{
			value: "RESOURCE_POLICY_ALARM",
		},
		EVENT_SYS: RelationType{
			value: "EVENT.SYS",
		},
	}
}

func (c RelationType) Value() string {
	return c.value
}

func (c RelationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RelationType) UnmarshalJSON(b []byte) error {
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
