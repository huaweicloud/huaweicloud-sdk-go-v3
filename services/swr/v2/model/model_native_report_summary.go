package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NativeReportSummary 云原生制品扫描报告摘要
type NativeReportSummary struct {

	// 报告ID
	ReportId *string `json:"report_id,omitempty"`

	// 制品扫描状态，Pending(扫描等待中)、Running(扫描中)、Error(扫描失败)、Success(扫描成功)
	ScanStatus *NativeReportSummaryScanStatus `json:"scan_status,omitempty"`

	// 制品扫描报告的总体严重程度，None(无评分), Low(低危), Medium(中危), High(高危), Critical(严重), Security(安全)
	Severity *NativeReportSummarySeverity `json:"severity,omitempty"`

	// 生成报告所花费的秒数时间
	Duration *int32 `json:"duration,omitempty"`

	Summary *VulnerabilitySummary `json:"summary,omitempty"`

	// 生成报告的扫描进程的开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 生成报告的扫描进程的结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 扫描的完成百分比，该值介于0和100之间
	CompletePercent *int32 `json:"complete_percent,omitempty"`

	Scanner *Scanner `json:"scanner,omitempty"`
}

func (o NativeReportSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NativeReportSummary struct{}"
	}

	return strings.Join([]string{"NativeReportSummary", string(data)}, " ")
}

type NativeReportSummaryScanStatus struct {
	value string
}

type NativeReportSummaryScanStatusEnum struct {
	SUCCESS NativeReportSummaryScanStatus
	ERROR   NativeReportSummaryScanStatus
	RUNNING NativeReportSummaryScanStatus
	PENDING NativeReportSummaryScanStatus
}

func GetNativeReportSummaryScanStatusEnum() NativeReportSummaryScanStatusEnum {
	return NativeReportSummaryScanStatusEnum{
		SUCCESS: NativeReportSummaryScanStatus{
			value: "Success",
		},
		ERROR: NativeReportSummaryScanStatus{
			value: "Error",
		},
		RUNNING: NativeReportSummaryScanStatus{
			value: "Running",
		},
		PENDING: NativeReportSummaryScanStatus{
			value: "Pending",
		},
	}
}

func (c NativeReportSummaryScanStatus) Value() string {
	return c.value
}

func (c NativeReportSummaryScanStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NativeReportSummaryScanStatus) UnmarshalJSON(b []byte) error {
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

type NativeReportSummarySeverity struct {
	value string
}

type NativeReportSummarySeverityEnum struct {
	NONE     NativeReportSummarySeverity
	LOW      NativeReportSummarySeverity
	MEDIUM   NativeReportSummarySeverity
	HIGH     NativeReportSummarySeverity
	CRITICAL NativeReportSummarySeverity
	SECURITY NativeReportSummarySeverity
}

func GetNativeReportSummarySeverityEnum() NativeReportSummarySeverityEnum {
	return NativeReportSummarySeverityEnum{
		NONE: NativeReportSummarySeverity{
			value: "None",
		},
		LOW: NativeReportSummarySeverity{
			value: "Low",
		},
		MEDIUM: NativeReportSummarySeverity{
			value: "Medium",
		},
		HIGH: NativeReportSummarySeverity{
			value: "High",
		},
		CRITICAL: NativeReportSummarySeverity{
			value: "Critical",
		},
		SECURITY: NativeReportSummarySeverity{
			value: "Security",
		},
	}
}

func (c NativeReportSummarySeverity) Value() string {
	return c.value
}

func (c NativeReportSummarySeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NativeReportSummarySeverity) UnmarshalJSON(b []byte) error {
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
