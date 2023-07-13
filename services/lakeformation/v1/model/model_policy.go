package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Policy struct {

	// 排除的允许策略
	AllowExceptions *[]PolicyItem `json:"allow_exceptions,omitempty"`

	// 条件属性
	Conditions *[]PolicyItemCondition `json:"conditions,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 类掩码策略条目
	DataMaskPolicyItems *[]DataMaskPolicyItem `json:"data_mask_policy_items,omitempty"`

	// 拒绝排除策略
	DenyExceptions *[]PolicyItem `json:"deny_exceptions,omitempty"`

	// 拒绝策略
	DenyPolicyItems *[]PolicyItem `json:"deny_policy_items,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 唯一guid
	Guid *string `json:"guid,omitempty"`

	// 主键
	Id *int64 `json:"id,omitempty"`

	// 是否支持审计
	IsAuditEnabled *bool `json:"is_audit_enabled,omitempty"`

	// 是否默认策略
	IsDefaultPolicy *bool `json:"is_default_policy,omitempty"`

	// 是否拒绝所有
	IsDenyAllElse *bool `json:"is_deny_all_else,omitempty"`

	// 是否启动
	IsEnabled *bool `json:"is_enabled,omitempty"`

	// 名
	Name *string `json:"name,omitempty"`

	// 选项
	Options *interface{} `json:"options,omitempty"`

	// 策略信息数组
	PolicyItems *[]PolicyItem `json:"policy_items,omitempty"`

	// 策略便签
	PolicyLabels *[]string `json:"policy_labels,omitempty"`

	// 策略优先级
	PolicyPriority *int32 `json:"policy_priority,omitempty"`

	// 策略类型
	PolicyType *int32 `json:"policy_type,omitempty"`

	// 资源签名
	ResourceSignature *string `json:"resource_signature,omitempty"`

	// 资源
	Resources map[string]PolicyResource `json:"resources,omitempty"`

	// 行过滤策略条目
	RowFilterPolicyItems *[]RowFilterPolicyItem `json:"row_filter_policy_items,omitempty"`

	// 服务
	Service *string `json:"service,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 验证周期
	ValiditySchedules *[]ValiditySchedule `json:"validity_schedules,omitempty"`

	// 版本
	Version *int64 `json:"version,omitempty"`

	// 域
	ZoneName *string `json:"zone_name,omitempty"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}
