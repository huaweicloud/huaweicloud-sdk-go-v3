package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ShowInternetBandwidth struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 线路
	Isp *string `json:"isp,omitempty"`

	// 全域公网带宽大小（入云方向）
	IngressSize *int32 `json:"ingress_size,omitempty"`

	// 接入点信息
	AccessSite *string `json:"access_site,omitempty"`

	// 全域公网带宽大小（出云方向）
	Size *int32 `json:"size,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 增强95保底率
	Ratio95peak *int32 `json:"ratio_95peak,omitempty"`

	// 订单信息
	BillingInfo *string `json:"billing_info,omitempty"`

	// 冻结原因
	FreezenInfo *string `json:"freezen_info,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 状态
	Status *ShowInternetBandwidthStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 是否包周期
	IsPrePaid *bool `json:"is_pre_paid,omitempty"`

	// 全域公网带宽标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]Tag `json:"sys_tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 全域公网带宽类型
	Type *string `json:"type,omitempty"`

	// 全域公网带宽资源的锁状态
	LockInfos *[]LockInfo `json:"lock_infos,omitempty"`
}

func (o ShowInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInternetBandwidth struct{}"
	}

	return strings.Join([]string{"ShowInternetBandwidth", string(data)}, " ")
}

type ShowInternetBandwidthStatus struct {
	value string
}

type ShowInternetBandwidthStatusEnum struct {
	NORMAL  ShowInternetBandwidthStatus
	FREEZED ShowInternetBandwidthStatus
}

func GetShowInternetBandwidthStatusEnum() ShowInternetBandwidthStatusEnum {
	return ShowInternetBandwidthStatusEnum{
		NORMAL: ShowInternetBandwidthStatus{
			value: "NORMAL",
		},
		FREEZED: ShowInternetBandwidthStatus{
			value: "FREEZED",
		},
	}
}

func (c ShowInternetBandwidthStatus) Value() string {
	return c.value
}

func (c ShowInternetBandwidthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInternetBandwidthStatus) UnmarshalJSON(b []byte) error {
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
