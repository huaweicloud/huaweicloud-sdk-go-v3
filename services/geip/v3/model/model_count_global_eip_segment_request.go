package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountGlobalEipSegmentRequest Request Object
type CountGlobalEipSegmentRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 只显示指定的字段
	Fields *[]CountGlobalEipSegmentRequestFields `json:"fields,omitempty"`

	// 根据资源ID过滤
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
	IpVersion *[]CountGlobalEipSegmentRequestIpVersion `json:"ip_version,omitempty"`

	// 根据分配的CIDR过滤
	Cidr *[]string `json:"cidr,omitempty"`

	// 根据分配的IPv6 CIDR过滤
	CidrV6 *[]string `json:"cidr_v6,omitempty"`

	// 根据是否冻结过滤
	Freezen *[]bool `json:"freezen,omitempty"`

	// 根据是否绑定全域公网带宽过滤
	InternetBandwidthIsNull *[]bool `json:"internet_bandwidth_is_null,omitempty"`

	// 根据资源状态过滤
	Status *[]CountGlobalEipSegmentRequestStatus `json:"status,omitempty"`

	// 根据绑定实例所属的局点过滤
	AssociateInstanceRegion *[]string `json:"associate_instance.region,omitempty"`

	// 根据绑定实例所属的边缘信息过滤
	AssociateInstancePublicBorderGroup *[]string `json:"associate_instance.public_border_group,omitempty"`

	// 根据绑定实例所在的站点过滤
	AssociateInstanceInstanceSite *[]string `json:"associate_instance.instance_site,omitempty"`

	// 根据绑定实例的类型过滤
	AssociateInstanceInstanceType *[]string `json:"associate_instance.instance_type,omitempty"`

	// 根据绑定实例的ID过滤
	AssociateInstanceInstanceId *[]string `json:"associate_instance.instance_id,omitempty"`

	// 根据绑定实例所属的项目ID过滤
	AssociateInstanceProjectId *[]string `json:"associate_instance.project_id,omitempty"`

	// 根据企业项目ID过滤
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据标签过滤
	Tags *[]string `json:"tags,omitempty"`
}

func (o CountGlobalEipSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGlobalEipSegmentRequest struct{}"
	}

	return strings.Join([]string{"CountGlobalEipSegmentRequest", string(data)}, " ")
}

type CountGlobalEipSegmentRequestFields struct {
	value string
}

type CountGlobalEipSegmentRequestFieldsEnum struct {
	COUNT CountGlobalEipSegmentRequestFields
}

func GetCountGlobalEipSegmentRequestFieldsEnum() CountGlobalEipSegmentRequestFieldsEnum {
	return CountGlobalEipSegmentRequestFieldsEnum{
		COUNT: CountGlobalEipSegmentRequestFields{
			value: "count",
		},
	}
}

func (c CountGlobalEipSegmentRequestFields) Value() string {
	return c.value
}

func (c CountGlobalEipSegmentRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountGlobalEipSegmentRequestFields) UnmarshalJSON(b []byte) error {
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

type CountGlobalEipSegmentRequestIpVersion struct {
	value int32
}

type CountGlobalEipSegmentRequestIpVersionEnum struct {
	E_4 CountGlobalEipSegmentRequestIpVersion
	E_6 CountGlobalEipSegmentRequestIpVersion
}

func GetCountGlobalEipSegmentRequestIpVersionEnum() CountGlobalEipSegmentRequestIpVersionEnum {
	return CountGlobalEipSegmentRequestIpVersionEnum{
		E_4: CountGlobalEipSegmentRequestIpVersion{
			value: 4,
		}, E_6: CountGlobalEipSegmentRequestIpVersion{
			value: 6,
		},
	}
}

func (c CountGlobalEipSegmentRequestIpVersion) Value() int32 {
	return c.value
}

func (c CountGlobalEipSegmentRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountGlobalEipSegmentRequestIpVersion) UnmarshalJSON(b []byte) error {
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

type CountGlobalEipSegmentRequestStatus struct {
	value string
}

type CountGlobalEipSegmentRequestStatusEnum struct {
	IDLE           CountGlobalEipSegmentRequestStatus
	INUSE          CountGlobalEipSegmentRequestStatus
	PENDING_CREATE CountGlobalEipSegmentRequestStatus
	PENDING_UPDATE CountGlobalEipSegmentRequestStatus
}

func GetCountGlobalEipSegmentRequestStatusEnum() CountGlobalEipSegmentRequestStatusEnum {
	return CountGlobalEipSegmentRequestStatusEnum{
		IDLE: CountGlobalEipSegmentRequestStatus{
			value: "idle",
		},
		INUSE: CountGlobalEipSegmentRequestStatus{
			value: "inuse",
		},
		PENDING_CREATE: CountGlobalEipSegmentRequestStatus{
			value: "pending_create",
		},
		PENDING_UPDATE: CountGlobalEipSegmentRequestStatus{
			value: "pending_update",
		},
	}
}

func (c CountGlobalEipSegmentRequestStatus) Value() string {
	return c.value
}

func (c CountGlobalEipSegmentRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountGlobalEipSegmentRequestStatus) UnmarshalJSON(b []byte) error {
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
