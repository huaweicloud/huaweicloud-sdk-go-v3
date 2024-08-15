package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReportCustomEventRequestBody struct {

	// 告警id
	AlarmId string `json:"alarmId"`

	// 告警名称
	AlarmName string `json:"alarmName"`

	// 告警级别。取值为Critical（紧急）, Major（重要）, Minor（次要）, Info（提示）
	AlarmLevel ReportCustomEventRequestBodyAlarmLevel `json:"alarmLevel"`

	// 告警发生时间
	Time int64 `json:"time"`

	// 告警发生时间
	NameSpace string `json:"nameSpace"`

	// 告警发生区域
	RegionId *string `json:"regionId,omitempty"`

	// 应用id
	ApplicationId *string `json:"applicationId,omitempty"`

	// 资源名称
	ResourceName *string `json:"resourceName,omitempty"`

	// 资源ID
	ResourceId *string `json:"resourceId,omitempty"`

	// 告警描述
	AlarmDesc string `json:"alarmDesc"`

	// 原始告警URL
	Url *string `json:"URL,omitempty"`

	// 告警状态。一般取值为alarm（告警中）和ok（已恢复）
	AlarmStatus *ReportCustomEventRequestBodyAlarmStatus `json:"alarmStatus,omitempty"`

	// 告警源
	AlarmSource string `json:"alarmSource"`

	// 告警附加信息
	Additional *interface{} `json:"additional,omitempty"`
}

func (o ReportCustomEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportCustomEventRequestBody struct{}"
	}

	return strings.Join([]string{"ReportCustomEventRequestBody", string(data)}, " ")
}

type ReportCustomEventRequestBodyAlarmLevel struct {
	value string
}

type ReportCustomEventRequestBodyAlarmLevelEnum struct {
	CRITICAL ReportCustomEventRequestBodyAlarmLevel
	MAJOR    ReportCustomEventRequestBodyAlarmLevel
	MINOR    ReportCustomEventRequestBodyAlarmLevel
	INFO     ReportCustomEventRequestBodyAlarmLevel
}

func GetReportCustomEventRequestBodyAlarmLevelEnum() ReportCustomEventRequestBodyAlarmLevelEnum {
	return ReportCustomEventRequestBodyAlarmLevelEnum{
		CRITICAL: ReportCustomEventRequestBodyAlarmLevel{
			value: "Critical",
		},
		MAJOR: ReportCustomEventRequestBodyAlarmLevel{
			value: "Major",
		},
		MINOR: ReportCustomEventRequestBodyAlarmLevel{
			value: "Minor",
		},
		INFO: ReportCustomEventRequestBodyAlarmLevel{
			value: "Info",
		},
	}
}

func (c ReportCustomEventRequestBodyAlarmLevel) Value() string {
	return c.value
}

func (c ReportCustomEventRequestBodyAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReportCustomEventRequestBodyAlarmLevel) UnmarshalJSON(b []byte) error {
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

type ReportCustomEventRequestBodyAlarmStatus struct {
	value string
}

type ReportCustomEventRequestBodyAlarmStatusEnum struct {
	ALARM ReportCustomEventRequestBodyAlarmStatus
	OK    ReportCustomEventRequestBodyAlarmStatus
}

func GetReportCustomEventRequestBodyAlarmStatusEnum() ReportCustomEventRequestBodyAlarmStatusEnum {
	return ReportCustomEventRequestBodyAlarmStatusEnum{
		ALARM: ReportCustomEventRequestBodyAlarmStatus{
			value: "alarm",
		},
		OK: ReportCustomEventRequestBodyAlarmStatus{
			value: "ok",
		},
	}
}

func (c ReportCustomEventRequestBodyAlarmStatus) Value() string {
	return c.value
}

func (c ReportCustomEventRequestBodyAlarmStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReportCustomEventRequestBodyAlarmStatus) UnmarshalJSON(b []byte) error {
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
