package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Event event
type Event struct {

	// SA数据对象版本号，数据接入时需携带版本号。版本号由SA服务团队负责更新，数据源只可填写SA给定的版本号。目前版本为1.0.0。
	Version string `json:"version"`

	Environment *Environment `json:"environment"`

	DataSource *DataSource `json:"data_source"`

	// 首次发现时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。
	FirstObservedTime *sdktime.SdkTime `json:"first_observed_time"`

	// 最新发现时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。
	LastObservedTime *sdktime.SdkTime `json:"last_observed_time,omitempty"`

	// 记录时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 数据接收时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。  是指事件数据被SA侧接收的时间，由SA接收时填写，产品上报数据时不用填写。
	ArriveTime *sdktime.SdkTime `json:"arrive_time,omitempty"`

	// 事件唯一标识，UUID格式。
	EventId string `json:"event_id"`

	// 事件标题，最大255字符。
	Title string `json:"title"`

	// 事件描述信息，最大1024个字符
	Description string `json:"description"`

	// 事件URL链接，指向数据源产品中有关当前事件说明的页面。
	SourceUrl *string `json:"source_url,omitempty"`

	// 事件发生次数，默认为1，必填。
	Count int32 `json:"count"`

	// 事件的置信度。置信度的定义旨在说明识别的行为或问题的可能性。 取值范围：0-100，0表示置信度为0%，100表示置信度为100%。
	Confidence *int32 `json:"confidence,omitempty"`

	Severity *Severity `json:"severity"`

	// 关键性，是指事件涉及的资源的重要性级别。 取值范围：0-100，0表示资源不关键，100表示最关键资源。
	Criticality *int32 `json:"criticality,omitempty"`

	Type *Type `json:"type"`

	Compliance *Compliance `json:"compliance,omitempty"`

	Network *Network `json:"network,omitempty"`

	VulnerabilityPatch *VulnerabilityPatch `json:"vulnerability_patch,omitempty"`

	Malware *Malware `json:"malware,omitempty"`

	ThreatIntel *ThreatIntel `json:"threat_intel,omitempty"`

	Resource *Resource `json:"resource"`

	Remediation *Remediation `json:"remediation,omitempty"`

	// 数据源自定义信息，最多支持50个key/value对，约束条件： 1、该对象不能包含冗余数据，并且不能与已定义的SSA事件格式字段冲突。 2、字段名称可以包含字母数字字符、空格和以下符号：_ . / = + \\ - @。 示例： \"data_source_fields\": {     \"key1\": \"value1\",     \"key2\", \"value2\",   }
	DataSourceFields *interface{} `json:"data_source_fields,omitempty"`

	// 事件验证状态，标识事件的准确性。 Unknown – 未知，默认 True_positive – 确认 False_positive – 误报。
	VerificationState *string `json:"verification_state,omitempty"`

	// 事件处理状态，New/Ignored/Resolved；默认New。
	HandleStatus string `json:"handle_status"`

	// 阶段：Prepartion|Detection and Analysis|Containm，Eradication& Recovery| Post-Incident-Activity。
	Phase *string `json:"phase,omitempty"`

	// 约束闭环时间：单位：天。
	Sla *int32 `json:"sla,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
