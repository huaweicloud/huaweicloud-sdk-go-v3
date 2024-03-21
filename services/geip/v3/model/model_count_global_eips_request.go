package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountGlobalEipsRequest Request Object
type CountGlobalEipsRequest struct {
	Id *[]string `json:"id,omitempty"`

	InternetBandwidthId *[]string `json:"internet_bandwidth_id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	NameLike *string `json:"name_like,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	GeipPoolName *[]string `json:"geip_pool_name,omitempty"`

	Isp *[]string `json:"isp,omitempty"`

	IpVersion *[]CountGlobalEipsRequestIpVersion `json:"ip_version,omitempty"`

	IpAddress *[]string `json:"ip_address,omitempty"`

	Ipv6Address *[]string `json:"ipv6_address,omitempty"`

	Freezen *[]bool `json:"freezen,omitempty"`

	Polluted *[]bool `json:"polluted,omitempty"`

	InternetBandwidthIsNull *[]bool `json:"internet_bandwidth_is_null,omitempty"`

	GcbBandwidthIsNull *[]bool `json:"gcb_bandwidth_is_null,omitempty"`

	Status *[]CountGlobalEipsRequestStatus `json:"status,omitempty"`

	AssociateInstanceInfoRegion *[]string `json:"associate_instance_info.region,omitempty"`

	AssociateInstanceInfoPublicBorderGroup *[]string `json:"associate_instance_info.public_border_group,omitempty"`

	AssociateInstanceInfoInstanceSite *[]string `json:"associate_instance_info.instance_site,omitempty"`

	AssociateInstanceInfoInstanceType *[]string `json:"associate_instance_info.instance_type,omitempty"`

	AssociateInstanceInfoInstanceId *[]string `json:"associate_instance_info.instance_id,omitempty"`

	AssociateInstanceInfoProjectId *[]string `json:"associate_instance_info.project_id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

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
