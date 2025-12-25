package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateReportRequestBody 报告更新请求体
type UpdateReportRequestBody struct {

	// 报告状态 enable, disable
	Status *UpdateReportRequestBodyStatus `json:"status,omitempty"`

	// 报告名称
	ReportName string `json:"report_name"`

	// 报告周期 weekly, daily, annual, monthly
	ReportPeriod *UpdateReportRequestBodyReportPeriod `json:"report_period,omitempty"`

	ReportRange *UpdateReportRequestBodyReportRange `json:"report_range,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 通知任务id
	NotificationTask *string `json:"notification_task,omitempty"`

	// 布局id
	LayoutId *string `json:"layout_id,omitempty"`

	// 报告Id
	ReportId *string `json:"report_id,omitempty"`

	// 报告发送规则
	ReportRuleInfos *[]ReportRuleRequest `json:"report_rule_infos,omitempty"`

	// 邮件标题
	Title *string `json:"title,omitempty"`

	// 收件人
	To *string `json:"to,omitempty"`

	// 抄送人
	Cc *string `json:"cc,omitempty"`

	// 备注
	Subject *string `json:"subject,omitempty"`

	// 邮件内容
	Content *string `json:"content,omitempty"`

	// 报告类型，支持图片、pdf、html
	ReportFileType *string `json:"report_file_type,omitempty"`

	// 报告发送频率，daily，monthly，weekly
	Frequency *string `json:"frequency,omitempty"`

	// 报告页面内容
	BindingWizard string `json:"binding_wizard"`

	// 信息
	Timezone *string `json:"timezone,omitempty"`
}

func (o UpdateReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateReportRequestBody", string(data)}, " ")
}

type UpdateReportRequestBodyStatus struct {
	value string
}

type UpdateReportRequestBodyStatusEnum struct {
	ENABLE  UpdateReportRequestBodyStatus
	DISABLE UpdateReportRequestBodyStatus
}

func GetUpdateReportRequestBodyStatusEnum() UpdateReportRequestBodyStatusEnum {
	return UpdateReportRequestBodyStatusEnum{
		ENABLE: UpdateReportRequestBodyStatus{
			value: "enable",
		},
		DISABLE: UpdateReportRequestBodyStatus{
			value: "disable",
		},
	}
}

func (c UpdateReportRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateReportRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateReportRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type UpdateReportRequestBodyReportPeriod struct {
	value string
}

type UpdateReportRequestBodyReportPeriodEnum struct {
	WEEKLY  UpdateReportRequestBodyReportPeriod
	DAILY   UpdateReportRequestBodyReportPeriod
	ANNUAL  UpdateReportRequestBodyReportPeriod
	MONTHLY UpdateReportRequestBodyReportPeriod
}

func GetUpdateReportRequestBodyReportPeriodEnum() UpdateReportRequestBodyReportPeriodEnum {
	return UpdateReportRequestBodyReportPeriodEnum{
		WEEKLY: UpdateReportRequestBodyReportPeriod{
			value: "weekly",
		},
		DAILY: UpdateReportRequestBodyReportPeriod{
			value: "daily",
		},
		ANNUAL: UpdateReportRequestBodyReportPeriod{
			value: "annual",
		},
		MONTHLY: UpdateReportRequestBodyReportPeriod{
			value: "monthly",
		},
	}
}

func (c UpdateReportRequestBodyReportPeriod) Value() string {
	return c.value
}

func (c UpdateReportRequestBodyReportPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateReportRequestBodyReportPeriod) UnmarshalJSON(b []byte) error {
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
