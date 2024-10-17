package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlarmLogRequest struct {
	Time *AlarmLogRequestTime `json:"time"`

	// 风险级别 - LOW - MEDIUM - HIGH
	Risk *AlarmLogRequestRisk `json:"risk,omitempty"`

	// 告警类型 - RISK_RULE: 风险规则 - RISK_CPU: CPU超限 - RISK_MEMORY: 内存超限 - RISK_DISK: 磁盘超限 - RISK_DISK_CAPACITY: 磁盘容量不足六个月 - RISK_BACKUP: 备份失败 - AUDIT_QPS_OVERFLOW: 流量超限入库延迟告警 - RISK_AGENT: Agent异常 - AUDIT_BACKUP_FAILED: 实例备份失败(运维侧)
	Type *string `json:"type,omitempty"`

	// 告警确认状态 - DONE: 已确认 - UNDO: 未确认
	Status *AlarmLogRequestStatus `json:"status,omitempty"`

	// 页码
	Page *int32 `json:"page,omitempty"`

	// 每页条数
	Size *int32 `json:"size,omitempty"`
}

func (o AlarmLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLogRequest struct{}"
	}

	return strings.Join([]string{"AlarmLogRequest", string(data)}, " ")
}

type AlarmLogRequestRisk struct {
	value string
}

type AlarmLogRequestRiskEnum struct {
	LOW    AlarmLogRequestRisk
	MEDIUM AlarmLogRequestRisk
	HIGH   AlarmLogRequestRisk
}

func GetAlarmLogRequestRiskEnum() AlarmLogRequestRiskEnum {
	return AlarmLogRequestRiskEnum{
		LOW: AlarmLogRequestRisk{
			value: "LOW",
		},
		MEDIUM: AlarmLogRequestRisk{
			value: "MEDIUM",
		},
		HIGH: AlarmLogRequestRisk{
			value: "HIGH",
		},
	}
}

func (c AlarmLogRequestRisk) Value() string {
	return c.value
}

func (c AlarmLogRequestRisk) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmLogRequestRisk) UnmarshalJSON(b []byte) error {
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

type AlarmLogRequestStatus struct {
	value string
}

type AlarmLogRequestStatusEnum struct {
	DONE AlarmLogRequestStatus
	UNDO AlarmLogRequestStatus
}

func GetAlarmLogRequestStatusEnum() AlarmLogRequestStatusEnum {
	return AlarmLogRequestStatusEnum{
		DONE: AlarmLogRequestStatus{
			value: "DONE",
		},
		UNDO: AlarmLogRequestStatus{
			value: "UNDO",
		},
	}
}

func (c AlarmLogRequestStatus) Value() string {
	return c.value
}

func (c AlarmLogRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmLogRequestStatus) UnmarshalJSON(b []byte) error {
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
