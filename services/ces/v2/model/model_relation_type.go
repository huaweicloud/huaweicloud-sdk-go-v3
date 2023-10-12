package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RelationType 关联类型。ALARM_RULE：关联告警规则，RESOURCE：关联资源，RESOURCE_POLICY_NOTIFICATION：关联资源策略屏蔽告警通知，RESOURCE_POLICY_ALARM：关联资源策略不计算告警。
type RelationType struct {
	value string
}

type RelationTypeEnum struct {
	ALARM_RULE                   RelationType
	RESOURCE                     RelationType
	RESOURCE_POLICY_NOTIFICATION RelationType
	RESOURCE_POLICY_ALARM        RelationType
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
