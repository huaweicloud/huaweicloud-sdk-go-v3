package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Alert 告警实体信息
type Alert struct {

	// 告警对象的版本，该字段的值必须为云SSA服务确定的官方发布版本之一
	Version *string `json:"version,omitempty"`

	// 事件唯一标识，UUID格式，最大36个字符
	Id *string `json:"id,omitempty"`

	// 数据投递后，被委托用户的domain_id
	DomainId *string `json:"domain_id,omitempty"`

	// 数据投递后，被委托用户的region_id
	RegionId *string `json:"region_id,omitempty"`

	// 当前的工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 标签，仅展示
	Labels *string `json:"labels,omitempty"`

	Environment *AlertEnvironment `json:"environment,omitempty"`

	DataSource *AlertDataSource `json:"data_source,omitempty"`

	// 首次发现时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	FirstObservedTime *string `json:"first_observed_time,omitempty"`

	// 最近发现时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	LastObservedTime *string `json:"last_observed_time,omitempty"`

	// 记录时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	CreateTime *string `json:"create_time,omitempty"`

	// 接收时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	ArriveTime *string `json:"arrive_time,omitempty"`

	// 告警标题
	Title *string `json:"title,omitempty"`

	// 告警描述信息
	Description *string `json:"description,omitempty"`

	// 告警URL链接，指向数据源产品中有关当前事件说明的页面
	SourceUrl *string `json:"source_url,omitempty"`

	// 事件发生次数
	Count *int32 `json:"count,omitempty"`

	// 事件的置信度。置信度的定义旨在说明识别的行为或问题的可能性。 取值范围：0-100，0表示置信度为0%，100表示置信度为100%
	Confidence *int32 `json:"confidence,omitempty"`

	// 严重性等级，取值范围：Tips | Low | Medium | High | Fatal 说明： 0: Tips – 未发现任何问题。 1: Low – 无需针对问题执行任何操作。 2: Medium – 问题需要处理，但不紧急。 3: High – 问题必须优先处理。 4: Fatal – 问题必须立即处理，以防止产生进一步的损害
	Severity *AlertSeverity `json:"severity,omitempty"`

	// 关键性，是指事件涉及的资产的重要性级别。 取值范围：0-100，0表示资产不关键，100表示最关键资产
	Criticality *int32 `json:"criticality,omitempty"`

	AlertType *AlertAlertType `json:"alert_type,omitempty"`

	// 网络信息
	NetworkList *[]AlertNetworkList `json:"network_list,omitempty"`

	// 受影响资产
	ResourceList *[]AlertResourceList `json:"resource_list,omitempty"`

	Remediation *AlertRemediation `json:"remediation,omitempty"`

	// 验证状态，标识事件的准确性。可选类型如下： Unknown – 未知 True_Positive – 确认 False_Positive – 误报 默认填写Unknown
	VerificationState *AlertVerificationState `json:"verification_state,omitempty"`

	// 事件处理状态，可选类型如下： Open – 打开，默认 Block – 阻塞 Closed – 关闭 默认填写Open
	HandleStatus *AlertHandleStatus `json:"handle_status,omitempty"`

	// 约束闭环时间：设置风险接受持续时间。单位：小时
	Sla *int32 `json:"sla,omitempty"`

	// 更新时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	UpdateTime *string `json:"update_time,omitempty"`

	// 关闭时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	CloseTime *string `json:"close_time,omitempty"`

	// 周期/处置阶段编号 Preparation|Detection and Analysis|Contain，Eradication& Recovery|Post-Incident-Activity
	IpdrrPhase *AlertIpdrrPhase `json:"ipdrr_phase,omitempty"`

	// 调试字段
	Simulation *string `json:"simulation,omitempty"`

	// 告警调查员
	Actor *string `json:"actor,omitempty"`

	// 责任人、服务责任人
	Owner *string `json:"owner,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 关闭原因: 误检 - False detection 已解决 - Resolved 重复 - Repeated 其他 - Other
	CloseReason *AlertCloseReason `json:"close_reason,omitempty"`

	// 关闭评论
	CloseComment *string `json:"close_comment,omitempty"`

	Malware *AlertMalware `json:"malware,omitempty"`

	// 系统信息
	SystemInfo *interface{} `json:"system_info,omitempty"`

	// 进程信息
	Process *[]AlertProcess `json:"process,omitempty"`

	// 用户信息
	UserInfo *[]AlertUserInfo `json:"user_info,omitempty"`

	// 文件信息
	FileInfo *[]AlertFileInfo `json:"file_info,omitempty"`

	// 告警管理列表的布局字段
	SystemAlertTable *interface{} `json:"system_alert_table,omitempty"`
}

func (o Alert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Alert struct{}"
	}

	return strings.Join([]string{"Alert", string(data)}, " ")
}

type AlertSeverity struct {
	value string
}

type AlertSeverityEnum struct {
	TIPS   AlertSeverity
	LOW    AlertSeverity
	MEDIUM AlertSeverity
	HIGH   AlertSeverity
	FATAL  AlertSeverity
}

func GetAlertSeverityEnum() AlertSeverityEnum {
	return AlertSeverityEnum{
		TIPS: AlertSeverity{
			value: "Tips",
		},
		LOW: AlertSeverity{
			value: "Low",
		},
		MEDIUM: AlertSeverity{
			value: "Medium",
		},
		HIGH: AlertSeverity{
			value: "High",
		},
		FATAL: AlertSeverity{
			value: "Fatal",
		},
	}
}

func (c AlertSeverity) Value() string {
	return c.value
}

func (c AlertSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertSeverity) UnmarshalJSON(b []byte) error {
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

type AlertVerificationState struct {
	value string
}

type AlertVerificationStateEnum struct {
	UNKNOWN        AlertVerificationState
	TRUE_POSITIVE  AlertVerificationState
	FALSE_POSITIVE AlertVerificationState
}

func GetAlertVerificationStateEnum() AlertVerificationStateEnum {
	return AlertVerificationStateEnum{
		UNKNOWN: AlertVerificationState{
			value: "Unknown",
		},
		TRUE_POSITIVE: AlertVerificationState{
			value: "True_Positive",
		},
		FALSE_POSITIVE: AlertVerificationState{
			value: "False_Positive",
		},
	}
}

func (c AlertVerificationState) Value() string {
	return c.value
}

func (c AlertVerificationState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertVerificationState) UnmarshalJSON(b []byte) error {
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

type AlertHandleStatus struct {
	value string
}

type AlertHandleStatusEnum struct {
	OPEN   AlertHandleStatus
	BLOCK  AlertHandleStatus
	CLOSED AlertHandleStatus
}

func GetAlertHandleStatusEnum() AlertHandleStatusEnum {
	return AlertHandleStatusEnum{
		OPEN: AlertHandleStatus{
			value: "Open",
		},
		BLOCK: AlertHandleStatus{
			value: "Block",
		},
		CLOSED: AlertHandleStatus{
			value: "Closed",
		},
	}
}

func (c AlertHandleStatus) Value() string {
	return c.value
}

func (c AlertHandleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertHandleStatus) UnmarshalJSON(b []byte) error {
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

type AlertIpdrrPhase struct {
	value string
}

type AlertIpdrrPhaseEnum struct {
	PREPARATION                 AlertIpdrrPhase
	DETECTION_AND_ANALYSIS      AlertIpdrrPhase
	CONTAINERADICATION_RECOVERY AlertIpdrrPhase
	POST_INCIDENT_ACTIVITY      AlertIpdrrPhase
}

func GetAlertIpdrrPhaseEnum() AlertIpdrrPhaseEnum {
	return AlertIpdrrPhaseEnum{
		PREPARATION: AlertIpdrrPhase{
			value: "Preparation",
		},
		DETECTION_AND_ANALYSIS: AlertIpdrrPhase{
			value: "Detection and Analysis",
		},
		CONTAINERADICATION_RECOVERY: AlertIpdrrPhase{
			value: "Contain，Eradication& Recovery",
		},
		POST_INCIDENT_ACTIVITY: AlertIpdrrPhase{
			value: "Post-Incident-Activity",
		},
	}
}

func (c AlertIpdrrPhase) Value() string {
	return c.value
}

func (c AlertIpdrrPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertIpdrrPhase) UnmarshalJSON(b []byte) error {
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

type AlertCloseReason struct {
	value string
}

type AlertCloseReasonEnum struct {
	FALSE_DETECTION AlertCloseReason
	RESOLVED        AlertCloseReason
	REPEATED        AlertCloseReason
	OTHER           AlertCloseReason
}

func GetAlertCloseReasonEnum() AlertCloseReasonEnum {
	return AlertCloseReasonEnum{
		FALSE_DETECTION: AlertCloseReason{
			value: "False detection",
		},
		RESOLVED: AlertCloseReason{
			value: "Resolved",
		},
		REPEATED: AlertCloseReason{
			value: "Repeated",
		},
		OTHER: AlertCloseReason{
			value: "Other",
		},
	}
}

func (c AlertCloseReason) Value() string {
	return c.value
}

func (c AlertCloseReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertCloseReason) UnmarshalJSON(b []byte) error {
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
