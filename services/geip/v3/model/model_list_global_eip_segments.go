package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListGlobalEipSegments struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 接入点信息
	AccessSite *string `json:"access_site,omitempty"`

	// 全域弹性公网IP池子名称
	GeipPoolName *string `json:"geip_pool_name,omitempty"`

	// 线路
	Isp *string `json:"isp,omitempty"`

	// IPv4或IPv6
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 全域公网IP段的cidr
	Cidr *string `json:"cidr,omitempty"`

	// 指定cidr-v6创建
	CidrV6 *string `json:"cidr_v6,omitempty"`

	// 是否冻结
	Freezen *bool `json:"freezen,omitempty"`

	// 冻结原因
	FreezenInfo *string `json:"freezen_info,omitempty"`

	// 状态
	Status *ListGlobalEipSegmentsStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	InternetBandwidth *InternetBandwidthInfo `json:"internet_bandwidth,omitempty"`

	AssociateInstance *InstanceInfo `json:"associate_instance,omitempty"`

	// 是否包周期
	IsPrePaid *bool `json:"is_pre_paid,omitempty"`

	// 全域弹性公网IP段标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]Tag `json:"sys_tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListGlobalEipSegments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipSegments struct{}"
	}

	return strings.Join([]string{"ListGlobalEipSegments", string(data)}, " ")
}

type ListGlobalEipSegmentsStatus struct {
	value string
}

type ListGlobalEipSegmentsStatusEnum struct {
	IDLE  ListGlobalEipSegmentsStatus
	INUSE ListGlobalEipSegmentsStatus
}

func GetListGlobalEipSegmentsStatusEnum() ListGlobalEipSegmentsStatusEnum {
	return ListGlobalEipSegmentsStatusEnum{
		IDLE: ListGlobalEipSegmentsStatus{
			value: "IDLE",
		},
		INUSE: ListGlobalEipSegmentsStatus{
			value: "INUSE",
		},
	}
}

func (c ListGlobalEipSegmentsStatus) Value() string {
	return c.value
}

func (c ListGlobalEipSegmentsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipSegmentsStatus) UnmarshalJSON(b []byte) error {
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
