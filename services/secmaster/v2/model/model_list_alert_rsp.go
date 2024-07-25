package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRsp 告警实体信息
type ListAlertRsp struct {

	// 告警对象的版本，该字段的值必须为华为云SSA服务确定的官方发布版本之一
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
	Severity *ListAlertRspSeverity `json:"severity,omitempty"`

	// 关键性，是指事件涉及的资源的重要性级别。 取值范围：0-100，0表示资源不关键，100表示最关键资源
	Criticality *int32 `json:"criticality,omitempty"`

	AlertType *AlertAlertType `json:"alert_type,omitempty"`

	// 网络信息
	NetworkList *[]AlertNetworkList `json:"network_list,omitempty"`

	// 受影响资源
	ResourceList *[]AlertResourceList `json:"resource_list,omitempty"`

	Remediation *AlertRemediation `json:"remediation,omitempty"`

	// 验证状态，标识事件的准确性。可选类型如下： Unknown – 未知 True_Positive – 确认 False_Positive – 误报 默认填写Unknown
	VerificationState *ListAlertRspVerificationState `json:"verification_state,omitempty"`

	// 事件处理状态，可选类型如下： Open – 打开，默认 Block – 阻塞 Closed – 关闭 默认填写Open
	HandleStatus *ListAlertRspHandleStatus `json:"handle_status,omitempty"`

	// 约束闭环时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	Sla *string `json:"sla,omitempty"`

	// 更新时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	UpdateTime *string `json:"update_time,omitempty"`

	// 关闭时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	CloseTime *string `json:"close_time,omitempty"`

	// 周期/处置阶段编号 Prepartion|Detection and Analysis|Containm，Eradication& Recovery|Post-Incident-Activity
	IpdrrPhase *ListAlertRspIpdrrPhase `json:"ipdrr_phase,omitempty"`

	// 周期/处置阶段编号 Prepartion|Detection and Analysis|Containm，Eradication& Recovery|Post-Incident-Activity
	ChopPhase *ListAlertRspChopPhase `json:"chop_phase,omitempty"`

	// 周期/处置阶段编号 Prepartion|Detection and Analysis|Containm，Eradication& Recovery|Post-Incident-Activity
	PpdrPhase *ListAlertRspPpdrPhase `json:"ppdr_phase,omitempty"`

	// 调试字段
	Simulation *string `json:"simulation,omitempty"`

	// 告警调查员
	Actor *string `json:"actor,omitempty"`

	// 责任人、服务责任人
	Owner *string `json:"owner,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 关闭原因: 误检 - False detection 已解决 - Resolved 重复 - Repeated 其他 - Other
	CloseReason *ListAlertRspCloseReason `json:"close_reason,omitempty"`

	// 关闭评论
	CloseComment *string `json:"close_comment,omitempty"`

	Malware *ShowAlertRspMalware `json:"malware,omitempty"`

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

func (o ListAlertRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRsp struct{}"
	}

	return strings.Join([]string{"ListAlertRsp", string(data)}, " ")
}

type ListAlertRspSeverity struct {
	value string
}

type ListAlertRspSeverityEnum struct {
	TIPS   ListAlertRspSeverity
	LOW    ListAlertRspSeverity
	MEDIUM ListAlertRspSeverity
	HIGH   ListAlertRspSeverity
	FATAL  ListAlertRspSeverity
}

func GetListAlertRspSeverityEnum() ListAlertRspSeverityEnum {
	return ListAlertRspSeverityEnum{
		TIPS: ListAlertRspSeverity{
			value: "Tips",
		},
		LOW: ListAlertRspSeverity{
			value: "Low",
		},
		MEDIUM: ListAlertRspSeverity{
			value: "Medium",
		},
		HIGH: ListAlertRspSeverity{
			value: "High",
		},
		FATAL: ListAlertRspSeverity{
			value: "Fatal",
		},
	}
}

func (c ListAlertRspSeverity) Value() string {
	return c.value
}

func (c ListAlertRspSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspSeverity) UnmarshalJSON(b []byte) error {
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

type ListAlertRspVerificationState struct {
	value string
}

type ListAlertRspVerificationStateEnum struct {
	UNKNOWN        ListAlertRspVerificationState
	TRUE_POSITIVE  ListAlertRspVerificationState
	FALSE_POSITIVE ListAlertRspVerificationState
}

func GetListAlertRspVerificationStateEnum() ListAlertRspVerificationStateEnum {
	return ListAlertRspVerificationStateEnum{
		UNKNOWN: ListAlertRspVerificationState{
			value: "Unknown",
		},
		TRUE_POSITIVE: ListAlertRspVerificationState{
			value: "True_Positive",
		},
		FALSE_POSITIVE: ListAlertRspVerificationState{
			value: "False_Positive",
		},
	}
}

func (c ListAlertRspVerificationState) Value() string {
	return c.value
}

func (c ListAlertRspVerificationState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspVerificationState) UnmarshalJSON(b []byte) error {
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

type ListAlertRspHandleStatus struct {
	value string
}

type ListAlertRspHandleStatusEnum struct {
	OPEN   ListAlertRspHandleStatus
	BLOCK  ListAlertRspHandleStatus
	CLOSED ListAlertRspHandleStatus
}

func GetListAlertRspHandleStatusEnum() ListAlertRspHandleStatusEnum {
	return ListAlertRspHandleStatusEnum{
		OPEN: ListAlertRspHandleStatus{
			value: "Open",
		},
		BLOCK: ListAlertRspHandleStatus{
			value: "Block",
		},
		CLOSED: ListAlertRspHandleStatus{
			value: "Closed",
		},
	}
}

func (c ListAlertRspHandleStatus) Value() string {
	return c.value
}

func (c ListAlertRspHandleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspHandleStatus) UnmarshalJSON(b []byte) error {
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

type ListAlertRspIpdrrPhase struct {
	value string
}

type ListAlertRspIpdrrPhaseEnum struct {
	PREPARTION                   ListAlertRspIpdrrPhase
	DETECTION_AND_ANALYSIS       ListAlertRspIpdrrPhase
	CONTAINMERADICATION_RECOVERY ListAlertRspIpdrrPhase
	POST_INCIDENT_ACTIVITY       ListAlertRspIpdrrPhase
}

func GetListAlertRspIpdrrPhaseEnum() ListAlertRspIpdrrPhaseEnum {
	return ListAlertRspIpdrrPhaseEnum{
		PREPARTION: ListAlertRspIpdrrPhase{
			value: "Prepartion",
		},
		DETECTION_AND_ANALYSIS: ListAlertRspIpdrrPhase{
			value: "Detection and Analysis",
		},
		CONTAINMERADICATION_RECOVERY: ListAlertRspIpdrrPhase{
			value: "Containm，Eradication& Recovery",
		},
		POST_INCIDENT_ACTIVITY: ListAlertRspIpdrrPhase{
			value: "Post-Incident-Activity",
		},
	}
}

func (c ListAlertRspIpdrrPhase) Value() string {
	return c.value
}

func (c ListAlertRspIpdrrPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspIpdrrPhase) UnmarshalJSON(b []byte) error {
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

type ListAlertRspChopPhase struct {
	value string
}

type ListAlertRspChopPhaseEnum struct {
	PREPARTION                   ListAlertRspChopPhase
	DETECTION_AND_ANALYSIS       ListAlertRspChopPhase
	CONTAINMERADICATION_RECOVERY ListAlertRspChopPhase
	POST_INCIDENT_ACTIVITY       ListAlertRspChopPhase
}

func GetListAlertRspChopPhaseEnum() ListAlertRspChopPhaseEnum {
	return ListAlertRspChopPhaseEnum{
		PREPARTION: ListAlertRspChopPhase{
			value: "Prepartion",
		},
		DETECTION_AND_ANALYSIS: ListAlertRspChopPhase{
			value: "Detection and Analysis",
		},
		CONTAINMERADICATION_RECOVERY: ListAlertRspChopPhase{
			value: "Containm，Eradication& Recovery",
		},
		POST_INCIDENT_ACTIVITY: ListAlertRspChopPhase{
			value: "Post-Incident-Activity",
		},
	}
}

func (c ListAlertRspChopPhase) Value() string {
	return c.value
}

func (c ListAlertRspChopPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspChopPhase) UnmarshalJSON(b []byte) error {
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

type ListAlertRspPpdrPhase struct {
	value string
}

type ListAlertRspPpdrPhaseEnum struct {
	PREPARTION                   ListAlertRspPpdrPhase
	DETECTION_AND_ANALYSIS       ListAlertRspPpdrPhase
	CONTAINMERADICATION_RECOVERY ListAlertRspPpdrPhase
	POST_INCIDENT_ACTIVITY       ListAlertRspPpdrPhase
}

func GetListAlertRspPpdrPhaseEnum() ListAlertRspPpdrPhaseEnum {
	return ListAlertRspPpdrPhaseEnum{
		PREPARTION: ListAlertRspPpdrPhase{
			value: "Prepartion",
		},
		DETECTION_AND_ANALYSIS: ListAlertRspPpdrPhase{
			value: "Detection and Analysis",
		},
		CONTAINMERADICATION_RECOVERY: ListAlertRspPpdrPhase{
			value: "Containm，Eradication& Recovery",
		},
		POST_INCIDENT_ACTIVITY: ListAlertRspPpdrPhase{
			value: "Post-Incident-Activity",
		},
	}
}

func (c ListAlertRspPpdrPhase) Value() string {
	return c.value
}

func (c ListAlertRspPpdrPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspPpdrPhase) UnmarshalJSON(b []byte) error {
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

type ListAlertRspCloseReason struct {
	value string
}

type ListAlertRspCloseReasonEnum struct {
	FALSE_DETECTION ListAlertRspCloseReason
	RESOLVED        ListAlertRspCloseReason
	REPEATED        ListAlertRspCloseReason
	OTHER           ListAlertRspCloseReason
}

func GetListAlertRspCloseReasonEnum() ListAlertRspCloseReasonEnum {
	return ListAlertRspCloseReasonEnum{
		FALSE_DETECTION: ListAlertRspCloseReason{
			value: "False detection",
		},
		RESOLVED: ListAlertRspCloseReason{
			value: "Resolved",
		},
		REPEATED: ListAlertRspCloseReason{
			value: "Repeated",
		},
		OTHER: ListAlertRspCloseReason{
			value: "Other",
		},
	}
}

func (c ListAlertRspCloseReason) Value() string {
	return c.value
}

func (c ListAlertRspCloseReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRspCloseReason) UnmarshalJSON(b []byte) error {
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
