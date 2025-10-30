package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventTypeDetailResponseInfo 大类别对应的名称与数量列表
type EventTypeDetailResponseInfo struct {

	// 大类别对应的名字:   - malicious_software：恶意软件   - \"exploit_attack\"：漏洞利用   - \"system_abnormal_behavior\"：系统异常行为   - \"user_abnormal_behavior\"：用户异常行为   - \"network_abnormal_access\"：网络异常访问   - \"resource_recon\"：资源侦查   - \"risk_audit\"：风险审计   - \"information_destroy\"：信息破坏   - \"denial_of_service\"：拒绝服务攻击   - \"advanced_threat\"：高级威胁   - \"extend\"：其余未列出的，里面的类别显示到外层中
	EventTypeName *EventTypeDetailResponseInfoEventTypeName `json:"event_type_name,omitempty"`

	// 事件类别的总数量
	EventTypeNum *int32 `json:"event_type_num,omitempty"`

	// 小类别对应的名称与数量列表
	EventTypeList *[]EventTypeResponseInfo `json:"event_type_list,omitempty"`
}

func (o EventTypeDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventTypeDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"EventTypeDetailResponseInfo", string(data)}, " ")
}

type EventTypeDetailResponseInfoEventTypeName struct {
	value string
}

type EventTypeDetailResponseInfoEventTypeNameEnum struct {
	MALICIOUS_SOFTWARE       EventTypeDetailResponseInfoEventTypeName
	EXPLOIT_ATTACK           EventTypeDetailResponseInfoEventTypeName
	SYSTEM_ABNORMAL_BEHAVIOR EventTypeDetailResponseInfoEventTypeName
	USER_ABNORMAL_BEHAVIOR   EventTypeDetailResponseInfoEventTypeName
	NETWORK_ABNORMAL_ACCESS  EventTypeDetailResponseInfoEventTypeName
	RESOURCE_RECON           EventTypeDetailResponseInfoEventTypeName
	RISK_AUDIT               EventTypeDetailResponseInfoEventTypeName
	INFORMATION_DESTROY      EventTypeDetailResponseInfoEventTypeName
	DENIAL_OF_SERVICE        EventTypeDetailResponseInfoEventTypeName
	ADVANCED_THREAT          EventTypeDetailResponseInfoEventTypeName
	EXTEND                   EventTypeDetailResponseInfoEventTypeName
}

func GetEventTypeDetailResponseInfoEventTypeNameEnum() EventTypeDetailResponseInfoEventTypeNameEnum {
	return EventTypeDetailResponseInfoEventTypeNameEnum{
		MALICIOUS_SOFTWARE: EventTypeDetailResponseInfoEventTypeName{
			value: "malicious_software",
		},
		EXPLOIT_ATTACK: EventTypeDetailResponseInfoEventTypeName{
			value: "exploit_attack",
		},
		SYSTEM_ABNORMAL_BEHAVIOR: EventTypeDetailResponseInfoEventTypeName{
			value: "system_abnormal_behavior",
		},
		USER_ABNORMAL_BEHAVIOR: EventTypeDetailResponseInfoEventTypeName{
			value: "user_abnormal_behavior",
		},
		NETWORK_ABNORMAL_ACCESS: EventTypeDetailResponseInfoEventTypeName{
			value: "network_abnormal_access",
		},
		RESOURCE_RECON: EventTypeDetailResponseInfoEventTypeName{
			value: "resource_recon",
		},
		RISK_AUDIT: EventTypeDetailResponseInfoEventTypeName{
			value: "risk_audit",
		},
		INFORMATION_DESTROY: EventTypeDetailResponseInfoEventTypeName{
			value: "information_destroy",
		},
		DENIAL_OF_SERVICE: EventTypeDetailResponseInfoEventTypeName{
			value: "denial_of_service",
		},
		ADVANCED_THREAT: EventTypeDetailResponseInfoEventTypeName{
			value: "advanced_threat",
		},
		EXTEND: EventTypeDetailResponseInfoEventTypeName{
			value: "extend",
		},
	}
}

func (c EventTypeDetailResponseInfoEventTypeName) Value() string {
	return c.value
}

func (c EventTypeDetailResponseInfoEventTypeName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventTypeDetailResponseInfoEventTypeName) UnmarshalJSON(b []byte) error {
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
