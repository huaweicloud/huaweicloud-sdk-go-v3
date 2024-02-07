package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BatchCreateInternetBandwidth struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 线路
	Isp string `json:"isp"`

	// 全域公网带宽大小（入云方向）
	IngressSize *int32 `json:"ingress_size,omitempty"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	// 全域公网带宽大小（出云方向）
	Size int32 `json:"size"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 状态
	Status *BatchCreateInternetBandwidthStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 是否创建成功标识，取值：successful、failed。
	RetStatus *string `json:"ret_status,omitempty"`

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

func (o BatchCreateInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidth struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidth", string(data)}, " ")
}

type BatchCreateInternetBandwidthStatus struct {
	value string
}

type BatchCreateInternetBandwidthStatusEnum struct {
	NORMAL  BatchCreateInternetBandwidthStatus
	FREEZED BatchCreateInternetBandwidthStatus
}

func GetBatchCreateInternetBandwidthStatusEnum() BatchCreateInternetBandwidthStatusEnum {
	return BatchCreateInternetBandwidthStatusEnum{
		NORMAL: BatchCreateInternetBandwidthStatus{
			value: "NORMAL",
		},
		FREEZED: BatchCreateInternetBandwidthStatus{
			value: "FREEZED",
		},
	}
}

func (c BatchCreateInternetBandwidthStatus) Value() string {
	return c.value
}

func (c BatchCreateInternetBandwidthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInternetBandwidthStatus) UnmarshalJSON(b []byte) error {
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
