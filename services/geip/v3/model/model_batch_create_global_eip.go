package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BatchCreateGlobalEip struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	// 全域弹性公网IP池子名称
	GeipPoolName string `json:"geip_pool_name"`

	// 线路
	Isp *string `json:"isp,omitempty"`

	// IPv4或IPv6
	IpVersion *int32 `json:"ip_version,omitempty"`

	// IPv4地址
	IpAddress *string `json:"ip_address,omitempty"`

	// IPv6地址
	Ipv6Address *string `json:"ipv6_address,omitempty"`

	// 是否冻结
	Freezen *bool `json:"freezen,omitempty"`

	// 冻结原因
	FreezenInfo *string `json:"freezen_info,omitempty"`

	// 是否污染
	Polluted *bool `json:"polluted,omitempty"`

	// 状态
	Status *BatchCreateGlobalEipStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	InternetBandwidthInfo *InternetBandwidthInfo `json:"internet_bandwidth_info"`

	GlobalConnectionBandwidthInfo *GlobalConnectionBandwidthInfo `json:"global_connection_bandwidth_info,omitempty"`

	AssociateInstanceInfo *InstanceInfo `json:"associate_instance_info,omitempty"`

	// 是否包周期
	IsPrePaid *bool `json:"is_pre_paid,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]Tag `json:"sys_tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o BatchCreateGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEip struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEip", string(data)}, " ")
}

type BatchCreateGlobalEipStatus struct {
	value string
}

type BatchCreateGlobalEipStatusEnum struct {
	PENDING_CREATE BatchCreateGlobalEipStatus
	IDLE           BatchCreateGlobalEipStatus
	INUSE          BatchCreateGlobalEipStatus
	PENDING_UPDATE BatchCreateGlobalEipStatus
}

func GetBatchCreateGlobalEipStatusEnum() BatchCreateGlobalEipStatusEnum {
	return BatchCreateGlobalEipStatusEnum{
		PENDING_CREATE: BatchCreateGlobalEipStatus{
			value: "PENDING_CREATE",
		},
		IDLE: BatchCreateGlobalEipStatus{
			value: "IDLE",
		},
		INUSE: BatchCreateGlobalEipStatus{
			value: "INUSE",
		},
		PENDING_UPDATE: BatchCreateGlobalEipStatus{
			value: "PENDING_UPDATE",
		},
	}
}

func (c BatchCreateGlobalEipStatus) Value() string {
	return c.value
}

func (c BatchCreateGlobalEipStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateGlobalEipStatus) UnmarshalJSON(b []byte) error {
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
