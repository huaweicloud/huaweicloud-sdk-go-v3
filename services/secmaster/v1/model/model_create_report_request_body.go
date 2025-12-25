package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateReportRequestBody 创建报告请求体
type CreateReportRequestBody struct {

	// 报告名称
	ReportName string `json:"report_name"`

	// 报告周期 weekly, daily, annual, monthly
	ReportPeriod CreateReportRequestBodyReportPeriod `json:"report_period"`

	ReportRange *CreateReportRequestBodyReportRange `json:"report_range"`

	// 语言
	Language string `json:"language"`

	// 通知任务id
	NotificationTask *string `json:"notification_task,omitempty"`

	// 布局id
	LayoutId string `json:"layout_id"`

	// 邮件标题
	Title *string `json:"title,omitempty"`

	// 收件人邮箱
	To *string `json:"to,omitempty"`

	// 抄送人邮箱
	Cc *string `json:"cc,omitempty"`

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

	// 报告发送规则
	ReportRuleInfos *[]ReportRuleRequest `json:"report_rule_infos,omitempty"`
}

func (o CreateReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReportRequestBody", string(data)}, " ")
}

type CreateReportRequestBodyReportPeriod struct {
	value string
}

type CreateReportRequestBodyReportPeriodEnum struct {
	WEEKLY  CreateReportRequestBodyReportPeriod
	DAILY   CreateReportRequestBodyReportPeriod
	ANNUAL  CreateReportRequestBodyReportPeriod
	MONTHLY CreateReportRequestBodyReportPeriod
}

func GetCreateReportRequestBodyReportPeriodEnum() CreateReportRequestBodyReportPeriodEnum {
	return CreateReportRequestBodyReportPeriodEnum{
		WEEKLY: CreateReportRequestBodyReportPeriod{
			value: "weekly",
		},
		DAILY: CreateReportRequestBodyReportPeriod{
			value: "daily",
		},
		ANNUAL: CreateReportRequestBodyReportPeriod{
			value: "annual",
		},
		MONTHLY: CreateReportRequestBodyReportPeriod{
			value: "monthly",
		},
	}
}

func (c CreateReportRequestBodyReportPeriod) Value() string {
	return c.value
}

func (c CreateReportRequestBodyReportPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateReportRequestBodyReportPeriod) UnmarshalJSON(b []byte) error {
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
