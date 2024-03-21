package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type DetachInternetBandwidthGlobalEips struct {

	// 全域弹性公网IP的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：全域弹性公网IP名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// - 租户账号ID，获取租户账号ID请参见[租户账号ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)
	DomainId *string `json:"domain_id,omitempty"`

	// - 功能说明：用户自定义的资源描述 - 约束：   - 值的长度最大512字符，由数字、字母、中文、_(下划线)、-（中划线）、.（点）组成。
	Description *string `json:"description,omitempty"`

	// 接入点信息
	AccessSite *string `json:"access_site,omitempty"`

	// 全域弹性公网IP池子名称
	GeipPoolName *string `json:"geip_pool_name,omitempty"`

	// 全域弹性公网IP所属线路
	Isp *string `json:"isp,omitempty"`

	// - 功能说明：全域弹性公网IP的版本 - 取值范围：4、6
	IpVersion *DetachInternetBandwidthGlobalEipsIpVersion `json:"ip_version,omitempty"`

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
	Status *DetachInternetBandwidthGlobalEipsStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	InternetBandwidthInfo *InternetBandwidthInfo `json:"internet_bandwidth_info,omitempty"`

	GlobalConnectionBandwidthInfo *GlobalConnectionBandwidthInfo `json:"global_connection_bandwidth_info,omitempty"`

	AssociateInstanceInfo *InstanceInfo `json:"associate_instance_info,omitempty"`

	// 是否包周期
	IsPrePaid *bool `json:"is_pre_paid,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]Tag `json:"sys_tags,omitempty"`

	// - 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。 - 创建全域弹性公网IP时，给全域弹性公网IP绑定企业项目ID。 - 不指定该参数时，默认值是 0 - 关于企业项目ID的获取及企业项目特性的详细信息，请参见[《企业管理用户指南》](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。
	EnterpriseProjectId *interface{} `json:"enterprise_project_id,omitempty"`
}

func (o DetachInternetBandwidthGlobalEips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInternetBandwidthGlobalEips struct{}"
	}

	return strings.Join([]string{"DetachInternetBandwidthGlobalEips", string(data)}, " ")
}

type DetachInternetBandwidthGlobalEipsIpVersion struct {
	value int32
}

type DetachInternetBandwidthGlobalEipsIpVersionEnum struct {
	E_4 DetachInternetBandwidthGlobalEipsIpVersion
	E_6 DetachInternetBandwidthGlobalEipsIpVersion
}

func GetDetachInternetBandwidthGlobalEipsIpVersionEnum() DetachInternetBandwidthGlobalEipsIpVersionEnum {
	return DetachInternetBandwidthGlobalEipsIpVersionEnum{
		E_4: DetachInternetBandwidthGlobalEipsIpVersion{
			value: 4,
		}, E_6: DetachInternetBandwidthGlobalEipsIpVersion{
			value: 6,
		},
	}
}

func (c DetachInternetBandwidthGlobalEipsIpVersion) Value() int32 {
	return c.value
}

func (c DetachInternetBandwidthGlobalEipsIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetachInternetBandwidthGlobalEipsIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type DetachInternetBandwidthGlobalEipsStatus struct {
	value string
}

type DetachInternetBandwidthGlobalEipsStatusEnum struct {
	PENDING_CREATE DetachInternetBandwidthGlobalEipsStatus
	IDLE           DetachInternetBandwidthGlobalEipsStatus
	INUSE          DetachInternetBandwidthGlobalEipsStatus
	PENDING_UPDATE DetachInternetBandwidthGlobalEipsStatus
}

func GetDetachInternetBandwidthGlobalEipsStatusEnum() DetachInternetBandwidthGlobalEipsStatusEnum {
	return DetachInternetBandwidthGlobalEipsStatusEnum{
		PENDING_CREATE: DetachInternetBandwidthGlobalEipsStatus{
			value: "PENDING_CREATE",
		},
		IDLE: DetachInternetBandwidthGlobalEipsStatus{
			value: "IDLE",
		},
		INUSE: DetachInternetBandwidthGlobalEipsStatus{
			value: "INUSE",
		},
		PENDING_UPDATE: DetachInternetBandwidthGlobalEipsStatus{
			value: "PENDING_UPDATE",
		},
	}
}

func (c DetachInternetBandwidthGlobalEipsStatus) Value() string {
	return c.value
}

func (c DetachInternetBandwidthGlobalEipsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetachInternetBandwidthGlobalEipsStatus) UnmarshalJSON(b []byte) error {
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
