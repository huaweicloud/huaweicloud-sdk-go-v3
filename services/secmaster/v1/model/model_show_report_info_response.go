package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReportInfoResponse Response Object
type ShowReportInfoResponse struct {

	// 报告id
	Id *string `json:"id,omitempty"`

	// 报告名称
	ReportName *string `json:"report_name,omitempty"`

	// 报告周期 weekly, daily, annual, monthly
	ReportPeriod *ShowReportInfoResponseReportPeriod `json:"report_period,omitempty"`

	ReportRange *CreateReportRequestBodyReportRange `json:"report_range,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 通知任务id
	NotificationTask *string `json:"notification_task,omitempty"`

	// 布局id
	LayoutId *string `json:"layout_id,omitempty"`

	// 报告状态 enable, disable
	Status *ShowReportInfoResponseStatus `json:"status,omitempty"`

	// 报告是否已生成
	IsGenerated *bool `json:"is_generated,omitempty"`

	// 报告发送规则
	ReportRuleInfos *[]ReportRuleInfo `json:"report_rule_infos,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReportInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowReportInfoResponse", string(data)}, " ")
}

type ShowReportInfoResponseReportPeriod struct {
	value string
}

type ShowReportInfoResponseReportPeriodEnum struct {
	WEEKLY  ShowReportInfoResponseReportPeriod
	DAILY   ShowReportInfoResponseReportPeriod
	ANNUAL  ShowReportInfoResponseReportPeriod
	MONTHLY ShowReportInfoResponseReportPeriod
}

func GetShowReportInfoResponseReportPeriodEnum() ShowReportInfoResponseReportPeriodEnum {
	return ShowReportInfoResponseReportPeriodEnum{
		WEEKLY: ShowReportInfoResponseReportPeriod{
			value: "weekly",
		},
		DAILY: ShowReportInfoResponseReportPeriod{
			value: "daily",
		},
		ANNUAL: ShowReportInfoResponseReportPeriod{
			value: "annual",
		},
		MONTHLY: ShowReportInfoResponseReportPeriod{
			value: "monthly",
		},
	}
}

func (c ShowReportInfoResponseReportPeriod) Value() string {
	return c.value
}

func (c ShowReportInfoResponseReportPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReportInfoResponseReportPeriod) UnmarshalJSON(b []byte) error {
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

type ShowReportInfoResponseStatus struct {
	value string
}

type ShowReportInfoResponseStatusEnum struct {
	ENABLE  ShowReportInfoResponseStatus
	DISABLE ShowReportInfoResponseStatus
}

func GetShowReportInfoResponseStatusEnum() ShowReportInfoResponseStatusEnum {
	return ShowReportInfoResponseStatusEnum{
		ENABLE: ShowReportInfoResponseStatus{
			value: "enable",
		},
		DISABLE: ShowReportInfoResponseStatus{
			value: "disable",
		},
	}
}

func (c ShowReportInfoResponseStatus) Value() string {
	return c.value
}

func (c ShowReportInfoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReportInfoResponseStatus) UnmarshalJSON(b []byte) error {
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
