package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlarmSummaryRequest Request Object
type ShowAlarmSummaryRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 告警命名空间，取值范围：SYS.CBR,SYS.RDS,SYS.GaussDB
	Namespace *ShowAlarmSummaryRequestNamespace `json:"namespace,omitempty"`

	// 告警状态，取值范围：ok,alarm,invalid
	Status *ShowAlarmSummaryRequestStatus `json:"status,omitempty"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 查询范围起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围结束时间，例如：2025-03-21T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowAlarmSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmSummaryRequest", string(data)}, " ")
}

type ShowAlarmSummaryRequestNamespace struct {
	value string
}

type ShowAlarmSummaryRequestNamespaceEnum struct {
	SYS_CBR      ShowAlarmSummaryRequestNamespace
	SYS_RDS      ShowAlarmSummaryRequestNamespace
	SYS_GAUSS_DB ShowAlarmSummaryRequestNamespace
}

func GetShowAlarmSummaryRequestNamespaceEnum() ShowAlarmSummaryRequestNamespaceEnum {
	return ShowAlarmSummaryRequestNamespaceEnum{
		SYS_CBR: ShowAlarmSummaryRequestNamespace{
			value: "SYS.CBR",
		},
		SYS_RDS: ShowAlarmSummaryRequestNamespace{
			value: "SYS.RDS",
		},
		SYS_GAUSS_DB: ShowAlarmSummaryRequestNamespace{
			value: "SYS.GaussDB",
		},
	}
}

func (c ShowAlarmSummaryRequestNamespace) Value() string {
	return c.value
}

func (c ShowAlarmSummaryRequestNamespace) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlarmSummaryRequestNamespace) UnmarshalJSON(b []byte) error {
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

type ShowAlarmSummaryRequestStatus struct {
	value string
}

type ShowAlarmSummaryRequestStatusEnum struct {
	OK      ShowAlarmSummaryRequestStatus
	ALARM   ShowAlarmSummaryRequestStatus
	INVALID ShowAlarmSummaryRequestStatus
}

func GetShowAlarmSummaryRequestStatusEnum() ShowAlarmSummaryRequestStatusEnum {
	return ShowAlarmSummaryRequestStatusEnum{
		OK: ShowAlarmSummaryRequestStatus{
			value: "ok",
		},
		ALARM: ShowAlarmSummaryRequestStatus{
			value: "alarm",
		},
		INVALID: ShowAlarmSummaryRequestStatus{
			value: "invalid",
		},
	}
}

func (c ShowAlarmSummaryRequestStatus) Value() string {
	return c.value
}

func (c ShowAlarmSummaryRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlarmSummaryRequestStatus) UnmarshalJSON(b []byte) error {
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
