package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmType 告警规则类型，ALL_INSTANCE为全部资源指标告警，RESOURCE_GROUP为资源分组指标告警，MULTI_INSTANCE为指定资源指标告警，EVENT.SYS为系统事件告警，EVENT.CUSTOM自定义事件告警，DNSHealthCheck为健康检查告警；
type AlarmType struct {
	value string
}

type AlarmTypeEnum struct {
	EVENT_SYS        AlarmType
	EVENT_CUSTOM     AlarmType
	DNS_HEALTH_CHECK AlarmType
	RESOURCE_GROUP   AlarmType
	MULTI_INSTANCE   AlarmType
	ALL_INSTANCE     AlarmType
}

func GetAlarmTypeEnum() AlarmTypeEnum {
	return AlarmTypeEnum{
		EVENT_SYS: AlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: AlarmType{
			value: "EVENT.CUSTOM",
		},
		DNS_HEALTH_CHECK: AlarmType{
			value: "DNSHealthCheck",
		},
		RESOURCE_GROUP: AlarmType{
			value: "RESOURCE_GROUP",
		},
		MULTI_INSTANCE: AlarmType{
			value: "MULTI_INSTANCE",
		},
		ALL_INSTANCE: AlarmType{
			value: "ALL_INSTANCE",
		},
	}
}

func (c AlarmType) Value() string {
	return c.value
}

func (c AlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmType) UnmarshalJSON(b []byte) error {
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
