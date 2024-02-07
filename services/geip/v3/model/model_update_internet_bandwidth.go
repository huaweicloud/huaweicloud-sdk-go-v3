package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type UpdateInternetBandwidth struct {

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

	// 冻结原因
	FreezenInfo *string `json:"freezen_info,omitempty"`

	// 订单信息
	BillingInfo *string `json:"billing_info,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 状态
	Status *UpdateInternetBandwidthStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 是否包周期
	IsPrePaid *bool `json:"is_pre_paid,omitempty"`

	// 全域公网带宽标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 全域公网带宽资源的锁状态
	LockInfos *[]LockInfo `json:"lock_infos,omitempty"`
}

func (o UpdateInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternetBandwidth struct{}"
	}

	return strings.Join([]string{"UpdateInternetBandwidth", string(data)}, " ")
}

type UpdateInternetBandwidthStatus struct {
	value string
}

type UpdateInternetBandwidthStatusEnum struct {
	NORMAL  UpdateInternetBandwidthStatus
	FREEZED UpdateInternetBandwidthStatus
}

func GetUpdateInternetBandwidthStatusEnum() UpdateInternetBandwidthStatusEnum {
	return UpdateInternetBandwidthStatusEnum{
		NORMAL: UpdateInternetBandwidthStatus{
			value: "NORMAL",
		},
		FREEZED: UpdateInternetBandwidthStatus{
			value: "FREEZED",
		},
	}
}

func (c UpdateInternetBandwidthStatus) Value() string {
	return c.value
}

func (c UpdateInternetBandwidthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInternetBandwidthStatus) UnmarshalJSON(b []byte) error {
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
