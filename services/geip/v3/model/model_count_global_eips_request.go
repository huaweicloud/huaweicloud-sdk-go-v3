package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountGlobalEipsRequest Request Object
type CountGlobalEipsRequest struct {

	// 根据ID过滤
	Id *[]string `json:"id,omitempty"`

	// 根据全域公网带宽的ID过滤
	InternetBandwidthId *[]string `json:"internet_bandwidth_id,omitempty"`

	// 根据名称过滤
	Name *[]string `json:"name,omitempty"`

	// 根据名称模糊匹配
	NameLike *string `json:"name_like,omitempty"`

	// 根据接入点过滤
	AccessSite *[]string `json:"access_site,omitempty"`

	// 根据全域弹性公网IP池名称过滤
	GeipPoolName *[]string `json:"geip_pool_name,omitempty"`

	// 根据运营商线路过滤
	Isp *[]string `json:"isp,omitempty"`

	// 根据IP版本过滤
	IpVersion *[]CountGlobalEipsRequestIpVersion `json:"ip_version,omitempty"`

	// 根据ip地址过滤
	IpAddress *[]string `json:"ip_address,omitempty"`

	// 根据ipv6地址过滤
	Ipv6Address *[]string `json:"ipv6_address,omitempty"`

	// 根据是否冻结过滤
	Freezen *[]bool `json:"freezen,omitempty"`

	// 根据是否污染过滤
	Polluted *[]bool `json:"polluted,omitempty"`

	// 根据是否绑定全域公网带宽过滤
	InternetBandwidthIsNull *[]bool `json:"internet_bandwidth_is_null,omitempty"`

	// 根据是否绑定骨干带宽过滤
	GcbBandwidthIsNull *[]bool `json:"gcb_bandwidth_is_null,omitempty"`

	// 根据资源状态过滤
	Status *[]CountGlobalEipsRequestStatus `json:"status,omitempty"`

	// 根据绑定实例所属的局点过滤
	AssociateInstanceInfoRegion *[]string `json:"associate_instance_info.region,omitempty"`

	// 根据绑定实例所属的边缘信息过滤
	AssociateInstanceInfoPublicBorderGroup *[]string `json:"associate_instance_info.public_border_group,omitempty"`

	// 根据绑定实例所在的站点过滤
	AssociateInstanceInfoInstanceSite *[]string `json:"associate_instance_info.instance_site,omitempty"`

	// 根据绑定实例的类型过滤
	AssociateInstanceInfoInstanceType *[]string `json:"associate_instance_info.instance_type,omitempty"`

	// 根据绑定实例的ID过滤
	AssociateInstanceInfoInstanceId *[]string `json:"associate_instance_info.instance_id,omitempty"`

	// query by associate_instance_info.project_id
	AssociateInstanceInfoProjectId *[]string `json:"associate_instance_info.project_id,omitempty"`

	// 根据企业项目ID过滤
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据标签过滤
	Tags *[]string `json:"tags,omitempty"`
}

func (o CountGlobalEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGlobalEipsRequest struct{}"
	}

	return strings.Join([]string{"CountGlobalEipsRequest", string(data)}, " ")
}

type CountGlobalEipsRequestIpVersion struct {
	value int32
}

type CountGlobalEipsRequestIpVersionEnum struct {
	E_4 CountGlobalEipsRequestIpVersion
	E_6 CountGlobalEipsRequestIpVersion
}

func GetCountGlobalEipsRequestIpVersionEnum() CountGlobalEipsRequestIpVersionEnum {
	return CountGlobalEipsRequestIpVersionEnum{
		E_4: CountGlobalEipsRequestIpVersion{
			value: 4,
		}, E_6: CountGlobalEipsRequestIpVersion{
			value: 6,
		},
	}
}

func (c CountGlobalEipsRequestIpVersion) Value() int32 {
	return c.value
}

func (c CountGlobalEipsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountGlobalEipsRequestIpVersion) UnmarshalJSON(b []byte) error {
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

type CountGlobalEipsRequestStatus struct {
	value string
}

type CountGlobalEipsRequestStatusEnum struct {
	IDLE           CountGlobalEipsRequestStatus
	INUSE          CountGlobalEipsRequestStatus
	PENDING_CREATE CountGlobalEipsRequestStatus
	PENDING_UPDATE CountGlobalEipsRequestStatus
}

func GetCountGlobalEipsRequestStatusEnum() CountGlobalEipsRequestStatusEnum {
	return CountGlobalEipsRequestStatusEnum{
		IDLE: CountGlobalEipsRequestStatus{
			value: "idle",
		},
		INUSE: CountGlobalEipsRequestStatus{
			value: "inuse",
		},
		PENDING_CREATE: CountGlobalEipsRequestStatus{
			value: "pending_create",
		},
		PENDING_UPDATE: CountGlobalEipsRequestStatus{
			value: "pending_update",
		},
	}
}

func (c CountGlobalEipsRequestStatus) Value() string {
	return c.value
}

func (c CountGlobalEipsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountGlobalEipsRequestStatus) UnmarshalJSON(b []byte) error {
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
