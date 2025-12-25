package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReportInfo 报告详情
type ReportInfo struct {

	// 报告id
	Id string `json:"id"`

	// 报告名称
	ReportName string `json:"report_name"`

	// 报告周期 weekly, daily, annual, monthly
	ReportPeriod ReportInfoReportPeriod `json:"report_period"`

	ReportRange *CreateReportRequestBodyReportRange `json:"report_range"`

	// 语言
	Language string `json:"language"`

	// 通知任务id
	NotificationTask string `json:"notification_task"`

	// 布局id
	LayoutId string `json:"layout_id"`

	// 报告状态 enable, disable
	Status ReportInfoStatus `json:"status"`

	// 报告是否已生成
	IsGenerated bool `json:"is_generated"`

	// 报告发送规则
	ReportRuleInfos *[]ReportRuleInfo `json:"report_rule_infos,omitempty"`
}

func (o ReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportInfo struct{}"
	}

	return strings.Join([]string{"ReportInfo", string(data)}, " ")
}

type ReportInfoReportPeriod struct {
	value string
}

type ReportInfoReportPeriodEnum struct {
	WEEKLY  ReportInfoReportPeriod
	DAILY   ReportInfoReportPeriod
	ANNUAL  ReportInfoReportPeriod
	MONTHLY ReportInfoReportPeriod
}

func GetReportInfoReportPeriodEnum() ReportInfoReportPeriodEnum {
	return ReportInfoReportPeriodEnum{
		WEEKLY: ReportInfoReportPeriod{
			value: "weekly",
		},
		DAILY: ReportInfoReportPeriod{
			value: "daily",
		},
		ANNUAL: ReportInfoReportPeriod{
			value: "annual",
		},
		MONTHLY: ReportInfoReportPeriod{
			value: "monthly",
		},
	}
}

func (c ReportInfoReportPeriod) Value() string {
	return c.value
}

func (c ReportInfoReportPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReportInfoReportPeriod) UnmarshalJSON(b []byte) error {
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

type ReportInfoStatus struct {
	value string
}

type ReportInfoStatusEnum struct {
	ENABLE  ReportInfoStatus
	DISABLE ReportInfoStatus
}

func GetReportInfoStatusEnum() ReportInfoStatusEnum {
	return ReportInfoStatusEnum{
		ENABLE: ReportInfoStatus{
			value: "enable",
		},
		DISABLE: ReportInfoStatus{
			value: "disable",
		},
	}
}

func (c ReportInfoStatus) Value() string {
	return c.value
}

func (c ReportInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReportInfoStatus) UnmarshalJSON(b []byte) error {
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
