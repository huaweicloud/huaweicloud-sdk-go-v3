package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmTypeResp **参数解释**： 告警规则类型 **取值范围**： 枚举值。 - ALL_INSTANCE：全部资源指标告警。 - RESOURCE_GROUP：资源分组指标告警。 - MULTI_INSTANCE：指定资源指标告警。 - EVENT.SYS：系统事件告警。 - EVENT.CUSTOM：自定义事件告警。 - DNSHealthCheck：健康检查告警。
type AlarmTypeResp struct {
	value string
}

type AlarmTypeRespEnum struct {
	EVENT_SYS        AlarmTypeResp
	EVENT_CUSTOM     AlarmTypeResp
	DNS_HEALTH_CHECK AlarmTypeResp
	RESOURCE_GROUP   AlarmTypeResp
	MULTI_INSTANCE   AlarmTypeResp
	ALL_INSTANCE     AlarmTypeResp
}

func GetAlarmTypeRespEnum() AlarmTypeRespEnum {
	return AlarmTypeRespEnum{
		EVENT_SYS: AlarmTypeResp{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: AlarmTypeResp{
			value: "EVENT.CUSTOM",
		},
		DNS_HEALTH_CHECK: AlarmTypeResp{
			value: "DNSHealthCheck",
		},
		RESOURCE_GROUP: AlarmTypeResp{
			value: "RESOURCE_GROUP",
		},
		MULTI_INSTANCE: AlarmTypeResp{
			value: "MULTI_INSTANCE",
		},
		ALL_INSTANCE: AlarmTypeResp{
			value: "ALL_INSTANCE",
		},
	}
}

func (c AlarmTypeResp) Value() string {
	return c.value
}

func (c AlarmTypeResp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmTypeResp) UnmarshalJSON(b []byte) error {
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
